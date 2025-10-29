package tests

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oasisprotocol/erc-8004/validator-agent/abi"
	"github.com/oasisprotocol/erc-8004/validator-agent/config"
	"github.com/oasisprotocol/erc-8004/validator-agent/logging"
	"github.com/oasisprotocol/erc-8004/validator-agent/rofl"
	"github.com/oasisprotocol/erc-8004/validator-agent/validator"
)

type RegistrationFile struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Endpoints   []struct {
		Name         string `json:"name"`
		Endpoint     string `json:"endpoint"`
		Version      string `json:"version,omitempty"`
		Capabilities struct {
		} `json:"capabilities,omitempty"`
	} `json:"endpoints"`
	Registrations []struct {
		AgentId       int    `json:"agentId"`
		AgentRegistry string `json:"agentRegistry"`
	} `json:"registrations"`
	SupportedTrust []string `json:"supportedTrust"`
}

func generateAgentRegistrationURI(identityRegistry string) (string, error) {
	registrationFile := RegistrationFile{
		Type:        "https://eips.ethereum.org/EIPS/eip-8004#registration-v1",
		Name:        "myAgentName",
		Description: "A natural language description of the Agent, which MAY include what it does, how it works, pricing, and interaction methods",
		Image:       "https://example.com/agentimage.png",
		Endpoints: []struct {
			Name         string `json:"name"`
			Endpoint     string `json:"endpoint"`
			Version      string `json:"version,omitempty"`
			Capabilities struct {
			} `json:"capabilities,omitempty"`
		}{
			{
				Name:     "A2A",
				Endpoint: "https://agent.example/.well-known/agent-card.json",
				Version:  "0.3.0",
			},
			{
				Name:         "MCP",
				Endpoint:     "https://mcp.agent.eth/",
				Capabilities: struct{}{},
				Version:      "2025-06-18",
			},
			{
				Name:     "OASF",
				Endpoint: "ipfs://{cid}",
				Version:  "0.7",
			},
			{
				Name:     "ENS",
				Endpoint: "vitalik.eth",
				Version:  "v1",
			},
			{
				Name:     "DID",
				Endpoint: "did:method:foobar",
				Version:  "v1",
			},
			{
				Name:     "agentWallet",
				Endpoint: "eip155:1:0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7",
			},
		},
		Registrations: []struct {
			AgentId       int    `json:"agentId"`
			AgentRegistry string `json:"agentRegistry"`
		}{
			{
				AgentId:       22,
				AgentRegistry: "eip155:31337:" + identityRegistry,
			},
		},
		SupportedTrust: []string{
			"reputation",
			"crypto-economic",
			"tee-attestation",
		},
	}

	registrationJSON, err := json.Marshal(registrationFile)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal registration file: %w", err)
	}

	return "data:application/json;base64," + base64.StdEncoding.EncodeToString(registrationJSON), nil
}

func generateValidationRequestURI() (string, error) {
	return "data:application/json;base64," + base64.StdEncoding.EncodeToString([]byte("{}")), nil
}

