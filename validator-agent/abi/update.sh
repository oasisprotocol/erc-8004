#!/bin/sh

# Extract ABI and contract bytecode from foundry-generated artifacts.
for c in IdentityRegistry ReputationRegistry ValidationRegistry; do
  jq .bytecode.object ../tests/contracts/out/${c}.sol/${c}.json > ${c}.bin
  # Trim leading "0x and trailing "
  sed -i 's/^.\{3\}\(.*\).$/\1/' ${c}.bin

  jq .abi ../tests/contracts/out/${c}.sol/${c}.json > ${c}.abi

  abigen --bin=${c}.bin --abi=${c}.abi --pkg=abi --type ${c} --out=${c}.go
done
