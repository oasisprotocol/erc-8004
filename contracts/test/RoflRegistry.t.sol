// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../src/RoflRegistry.sol";
import {Test, console} from "forge-std/Test.sol";
import {SapphireTest} from "@oasisprotocol-sapphire-foundry-0.1.2/BaseSapphireTest.sol";

contract RoflRegistryTest is SapphireTest {
    RoflRegistry public registry;

    address testPubkey = address(0x1234567890123456789012345678901234567890);
    string testContext = "test-context";
    bytes21 testAppId = bytes21(0x000000000000000000000000000000000000000000);

    function setUp() public override {
        super.setUp();
        registry = new RoflRegistry();
    }

    function testAddHasGetKey() public {
        assertTrue(registry.hasKey(testAppId, testContext)==false);

        registry.addKey(testPubkey, testContext);

        address pk = registry.getKey(testAppId, testContext);
        assertEq(pk, testPubkey);
        assertTrue(registry.hasKey(testAppId, testContext));
    }

    function testAddKeyMultipleContexts() public {
        address testPubkey2 = address(0x9876543210987654321098765432109876543210);
        string memory context2 = "another-context";

        registry.addKey(testPubkey, testContext);
        registry.addKey(testPubkey2, context2);

        address pk1 = registry.getKey(testAppId, testContext);
        address pk2 = registry.getKey(testAppId, context2);

        assertEq(pk1, testPubkey);
        assertEq(pk2, testPubkey2);
    }

    function testOverwriteKey() public {
        address newPubkey = address(0x1111111111111111111111111111111111111111);

        registry.addKey(testPubkey, testContext);
        registry.addKey(newPubkey, testContext);

        address key = registry.getKey(testAppId, testContext);
        assertEq(key, newPubkey);
    }
}