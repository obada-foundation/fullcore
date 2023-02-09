#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/obit/types/expected_keepers.go -package testutil -destination x/obit/testutil/expected_keepers_mocks.go
