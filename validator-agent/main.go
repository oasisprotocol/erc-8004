package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	coreLogging "github.com/oasisprotocol/oasis-core/go/common/logging"
	"github.com/spf13/cobra"

	"github.com/oasisprotocol/erc-8004/validator-agent/appd"
	"github.com/oasisprotocol/erc-8004/validator-agent/config"
	"github.com/oasisprotocol/erc-8004/validator-agent/logging"
	"github.com/oasisprotocol/erc-8004/validator-agent/rofl"
	"github.com/oasisprotocol/erc-8004/validator-agent/validator"
)

var (
	// Path to the configuration file.
	configFile string

	// Oasis-web3-gateway root command.
	rootCmd = &cobra.Command{
		Use:     "validator-agent",
		Short:   "validator-agent",
		Version: Software,
		RunE:    exec,
	}
)

func initVersions() {
	cobra.AddTemplateFunc("toolchain", func() interface{} { return Toolchain })
	cobra.AddTemplateFunc("sdk", func() interface{} { return GetOasisSDKVersion() })

	rootCmd.SetVersionTemplate(`Software version: {{.Version}}
Oasis SDK version: {{ sdk }}
Go toolchain version: {{ toolchain }}
`)
}

func init() {
	initVersions()

	rootCmd.Flags().StringVar(&configFile, "config", "config.yaml", "path to the config.yml file")
}

func exec(_ *cobra.Command, _ []string) error {
	if err := runRoot(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func main() {
	_ = rootCmd.Execute()
}

//nolint:gocyclo
func runRoot() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize server config
	cfg, err := config.InitConfig(configFile)
	if err != nil {
		return err
	}
	// Initialize logging.
	if err = logging.InitLogging(cfg.Log); err != nil {
		return err
	}
	logger := coreLogging.GetLogger("main")

	appdClient := appd.NewAppdClient(cfg.Appd)
	sk, err := appdClient.GenerateKey(cfg.Appd.SigningKeyId, "secp256k1")
	if err != nil {
		return fmt.Errorf("failed to generate key for validator agent %w", err)
	}
	{
		skEcdsa, err := crypto.ToECDSA(sk)
		if err != nil {
			return fmt.Errorf("failed to convert key to ECDSA: %w", err)
		}
		validatorAddress := crypto.PubkeyToAddress(skEcdsa.PublicKey)
		logger.Info("derived validator signing key", "address", validatorAddress.Hex())

		skEcdsaBase64 := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&skEcdsa.PublicKey))
		err = appdClient.AddMetadata(cfg.Appd.SigningKeyId, skEcdsaBase64)
		if err != nil {
			return fmt.Errorf("failed to store public key %s to replica metadata: %w", skEcdsaBase64, err)
		}
		logger.Debug("updated ROFL metadata with public signing key", "key_id", cfg.Appd.SigningKeyId, "pubkey", crypto.FromECDSAPub(&skEcdsa.PublicKey))
	}

	roflClient := rofl.NewRoflClient(cfg.Rofl)

	monitors := make(map[string]*validator.ValidationMonitor)
	for name, n := range cfg.ValidationNetworks {
		m := validator.NewValidationMonitor(name, n, sk, roflClient)
		monitors[name] = m
	}

	// Wait for all monitors to finish
	var wg sync.WaitGroup
	for _, m := range monitors {
		wg.Go(func() {
			m.Start(context.Background())
		})
	}
	wg.Wait()

	return nil
}
