package validator

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"net/http"
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

// ValidationReport is a ROFL validation report.
type ValidationReport struct {
	Score        uint8                 `json:"score"`
	AgentId      *big.Int              `json:"agent_id"`
	AppId        sdkRofl.AppID         `json:"app_id"`
	ScoreDetails map[string]uint8      `json:"score_details"`
	Error        error                 `json:"error"`
	AppInfo      *sdkRofl.AppConfig    `json:"app_info"`
	ReplicaInfo  *sdkRofl.Registration `json:"replica_info"`
}

// UploadResponse is IPFS response.
type UploadResponse struct {
	Data struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Cid       string `json:"cid"`
		Size      int    `json:"size"`
		CreatedAt string `json:"created_at"`
		MimeType  string `json:"mime_type"`
		Network   string `json:"network"`
	}
}

// NewValidationReport creates a new ROFL validation report.
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

func (vr *ValidationReport) ToInlineURI() (string, error) {
	validationJSON, err := json.Marshal(vr)
	if err != nil {
		return "", err
	}
	return "data:application/json;base64," + base64.StdEncoding.EncodeToString(validationJSON), nil
}

func (vr *ValidationReport) UploadToIpfs(addr common.Address, gateway_url string, jwt string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var rak []byte
	var err error
	if vr.ReplicaInfo != nil {
		rak, err = vr.ReplicaInfo.RAK.MarshalText()
		if err != nil {
			return "", err
		}
	}

	part, err := writer.CreateFormFile("file", fmt.Sprintf("validation-report-%s-%s-%s.json", vr.AppId, rak, addr.Hex()))
	if err != nil {
		return "", err
	}

	validationJSON, err := json.Marshal(vr)
	if err != nil {
		return "", err
	}

	_, err = part.Write(validationJSON)
	if err != nil {
		return "", err
	}

	err = writer.WriteField("network", "public")
	if err != nil {
		return "", err
	}

	err = writer.WriteField("name", fmt.Sprintf("Validation report for app %s and replica %s by %s", vr.AppId, rak, addr.Hex()))
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", gateway_url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("server returned error code %d: %s",
			resp.StatusCode, string(bodyBytes))
	}

	var response UploadResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ipfs://%s", response.Data.Cid), nil
}

const ScoreReplicaNotExpiredYet = "replica-not-expired-yet"
const ScoreAtLeastOneReplica = "at-least-one-replica"

