import os
import time

from erc8004 import Erc8004
from eth_account import Account
import logging
from rofl_metadata import RoflMetadata, ERC8004_AGENT_ID_KEY

def main():
    logging.basicConfig(
        level=logging.INFO,
    )

    metadata_handler = RoflMetadata()
    metadata = metadata_handler.client.get_metadata()

    private_key = os.getenv("PRIVATE_KEY")
    if not private_key:
        private_key_id = os.getenv("PRIVATE_KEY_ID", "erc8004_signing_key")
        private_key = metadata_handler.client.generate_key(private_key_id)

    address = Account.from_key(private_key).address
    logging.info(f"Using ERC-8004 account {address}")

    erc8004 = Erc8004(
        chain_id=int(os.getenv("CHAIN_ID", "11155111")),  # Ethereum Sepolia testnet
        rpc_url=os.getenv("RPC_URL"),
        signer_key=private_key,
        ipfs_type="pinata",  # Options: "pinata", "filecoinPin", "node"
        jwt=os.getenv("PINATA_JWT")
    )

    while not erc8004.sdk.web3_client.get_balance(address):
        logging.info(f"Account {address} has no balance. Please top it up in order to register a new ERC-8004 agent and submit the validation request.")
        time.sleep(5)

    # Send ValidationRequest to ERC-8004 registry.
    agent_id = os.getenv("AGENT_ID")
    if agent_id is None:
        agent_id = metadata.get(ERC8004_AGENT_ID_KEY, None)
    if agent_id is None:
        name = os.getenv("AGENT_NAME", "rofl-agent")
        description = os.getenv("AGENT_DESCRIPTION", "Oasis ROFL-powered trustless agent.")
        image = os.getenv("AGENT_IMAGE")
        version = os.getenv("AGENT_VERSION", "0.1.0")
        category = os.getenv("AGENT_CATEGORY", "rofl")
        app_id = metadata_handler.client.get_app_id()
        agent_id = erc8004.register_agent(app_id, name, description, image, version, category)

    # Update ROFL replica metadata with exposed public keys.
    key_ids = os.getenv("PUBLIC_KEY_IDS")

    # Store any new keys or agent ID.
    metadata_handler.store_public_keys(key_ids.split(",") if key_ids else [], metadata)

    erc8004.validation_request(agent_id, os.getenv("VALIDATOR_ADDRESS", "0x7b60d730B6FBb55F9e4D1DB33B2D476e9c19fd35"))

if __name__ == "__main__":
    main()
