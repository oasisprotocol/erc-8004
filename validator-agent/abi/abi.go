package abi

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type MetadataEntry struct {
	Key   string
	Value []byte
}

//go:embed IdentityRegistry.abi
var IdentityRegistryAbiRaw string

//go:embed ReputationRegistry.abi
var ReputationRegistryAbiRaw string

//go:embed ValidationRegistry.abi
var ValidationRegistryAbiRaw string

var IdentityRegistryAbi abi.ABI
var ReputationRegistryAbi abi.ABI
var ValidationRegistryAbi abi.ABI

func init() {
	var err error
	if IdentityRegistryAbi, err = abi.JSON(strings.NewReader(IdentityRegistryAbiRaw)); err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}
	if ReputationRegistryAbi, err = abi.JSON(strings.NewReader(ReputationRegistryAbiRaw)); err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}
	if ValidationRegistryAbi, err = abi.JSON(strings.NewReader(ValidationRegistryAbiRaw)); err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}
}