var (
	ScoreTable = map[string]uint8{
		ScoreReplicaNotExpiredYet: 50,
		ScoreAtLeastOneReplica:    50,
	}

	TagValidationInProgress = [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	TagValidationFinished   = [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
)

type ValidationMonitor struct {
	name       string
	vnCfg      *config.ValidationNetworkConfig
	ipfsCfg    *config.IpfsConfig
	signingKey []byte
	roflClient *rofl.RoflClient
	logger     *logging.Logger

	client             *ethclient.Client
	identityRegistry   *abi.IdentityRegistry
	validationRegistry *abi.ValidationRegistry
}

func NewValidationMonitor(name string, vnCfg *config.ValidationNetworkConfig, ipfsCfg *config.IpfsConfig, signingKey []byte, roflClient *rofl.RoflClient) *ValidationMonitor {
	return &ValidationMonitor{
		name:       name,
		vnCfg:      vnCfg,
		ipfsCfg:    ipfsCfg,
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
	vm.client, err = ethclient.DialContext(ctx, vm.vnCfg.RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer vm.client.Close()

	validationRegistryAddress := common.HexToAddress(vm.vnCfg.ValidationRegistry)
	vm.validationRegistry, err = abi.NewValidationRegistry(validationRegistryAddress, vm.client)
	if err != nil {
		vm.logger.Error("failed to parse validator registry", "err", err)
		return fmt.Errorf("failed to parse validator registry: %w", err)
	}

	identityRegistryAddress := common.HexToAddress(vm.vnCfg.IdentityRegistry)
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

	vm.logger.Debug("fetching unfinished validation requests", "validator_address", validatorAddress)
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	validatorRequestHashes, err := vm.validationRegistry.GetValidatorRequests(callOpts, validatorAddress)
	if err != nil {
		return fmt.Errorf("failed to get unfinished validation requests: %w", err)
	}
	validatorRequestHashesHex := make([]string, len(validatorRequestHashes))
	for i, rh := range validatorRequestHashes {
		validatorRequestHashesHex[i] = fmt.Sprintf("%x", rh)
	}
	vm.logger.Debug("unfinished validation requests fetched", "count", len(validatorRequestHashes), "request_hashes", validatorRequestHashesHex)

	// TODO: Binary search.
	for _, rh := range validatorRequestHashes {
		validationStatus, err := vm.validationRegistry.GetValidationStatus(callOpts, rh)
		if err != nil {
			vm.logger.Error("failed to get unfinished validation request", "req_hash", fmt.Sprintf("0x%x", rh), "err", err)
			continue
		}
		if bytes.Equal(validationStatus.Tag[:], TagValidationFinished[:]) {
			// Validation request already processed.
			continue
		}
		report := vm.generateValidationResponse(ctx, validationStatus.AgentId)
		txHash, err := vm.sendValidationResponse(rh, report)
		if err != nil {
			reportJson, _ := json.Marshal(report)
			vm.logger.Info("unable to send validation response", "report", reportJson, "request_hash", fmt.Sprintf("0x%x", rh), "agent_id", report.AgentId, "app_id", report.AppId, "tx_hash", txHash, "err", err)
		} else {
			vm.logger.Debug("validation response sent", "request_hash", fmt.Sprintf("%x", rh), "agent_id", report.AgentId, "app_id", report.AppId, "tx_hash", txHash)
		}
	}

	f, err := abi.NewValidationRegistryFilterer(validationRegistryAddress, vm.client)
	if err != nil {
		return fmt.Errorf("failed to create filterer: %w", err)
	}

	logs := make(chan *abi.ValidationRegistryValidationRequest)

	opts := &bind.WatchOpts{
		Context: ctx,
	}
	sub, err := f.WatchValidationRequest(opts, logs, []common.Address{validatorAddress}, []*big.Int{}, [][32]byte{})
	//sub, err := f.WatchValidationRequest(opts, logs, []common.Address{}, []*big.Int{}, [][32]byte{}) // Fetch all events.
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
				"agentId", event.AgentId,
				"requestURI", event.RequestUri,
			)

			if event.ValidatorAddress != validatorAddress {
				vm.logger.Info("validator address mismatch. Ignoring.", "event_validator_address", event.ValidatorAddress, "wanted_validator_address", validatorAddress)
				continue
			}

			report := vm.generateValidationResponse(ctx, event.AgentId)
			txHash, err := vm.sendValidationResponse(event.RequestHash, report)
			if err != nil {
				reportJson, _ := json.Marshal(report)
				vm.logger.Info("unable to send validation response", "report", reportJson, "request_hash", fmt.Sprintf("0x%x", event.RequestHash), "agent_id", report.AgentId, "app_id", report.AppId, "tx_hash", txHash, "err", err)
			} else {
				vm.logger.Debug("validation response sent", "request_hash", fmt.Sprintf("0x%x", event.RequestHash), "agent_id", report.AgentId, "app_id", report.AppId, "tx_hash", txHash)
			}
		}
	}
}

func (vm *ValidationMonitor) generateValidationResponse(ctx context.Context, agentId *big.Int) *ValidationReport {
	report := NewValidationReport(agentId)

	rawAppId, err := vm.identityRegistry.GetMetadata(&bind.CallOpts{Context: ctx}, agentId, IDENTITY_REGISTRY_METADATA_APP_ID_KEY)
	if err != nil {
		report.Error = fmt.Errorf("unable to get app ID for agent %s: %w", agentId, err)
		return report
	}
	if err = report.AppId.UnmarshalText(rawAppId); err != nil {
		report.Error = fmt.Errorf("unable to unmarshal text for app %s and agent %s: %w", string(rawAppId), agentId, err)
		return report
	}

	appInfo, replicas, lastEpoch, err := vm.roflClient.GetRoflReplicas(string(rawAppId))
	if err != nil {
		report.Error = fmt.Errorf("unable to get ROFL info for app %s and agent %s from Oasis chain: %w", string(rawAppId), agentId, err)
		return report
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
		vm.logger.Debug("ROFL has no replicas", "agentId", agentId, "appId", string(rawAppId))
	}

	return report
}

func (vm *ValidationMonitor) sendValidationResponse(requestHash [32]byte, report *ValidationReport) (common.Hash, error) {
	addr, auth, err := GetAddrAuth(vm.signingKey, vm.client)
	if err != nil {
		return common.Hash{}, err
	}

	reportURI := ""
	if vm.ipfsCfg.GatewayUrl != "" {
		reportURI, err = report.UploadToIpfs(addr, vm.ipfsCfg.GatewayUrl, vm.ipfsCfg.Jwt)
		vm.logger.Info("validation response stored to IPFS", "uri", reportURI)
	} else {
		reportURI, err = report.ToInlineURI()
	}
	if err != nil {
		return common.Hash{}, err
	}

	reportURIHash := crypto.Keccak256([]byte(reportURI))
	auth.GasLimit = 200000 // XXX: Gas estimation is off by ~3000 when sending tx the second time (bug in go-ethereum lib?). It takes ~112k gas for an IPFS-stored validation response.
	tx, err := vm.validationRegistry.ValidationResponse(auth, requestHash, report.Score, reportURI, [32]byte(reportURIHash), TagValidationFinished)
	if err != nil {
		if tx != nil {
			return tx.Hash(), err
		} else {
			return common.Hash{}, err
		}
	}

	var receipt *types.Receipt
	for i := 0; i < 300; i++ {
		time.Sleep(1000 * time.Millisecond)
		receipt, err = vm.client.TransactionReceipt(auth.Context, tx.Hash())
		if err == nil {
			break
		}
	}

	if receipt == nil {
		return common.Hash{}, fmt.Errorf("unable to get the receipt")
	}

	if receipt.Status != 1 {
		return receipt.TxHash, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return receipt.TxHash, nil
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

	return fromAddress, auth, nil
}
