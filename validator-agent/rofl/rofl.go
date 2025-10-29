package rofl

import (
	"context"
	"encoding/json"
	"fmt"

	beaconApi "github.com/oasisprotocol/oasis-core/go/beacon/api"
	"github.com/oasisprotocol/oasis-core/go/common/logging"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/client"
	sdkConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/connection"
	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/modules/rofl"

	"github.com/oasisprotocol/erc-8004/validator-agent/config"
)

type RoflClient struct {
	cfg    *config.RoflConfig
	logger *logging.Logger
}

func NewRoflClient(cfg *config.RoflConfig) *RoflClient {
	return &RoflClient{cfg: cfg, logger: logging.GetLogger("rofl")}
}

func (c *RoflClient) GetRoflReplicas(appId string) (*rofl.AppConfig, []*rofl.Registration, beaconApi.EpochTime, error) {
	networkName := c.cfg.Network

	// Setup client.
	ctx := context.Background()
	network := sdkConfig.DefaultNetworks.All[networkName]
	if network == nil {
		return nil, nil, 0, fmt.Errorf("network '%s' not found in configuration", networkName)
	}
	// Override network config with custom address if provided.
	if c.cfg.RPC != "" {
		// Clone the network config to avoid modifying the default.
		networkCopy := *network
		network = &networkCopy

		// Set the custom RPC address. The SDK's connection.Connect will automatically
		// determine whether to use TLS based on whether the address is local.
		// For Unix socket paths (starting with /), prepend "unix://" scheme.
		if len(c.cfg.RPC) > 0 && c.cfg.RPC[0] == '/' {
			network.RPC = "unix://" + c.cfg.RPC
		} else {
			network.RPC = c.cfg.RPC
		}
		c.logger.Debug("Connecting...", "network", networkName, "rpc", c.cfg.RPC)
	} else {
		c.logger.Debug("Connecting...", "network", networkName)
	}

	conn, err := connection.Connect(ctx, network)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("failed to connect: %w", err)
	}
	consensusBackend := conn.Consensus().Core()

	// Fetch and display latest block.
	block, err := consensusBackend.GetBlock(ctx, consensus.HeightLatest)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("failed to fetch latest block: %w", err)
	}
	c.logger.Debug("Latest Block Information", "height", block.Height, "hash", block.Hash, "time", block.Time)

	// Setup ROFL client.
	sapphirePT := network.ParaTimes.All["sapphire"]
	if sapphirePT == nil {
		return nil, nil, 0, fmt.Errorf("sapphire paratime not found in network configuration")
	}
	runtimeClient := conn.Runtime(sapphirePT)
	c.logger.Debug("Runtime client initialized", "runtime_id", sapphirePT.ID)
	roflClient := rofl.NewV1(runtimeClient)

	c.logger.Debug("Querying Sapphire ROFL App", "appId", appId)

	// Query ROFL apps.
	appID := rofl.NewAppIDFromBech32(appId)
	app, err := roflClient.App(ctx, client.RoundLatest, appID)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Failed to fetch ROFL app %s: %w", appId, err)
	}
	appJson, _ := json.Marshal(app)
	c.logger.Debug("Got App", "app", appJson)

	// Query instances for this app.
	instances, err := roflClient.AppInstances(ctx, client.RoundLatest, app.ID)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Failed to fetch replicas: %w\n", err)
	}
	instancesJson, _ := json.Marshal(instances)
	c.logger.Debug("Got App instances", "instances", instancesJson)

	// Query for last epoch
	lastEpoch, err := conn.Consensus().Beacon().GetEpoch(ctx, consensus.HeightLatest)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Failed to fetch last epoch: %w\n", err)
	}
	c.logger.Debug("Got last epoch", "epoch", lastEpoch)

	return app, instances, lastEpoch, nil
}
