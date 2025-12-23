import bech32
import cbor2
import logging
from ecdsa import SigningKey, SECP256k1, Ed25519
from oasis_rofl_client import RoflClient, KeyKind

ERC8004_SIGNING_KEY_ID = "erc8004_signing_key"
ERC8004_AGENT_ID_KEY = "erc8004_agent_id"

class RoflMetadata:
    def __init__(self):
        self.client = RoflClient()

    def store_public_keys(self, key_ids: list[str], metadata: dict[str,str] | None) -> dict[str, str]:
        pub_keys: dict[str, str] = {}

        for id_type in key_ids:
            key_id = id_type.split(':')[0]

            try:
                kind = KeyKind[id_type.split(':')[1]] if ":" in id_type else KeyKind.SECP256K1
                key = self.client.generate_key(key_id, kind)
                pub_keys[key_id] = RoflMetadata.derive_pubkey(bytes.fromhex(key), kind).hex()

            except Exception as e:
                logging.error(f"Failed to generate key {id_type}: {e}")
                continue

        for key, val in pub_keys.items():
            metadata[f"{key}.pk"] = val

        if metadata == self.client.get_metadata():
            logging.info(f"ROFL replica metadata unchanged. {metadata}")
        else:
            logging.info(f"Updating replica metadata. {metadata}")
            self.client.set_metadata(metadata)

        return metadata

    def get_app_metadata(self) -> dict[str, str]:
        """Returns app metadata stored on-chain"""

        app_id = self.client.get_app_id()
        _, bech32_data = bech32.bech32_decode(app_id)
        app_id_bytes = bytes(bech32.convertbits(bech32_data, 5, 8, False))

        response = self.client.query("rofl.App", cbor2.dumps({"id": app_id_bytes}))
        app_metadata = cbor2.loads(response)["metadata"]

        return {
            "name": app_metadata.get("net.oasis.rofl.name"),
            "version": app_metadata.get("net.oasis.rofl.version"),
            "repository": app_metadata.get("net.oasis.rofl.repository"),
            "author": app_metadata.get("net.oasis.rofl.author"),
            "license": app_metadata.get("net.oasis.rofl.license"),
            "homepage": app_metadata.get("net.oasis.rofl.homepage"),
            "description": app_metadata.get("net.oasis.rofl.description"),
        }

    def derive_pubkey(key: bytes, kind: KeyKind) -> bytes:
        match kind:
            case KeyKind.SECP256K1:
                sk = SigningKey.from_string(key, curve=SECP256k1)
                vk = sk.get_verifying_key()
                return vk.to_string()
            case KeyKind.ED25519:
                sk = SigningKey.from_string(key, curve=Ed25519)
                vk = sk.get_verifying_key()
                return vk.to_string()
            case _:
                raise Exception(f"Unsupported key type: {kind}")
