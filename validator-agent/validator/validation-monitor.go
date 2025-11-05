package validator

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oasisprotocol/oasis-core/go/common/logging"
	sdkRofl "github.com/oasisprotocol/oasis-sdk/client-sdk/go/modules/rofl"

	"github.com/oasisprotocol/erc-8004/validator-agent/abi"
	"github.com/oasisprotocol/erc-8004/validator-agent/config"
	"github.com/oasisprotocol/erc-8004/validator-agent/rofl"
)

const IDENTITY_REGISTRY_METADATA_APP_ID_KEY = "net.oasis.app_id"

type ValidationReport struct {
	Score        uint8                 `json:"score"`
	AgentId      *big.Int              `json:"agent_id"`
	AppId        sdkRofl.AppID         `json:"app_id"`
	ScoreDetails map[string]uint8      `json:"score_details"`
	AppInfo      *sdkRofl.AppConfig    `json:"app_info"`
	ReplicaInfo  *sdkRofl.Registration `json:"replica_info"`
}

func NewValidationReport(agentId *big.Int) *ValidationReport {
	return &ValidationReport{
		Score:        0,
		AgentId:      agentId,
		AppId:        sdkRofl.AppID{},
		ScoreDetails: make(map[string]uint8),
	}
}

func (vr *ValidationReport) AddScore(s string) {
	vr.Score += ScoreTable[s]
	vr.ScoreDetails[s] = ScoreTable[s]
}

func (vr *ValidationReport) ToURI() (string, error) {
	registrationJSON, err := json.Marshal(vr)
	if err != nil {
		return "", err
	}
	return "data:application/json;base64," + base64.StdEncoding.EncodeToString(registrationJSON), nil
}

const ScoreReplicaNotExpiredYet = "replica-not-expired-yet"
const ScoreAtLeastOneReplica = "at-least-one-replica"

var ScoreTable = map[string]uint8{
	ScoreReplicaNotExpiredYet: 50,
	ScoreAtLeastOneReplica:    50,
}

const TagValidationInProgress = 0
const TagValidationFinished = 1

type ValidationMonitor struct {
	name       string
	cfg        *config.ValidationNetworkConfig
	signingKey []byte
	roflClient *rofl.RoflClient
	logger     *logging.Logger

	client             *ethclient.Client
	identityRegistry   *abi.IdentityRegistry
	validationRegistry *abi.ValidationRegistry
}

func NewValidationMonitor(name string, cfg *config.ValidationNetworkConfig, signingKey []byte, roflClient *rofl.RoflClient) *ValidationMonitor {
	return &ValidationMonitor{
		name:       name,
		cfg:        cfg,
		signingKey: signingKey,
		roflClient: roflClient,
		logger:     logging.GetLogger(fmt.Sprintf("validator-monitor/%s", name)),
	}
}

func (vm *ValidationMonitor) Start(ctx context.Context) {
	if err := vm.StartMonitor(ctx); err != nil {
		vm.logger.Error("failed to start ValidationMonitor", "err", err)
	}
}

