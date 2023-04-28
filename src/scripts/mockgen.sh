#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/obit/types/expected_keepers.go -package testutil -destination x/obit/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=vendor/github.com/cosmos/cosmos-sdk/x/nft/expected_keepers.go -package nft -destination testutil/keeper/nft/keepers.go
