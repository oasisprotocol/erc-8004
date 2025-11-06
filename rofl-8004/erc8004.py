from agent0_sdk import SDK, Agent
import base64
from eth_account import Account
import json
import logging

METADATA_APP_ID_KEY= "net.oasis.app_id"

class Erc8004:
    def __init__(self, rpc_url, signer_key, chain_id=11155111, ipfs_type="pinata", jwt=None):
        self.signer_key = signer_key
        self.chain_id = chain_id

        self.sdk = SDK(
            chainId=chain_id,
            rpcUrl=rpc_url,
            signer=signer_key,
            ipfs=ipfs_type,
            pinataJwt=jwt,
        )

    def register_agent(self, app_id: str, name: str, description: str, image: str, version: str, category: str) -> str:
        """
        Registers a new agent in IdentityRegistry.
        :return: agent ID without chain ID
        """
        # Create agent
        agent = self.sdk.createAgent(
            name=name,
            description=description,
            image=image,
        )
        #agent = sdk.loadAgent("11155111:1143")

        # Configure endpoints (automatically extracts capabilities)
        ##agent.setMCP("https://mcp.example.com/")  # Extracts tools, prompts, resources
        ##agent.setA2A("https://a2a.example.com/agent-card.json")  # Extracts skills
        ##agent.setENS("myagent.eth")

        # Configure wallet and trust
        signer_address = Account.from_key(self.signer_key).address
        agent.setAgentWallet(signer_address, chainId=self.chain_id)
        agent.setTrust(reputation=True, cryptoEconomic=True, teeAttestation=True)

        # Add metadata and set status
        metadata = {METADATA_APP_ID_KEY: app_id}
        if version:
            metadata["version"] = version
        if category:
            metadata["category"] = category
        agent.setMetadata(metadata)
        agent.setActive(True)

        # Register on-chain with IPFS
        logging.info(f"Prepared agent metadata for submission {metadata}")
        agent.registerIPFS()
        logging.info(f"Agent registered: {agent.agentId}")
        logging.info(f"Agent URI: {agent.agentURI}")

        return agent.agentId

    def validation_request(self, agent_id: str, validator_address: str):
        """
        Submit validation request to ValidationRegistry for an agent to a specific validator.

        :param agent_id: agent ID
        :param validator_address: hex-encoded validator address
        """
        logging.info(f"Preparing ValidationRequest for agent {agent_id}")

        if ":" in agent_id:
            token_id = int(agent_id.split(":")[-1])
        else:
            token_id = int(agent_id)

        app_id = self.sdk.web3_client.call_contract(
            self.sdk.identity_registry,
            "getMetadata",
            token_id,
            METADATA_APP_ID_KEY,
        )
        if app_id == "":
            logging.error(f"{METADATA_APP_ID_KEY} not in agent metadata. Aborting.")
            exit(-1)
        else:
            logging.info(f"got {METADATA_APP_ID_KEY}: {app_id}")

        addr = self.sdk.web3_client.account.address
        tx_count = self.sdk.web3_client.get_transaction_count(addr)
        nonce = self.sdk.web3_client.w3.keccak(text=addr + str(tx_count)).hex()
        vr_json = json.dumps({"nonce": nonce})
        vr_uri = f"data:application/json;base64,{base64.b64encode(vr_json.encode('utf-8'))}"

        tx_hash = self.sdk.web3_client.transact_contract(
            self.sdk.validation_registry,
            "validationRequest(address,uint256,string,bytes32)",
            validator_address,
            token_id,
            vr_uri,
            self.sdk.web3_client.w3.keccak(text=vr_uri)
        )

        logging.info(f"Tx hash: {tx_hash}")

        receipt = self.sdk.web3_client.w3.eth.wait_for_transaction_receipt(tx_hash)

        logging.info(f"Receipt: {receipt}")
        logging.info(f"ValidationRequest for agent {agent_id} successfully submitted")
