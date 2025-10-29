package appd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasisprotocol/oasis-core/go/common/logging"

	"github.com/oasisprotocol/erc-8004/validator-agent/config"
)

type AppdInterface interface {
	// Generate a secret key from keyId and kind.
	GenerateKey(keyId string, kind string) ([]byte, error)

	// Add a key-value entry to replica metadata.
	AddMetadata(key string, value string) error
}

type AppdClient struct {
	client *http.Client

	logger *logging.Logger
	cfg    *config.AppdConfig
}

func (c *AppdClient) GenerateKey(keyId string, kind string) ([]byte, error) {
	// Make request to generate keys endpoint
	jsonData := fmt.Sprintf(`{"key_id":"%s","kind":"%s"}`, keyId, kind)
	resp, err := c.client.Post("http://localhost/rofl/v1/keys/generate", "application/json", strings.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	// Parse JSON response
	var result struct {
		Key string `json:"key"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %w", err)
	}

	return common.FromHex(result.Key), nil
}

func (c *AppdClient) AddMetadata(keyId string, value string) error {
	resp, err := c.client.Get("http://localhost/rofl/v1/metadata")
	if err != nil {
		return fmt.Errorf("failed making get request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed reading get response: %w", err)
	}
	metadata := make(map[string]string)
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	metadata[keyId] = value
	metadataJson, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w\n", err)
	}
	resp, err = c.client.Post("http://localhost/rofl/v1/metadata", "application/json", strings.NewReader(string(metadataJson)))
	if err != nil {
		return fmt.Errorf("failed to make port request: %w\n", err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed reading set response: %w\n", err)
	}

	return nil
}

type AppdMock struct {
	logger *logging.Logger
	cfg    *config.AppdConfig
}

func (c *AppdMock) GenerateKey(_ string, _ string) ([]byte, error) {
	return common.FromHex(c.cfg.MockSigningKey), nil
}

func (c *AppdMock) AddMetadata(_ string, _ string) error {
	return nil
}

func NewAppdClient(cfg *config.AppdConfig) AppdInterface {
	logger := logging.GetLogger("appd")
	if cfg.Mock {
		return &AppdMock{
			logger: logger,
			cfg:    cfg,
		}
	}

	// Create HTTP client with Unix socket transport
	return &AppdClient{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
					var d net.Dialer
					return d.DialContext(ctx, "unix", "/run/rofl-appd.sock")
				},
			},
			Timeout: 10 * time.Second,
		},
		logger: logger,
		cfg:    cfg,
	}
}
