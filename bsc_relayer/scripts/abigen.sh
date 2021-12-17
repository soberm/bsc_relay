#!/bin/sh

cd "$(dirname "$0")" || exit 1

if ! which abigen >/dev/null; then
  echo "error: abigen not installed" >&2
  exit 1
fi

abigen --abi ../../bsc_contracts/artifacts/contracts/BSCRelay.sol/BSCRelay.abi --pkg relay --type BSCRelayContract --out ../pkg/relay/bscrelaycontract.go