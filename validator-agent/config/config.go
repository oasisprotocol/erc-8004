package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/rawbytes"

	"github.com/oasisprotocol/oasis-core/go/common/logging"
)

// Config contains the CLI configuration.
type Config struct {
	Rofl               *RoflConfig                         `koanf:"rofl"`
	Log                *LogConfig                          `koanf:"log"`
	ValidationNetworks map[string]*ValidationNetworkConfig `koanf:"validation_networks"`
	Appd               *AppdConfig                         `koanf:"appd"`
}

// Validate performs config validation.
func (cfg *Config) Validate() error {
	if cfg.Log != nil {
		if err := cfg.Log.Validate(); err != nil {
			return err
		}
	}
	if cfg.ValidationNetworks != nil {
		for name, n := range cfg.ValidationNetworks {
			if err := n.Validate(); err != nil {
				return fmt.Errorf("validation_network %s: %w", name, err)
			}
		}
	}
	return nil
}

type RoflConfig struct {
	Network string `koanf:"network"`
	RPC     string `koanf:"rpc"`
}

func (cfg *RoflConfig) Validate() error {
	if cfg.Network != "mainnet" && cfg.Network != "testnet" && cfg.Network != "localnet" {
		return fmt.Errorf("network %s invalid. Must be mainnet, testnet or localnet", cfg.Network)
	}
	return nil
}

// LogConfig contains the logging configuration.
type LogConfig struct {
	Format string `koanf:"format"`
	Level  string `koanf:"level"`
	File   string `koanf:"file"`
}

// Validate validates the logging configuration.
func (cfg *LogConfig) Validate() error {
	var format logging.Format
	if err := format.Set(cfg.Format); err != nil {
		return err
	}
	var level logging.Level
	return level.Set(cfg.Level)
}

// ValidationNetworksConfig is the configuration for EVM networks to be monitored.
type ValidationNetworkConfig struct {
	RPC                string `koanf:"rpc"`
	IdentityRegistry   string `koanf:"identity_registry"`
	ValidationRegistry string `koanf:"validation_registry"`
}

// Validate validates the logging configuration.
func (cfg *ValidationNetworkConfig) Validate() error {
	if cfg.RPC == "" {
		return fmt.Errorf("rpc is required")
	}
	if !common.IsHexAddress(cfg.IdentityRegistry) {
		return fmt.Errorf("identity_registry must be a valid Ethereum address")
	}
	if !common.IsHexAddress(cfg.ValidationRegistry) {
		return fmt.Errorf("validation_registry must be a valid Ethereum address")
	}
	return nil
}

type AppdConfig struct {
	// SigningKeyId is the key ID used when calling appd to generate a keypair.
	SigningKeyId string `koanf:"signing_key_id"`

	// Mock will use mocked version of appd. Useful for testing.
	Mock bool `koanf:"mock"`

	// Signing key to use in mocked version.
	MockSigningKey string `koanf:"mock_signing_key"`
}

// Validate validates the logging configuration.
func (cfg *AppdConfig) Validate() error {
	if cfg.SigningKeyId == "" {
		return fmt.Errorf("signing_key_id is required")
	}
	return nil
}

// InitConfig initializes configuration from file.
func InitConfig(f string) (*Config, error) {
	var config Config
	k := koanf.New(".")

	fileContent, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	// Find all ${...} patterns in config and replace them with environment variables.
	re := regexp.MustCompile(`\$\{([^}]+)\}`)
	expandedContent := re.ReplaceAllStringFunc(string(fileContent), func(match string) string {
		// Extract the variable name (remove ${ and })
		varName := match[2 : len(match)-1]
		// Get the environment variable value
		if val := os.Getenv(varName); val != "" {
			return val
		}
		// If not found, return the original match
		return match
	})

	if err := k.Load(rawbytes.Provider([]byte(expandedContent)), yaml.Parser()); err != nil {
		return nil, err
	}

	// Load environment variables and merge into the loaded config.
	if err := k.Load(env.Provider("", ".", func(s string) string {
		// `__` is used as a hierarchy delimiter.
		return strings.ReplaceAll(strings.ToLower(s), "__", ".")
	}), nil); err != nil {
		return nil, err
	}

	// Unmarshal into config.
	if err := k.Unmarshal("", &config); err != nil {
		return nil, err
	}

	// Validate config.
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}
