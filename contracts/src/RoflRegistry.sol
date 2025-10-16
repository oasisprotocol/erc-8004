// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Subcall } from "@oasisprotocol/sapphire-contracts/contracts/Subcall.sol";

contract RoflRegistry {
    event KeyAdded(bytes21, string, bytes32);

    mapping(bytes => bytes32) private _keys;

    function addKey(bytes32 pubkey, string calldata context) external {
        _keys[abi.encodePacked(Subcall.getRoflAppId(), context)] = pubkey;

        emit KeyAdded(Subcall.getRoflAppId(), context, pubkey);
    }
    function getKey(bytes21 appId, string calldata context) external view returns (bytes32) {
        return _keys[abi.encodePacked(appId, context)];
    }

    function hasKey(bytes21 appId, string calldata context) external view returns (bool) {
        return _keys[abi.encodePacked(appId, context)] != bytes32(0);
    }
}
