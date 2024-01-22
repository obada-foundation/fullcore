PROJECT = obada/fullcore

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

BUILD_DIR ?= $(CURDIR)/build

CI_COMMIT_REF_SLUG ?= develop
CONTAINER_IMAGE = $(PROJECT):$(CI_COMMIT_REF_SLUG)
CONTAINER_TESTNET_IMAGE = $(CONTAINER_IMAGE)-testnet
CONTAINER_TAG_IMAGE = $(PROJECT):$(CI_COMMIT_TAG)
COMMIT_HASH = $(CI_COMMIT_SHA)
SHELL := /bin/sh
protoVer = 0.14.0
protoImageName = ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=docker run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)
containerProtoGen = cosmos-sdk-proto-gen-$(protoVer)
containerProtoFmt = cosmos-sdk-proto-fmt-$(protoVer)
.DEFAULT_GOAL := help

##################
###   Docker   ###
##################

docker/build:
	docker build -t $(CONTAINER_IMAGE) -f docker/Dockerfile .
	docker build -t $(CONTAINER_TESTNET_IMAGE) -f docker/testnet/Dockerfile .

docker/publish:
	docker push $(CONTAINER_IMAGE)
	docker push $(CONTAINER_TESTNET_IMAGE)

docker: docker/build docker/publish

##################
###     CI     ###
##################

ci/coverage:
	go test $(cd src && go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt

ci/test:
	go test -v ./... -cover

ci/lint:
	golangci-lint --config .golangci.yml run --print-issued-lines --out-format=github-actions ./...

##################
###  Protobuf  ###
##################

proto: proto/format proto/lint proto/gen proto/swagger-gen

proto/deps:
	go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm-proto@latest

proto/swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protocgen-docs.sh

proto/gen: proto/deps
	@echo "Generating protobuf files..."
	@$(protoImage) sh ./scripts/protocgen.sh
	@go mod tidy

proto/format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto/lint:
	@$(protoImage) buf lint proto/ --error-format=json

.PHONY: proto proto/gen proto/format proto/lint

##################
###   Others   ###
##################

src/run:
	go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	go mod tidy && go mod vendor

mockgen:
	./scripts/mockgen.sh

clean:
	rm -rf $(BUILD_DIR)

bin:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR) $(CURDIR)/cmd/fullcored
