import json
import os
import time
import sys

from erc8004 import Erc8004
from eth_account import Account
import logging
from rofl_metadata import RoflMetadata, ERC8004_SIGNING_KEY_ID, ERC8004_AGENT_ID_KEY
from oasis_rofl_client import KeyKind

ROFL_8004_CONFIG_FILENAME = ".rofl-8004/rofl-8004.json"

def main():
    logging.basicConfig(
        level=logging.INFO,
    )

    rpc_url = os.getenv("RPC_URL") or None
    if rpc_url is None:
        sys.exit("RPC_URL is required")

    jwt = os.getenv("PINATA_JWT") or None
    if jwt is None:
        sys.exit("PINATA_JWT is required")

    metadata_handler = RoflMetadata()
    replica_metadata = metadata_handler.client.get_metadata()

    private_key_id = os.getenv("ERC8004_SIGNING_KEY_ID", ERC8004_SIGNING_KEY_ID)
    private_key = os.getenv("ERC8004_SIGNING_KEY") or None
    if not private_key:
        private_key = metadata_handler.client.generate_key(private_key_id)

    address = Account.from_key(private_key).address
    logging.info(f"Using ERC-8004 account {address}")

    erc8004 = Erc8004(
        chain_id=int(os.getenv("CHAIN_ID", "11155111")),  # Ethereum Sepolia testnet
        rpc_url=rpc_url,
        signer_key=private_key,
        ipfs_type="pinata",  # Options: "pinata", "filecoinPin", "node"
        jwt=jwt,
    )

    while not erc8004.sdk.web3_client.get_balance(address):
        logging.info(f"Account {address} has no balance. Please top it up in order to register a new ERC-8004 agent and submit the validation request.")
        time.sleep(5)

    # Check for ROFL config on persistent storage.
    rofl_config = {}
    if os.path.exists(ROFL_8004_CONFIG_FILENAME):
        with open(ROFL_8004_CONFIG_FILENAME, "r") as f:
            rofl_config = json.load(f)
            logging.info(f"Loaded ROFL config file from {ROFL_8004_CONFIG_FILENAME}")

    agent_id = os.getenv("AGENT_ID") or None
    if agent_id is None:
        agent_id = rofl_config.get("agent_id")
    if agent_id is None:
        agent_id = replica_metadata.get(ERC8004_AGENT_ID_KEY)
    if agent_id is None:
        app_metadata = metadata_handler.get_app_metadata()
        name = os.getenv("AGENT_NAME") or app_metadata["name"] or "rofl-agent"
        description = os.getenv("AGENT_DESCRIPTION") or app_metadata["description"] or "Oasis ROFL-powered trustless agent."
        image = os.getenv("AGENT_IMAGE") or None
        version = os.getenv("AGENT_VERSION") or app_metadata["version"] or "0.1.0"
        category = os.getenv("AGENT_CATEGORY") or "rofl"
        mcp = os.getenv("AGENT_MCP") or None
        a2a = os.getenv("AGENT_A2A") or None
        ens = os.getenv("AGENT_ENS") or None
        app_id = metadata_handler.client.get_app_id()
        agent_id = erc8004.register_agent(app_id, name, description, image, version, category, mcp, a2a, ens)

    # Store agent ID to disk.
    rofl_config["agent_id"] = agent_id
    os.makedirs(os.path.dirname(ROFL_8004_CONFIG_FILENAME), exist_ok=True)
    with open(ROFL_8004_CONFIG_FILENAME, "w") as f:
        json.dump(rofl_config, f)
        logging.info(f"Stored ROFL config file to {ROFL_8004_CONFIG_FILENAME}")

    # Store agent ID to ROFL metadata.
    replica_metadata[ERC8004_AGENT_ID_KEY] = agent_id

    # Store the public key of the ERC-8004 signing key (either provided or TEE-derived) to ROFL metadata.
    replica_metadata[f"{private_key_id}.pk"] = RoflMetadata.derive_pubkey(bytes.fromhex(private_key), KeyKind.SECP256K1).hex()

    # Update ROFL replica metadata with exposed public keys.
    key_ids = os.getenv("PUBLIC_KEY_IDS")

    # Store any new keys or agent ID.
    metadata_handler.store_public_keys(key_ids.split(",") if key_ids else [], replica_metadata)

    # Send ValidationRequest to ERC-8004 registry.
    erc8004.validation_request(agent_id, os.getenv("VALIDATOR_ADDRESS", "0x7b60d730B6FBb55F9e4D1DB33B2D476e9c19fd35"))

if __name__ == "__main__":
    main()