func registerAgent(t *testing.T, client *ethclient.Client, agentKey string, identityRegistry string) *big.Int {
	agentURI, err := generateAgentRegistrationURI(identityRegistry)
	if err != nil {
		t.Fatalf("generateAgentRegistrationURI failed: %s", err)
	}

	_, auth, err := validator.GetAddrAuth(common.FromHex(agentKey), client)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := abi.NewIdentityRegistry(common.HexToAddress(identityRegistry), client)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := instance.Register0(auth, agentURI, []abi.IIdentityRegistryMetadataEntry{
		{
			Key:   validator.IDENTITY_REGISTRY_METADATA_APP_ID_KEY,
			Value: []byte("rofl1qpykfkl6ea78cyy67d35f7fmpk3pg36vashka4v9"),
		},
	})
	if err != nil {
		t.Fatalf("Failed to send tx: %s", err)
	}

	var receipt *types.Receipt
	for {
		receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if receipt.Status != 1 {
		t.Fatalf("Transaction failed with status: %d", receipt.Status)
	}

	// Parse the Registered event to extract agentId
	if len(receipt.Logs) == 0 {
		t.Fatalf("No logs found in transaction receipt")
	}

	for _, l := range receipt.Logs {
		if len(l.Topics) >= 4 {
			return new(big.Int).SetBytes(l.Topics[3].Bytes()) // agent id
		}
	}

	t.Fatal("Failed to extract agentId from transaction logs")
	return nil
}

func sendValidationRequest(t *testing.T, client *ethclient.Client, agentId *big.Int, agentKey string, validatorAddress string, validationRegistry string) [32]byte {
	_, auth, err := validator.GetAddrAuth(common.FromHex(agentKey), client)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := abi.NewValidationRegistry(common.HexToAddress(validationRegistry), client)
	if err != nil {
		t.Fatal(err)
	}

	validationRequestURI, err := generateValidationRequestURI()
	if err != nil {
		t.Fatal(err)
	}
	validationRequestURIHash := crypto.Keccak256([]byte(validationRequestURI))

	tx, err := instance.ValidationRequest(auth, common.HexToAddress(validatorAddress), agentId, validationRequestURI, [32]byte(validationRequestURIHash))
	if err != nil {
		t.Fatalf("Failed to send tx: %s", err)
	}

	var receipt *types.Receipt
	for {
		receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if receipt.Status != 1 {
		t.Fatalf("Transaction failed with status: %d", receipt.Status)
	}

	return [32]byte(validationRequestURIHash)
}

func waitForValidationResponse(t *testing.T, client *ethclient.Client, validatorAddress common.Address, validationRegistry string, agentId *big.Int, requestHash [32]byte) {
	instance, err := abi.NewValidationRegistry(common.HexToAddress(validationRegistry), client)
	if err != nil {
		t.Fatal(err)
	}

	opts := &bind.WatchOpts{Context: t.Context()}
	logs := make(chan *abi.ValidationRegistryValidationResponse)
	sub, err := instance.WatchValidationResponse(opts, logs, []common.Address{validatorAddress}, []*big.Int{agentId}, [][32]byte{requestHash})
	if err != nil {
		t.Fatal("failed to watch ValidationRequest events", "err", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-t.Context().Done():
			return
		case err := <-sub.Err():
			t.Fatal("subscription error", "err", err)
		case event := <-logs:
			t.Logf("received validation response: score=%d", event.Response)
			if event.Response != 100 {
				t.Fatalf("Expected validation score to be 100, got %d", event.Response)
			}
			return
		case <-time.After(30 * time.Second):
			t.Fatal("timeout waiting for validation response")
		}
	}
}

// TestE2EValidatorAgent tests the end-to-end flow of the validator agent
func TestE2EValidatorAgent(t *testing.T) {
	t.Run("TestValidationMonitor", func(t *testing.T) {
		err := logging.InitLogging(nil)
		if err != nil {
			t.Fatal(err)
		}

		signingKey := "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
		cfg := config.ValidationNetworkConfig{
			RPC:                "ws://localhost:8545",
			IdentityRegistry:   "0x5FbDB2315678afecb367f032d93F642f64180aa3",
			ValidationRegistry: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		}
		roflClient := rofl.NewRoflClient(&config.RoflConfig{Network: "mainnet"})

		vm := validator.NewValidationMonitor("localnet-monitor", &cfg, common.FromHex(signingKey), roflClient)

		go vm.Start(context.Background())

		client, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			t.Fatalf("Failed to connect to Ethereum client: %s", err)
		}
		defer client.Close()

		agentKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		agentId := registerAgent(t, client, agentKey, cfg.IdentityRegistry)
		t.Log(agentId)

		validatorAddress, _, err := validator.GetAddrAuth(common.FromHex(signingKey), client)
		if err != nil {
			t.Fatal(err)
		}
		requestHash := sendValidationRequest(t, client, agentId, agentKey, validatorAddress.Hex(), cfg.ValidationRegistry)
		waitForValidationResponse(t, client, validatorAddress, cfg.ValidationRegistry, agentId, requestHash)
	})
}
