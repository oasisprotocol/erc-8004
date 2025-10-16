// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Subcall } from "@oasisprotocol/sapphire-contracts/contracts/Subcall.sol";

contract RoflRegistry {
    event KeyAdded(bytes21, string, address);

    mapping(bytes => address) private _keys;

    function addKey(address pubkey, string calldata context) external {
        _keys[abi.encodePacked(Subcall.getRoflAppId(), context)] = pubkey;

        emit KeyAdded(Subcall.getRoflAppId(), context, pubkey);
    }
    function getKey(bytes21 appId, string calldata context) external view returns (address) {
        return _keys[abi.encodePacked(appId, context)];
    }

    function hasKey(bytes21 appId, string calldata context) external view returns (bool) {
        return _keys[abi.encodePacked(appId, context)] != address(0);
    }
}