func (vm *ValidationMonitor) StartMonitor(ctx context.Context) error {
	var err error
	vm.client, err = ethclient.DialContext(ctx, vm.cfg.RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer vm.client.Close()

	validationRegistryAddress := common.HexToAddress(vm.cfg.ValidationRegistry)
	vm.validationRegistry, err = abi.NewValidationRegistry(validationRegistryAddress, vm.client)
	if err != nil {
		vm.logger.Error("failed to parse validator registry", "err", err)
		return fmt.Errorf("failed to parse validator registry: %w", err)
	}

	identityRegistryAddress := common.HexToAddress(vm.cfg.IdentityRegistry)
	vm.identityRegistry, err = abi.NewIdentityRegistry(identityRegistryAddress, vm.client)
	if err != nil {
		vm.logger.Error("failed to parse identity registry", "err", err)
		return fmt.Errorf("failed to parse identity registry: %w", err)
	}

	privateKey, err := crypto.ToECDSA(vm.signingKey)
	if err != nil {
		return fmt.Errorf("failed to parse signing key: %w", err)
	}
	validatorAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	f, err := abi.NewValidationRegistryFilterer(validationRegistryAddress, vm.client)
	if err != nil {
		return fmt.Errorf("failed to create filterer: %w", err)
	}

	logs := make(chan *abi.ValidationRegistryValidationRequest)

	opts := &bind.WatchOpts{
		Context: ctx,
	}
	sub, err := f.WatchValidationRequest(opts, logs, []common.Address{validatorAddress}, []*big.Int{}, [][32]byte{})
	if err != nil {
		vm.logger.Error("failed to watch ValidationRequest events", "err", err)
		return fmt.Errorf("failed to watch ValidationRequest events: %w", err)
	}
	vm.logger.Info("listening for validation requests", "validation_registry", validationRegistryAddress, "validator_address", validatorAddress)
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			vm.logger.Error("subscription error", "err", err)
			return fmt.Errorf("subscription error: %w", err)
		case event := <-logs:
			vm.logger.Info("received ValidationRequest event from filterer",
				"requestHash", fmt.Sprintf("0x%x", event.RequestHash),
				"validator", event.ValidatorAddress.Hex(),
				"agentId", event.AgentId)

			report := NewValidationReport(event.AgentId)

			rawAppId, err := vm.identityRegistry.GetMetadata(&bind.CallOpts{Context: ctx}, event.AgentId, IDENTITY_REGISTRY_METADATA_APP_ID_KEY)
			if err != nil {
				vm.logger.Error("unable to get app ID for agent", "err", err, "agentId", event.AgentId)
				if err = vm.sendValidationResponse(event.RequestHash, report); err != nil {
					vm.logger.Error("unable to send validation response", "err", err, "report", report)
				}
				continue
			}
			if err = report.AppId.UnmarshalText(rawAppId); err != nil {
				vm.logger.Error("unable to unmarshal text for app ID", "err", err, "agentId", event.AgentId, "appId", string(rawAppId))
			}

			appInfo, replicas, lastEpoch, err := vm.roflClient.GetRoflReplicas(string(rawAppId))
			if err != nil {
				vm.logger.Error("unable to get ROFL info from Oasis chain", "err", err, "agentId", event.AgentId, "appId", string(rawAppId))
				if err = vm.sendValidationResponse(event.RequestHash, report); err != nil {
					vm.logger.Error("unable to send validation response", "err", err, "report", report)
				}
				continue
			}
			report.AppInfo = appInfo

			if len(replicas) != 0 {
				report.AddScore(ScoreAtLeastOneReplica)

				// Assume the most recent replica has the desired public keys.
				recentReplicaIdx := 0
				for i, r := range replicas {
					if r.Expiration > replicas[recentReplicaIdx].Expiration {
						recentReplicaIdx = i
					}
				}
				replica := replicas[recentReplicaIdx]

				report.ReplicaInfo = replica

				if replica.Expiration >= lastEpoch {
					report.AddScore(ScoreReplicaNotExpiredYet)
				}
			} else {
				vm.logger.Debug("ROFL has no replicas", "agentId", event.AgentId, "appId", string(rawAppId))
			}

			if err = vm.sendValidationResponse(event.RequestHash, report); err != nil {
				vm.logger.Error("unable to send validation response", "err", err, "report", report)
			}
		}
	}
}

func (vm *ValidationMonitor) sendValidationResponse(requestHash [32]byte, report *ValidationReport) error {
	_, auth, err := GetAddrAuth(vm.signingKey, vm.client)
	if err != nil {
		return err
	}

	reportURI, err := report.ToURI()
	if err != nil {
		return err
	}
	reportURIHash := crypto.Keccak256([]byte(reportURI))
	var tag [32]byte
	copy(tag[:], []byte{TagValidationFinished})
	tx, err := vm.validationRegistry.ValidationResponse(auth, requestHash, report.Score, reportURI, [32]byte(reportURIHash), tag)
	if err != nil {
		return err
	}

	var receipt *types.Receipt
	for i := 0; i < 300; i++ {
		receipt, err = vm.client.TransactionReceipt(auth.Context, tx.Hash())
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if receipt == nil {
		return fmt.Errorf("unable to get the receipt")
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return nil
}

func GetAddrAuth(key []byte, client *ethclient.Client) (common.Address, *bind.TransactOpts, error) {
	privateKey, err := crypto.ToECDSA(key)
	if err != nil {
		return common.Address{}, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, fmt.Errorf("error casting public key %s to ECDSA", publicKey)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to get account nonce: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to suggest gas price: %w", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to get transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice

	return fromAddress, auth, nil
}
