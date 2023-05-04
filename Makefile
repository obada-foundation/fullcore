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
protoVer = v0.7
protoImageName = tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen = cosmos-sdk-proto-gen-$(protoVer)
containerProtoFmt = cosmos-sdk-proto-fmt-$(protoVer)
.DEFAULT_GOAL := help

docker/build:
	docker build -t $(CONTAINER_IMAGE) -f docker/Dockerfile .
	docker build -t $(CONTAINER_TESTNET_IMAGE) -f docker/testnet/Dockerfile .

docker/publish:
	docker push $(CONTAINER_IMAGE)
	docker push $(CONTAINER_TESTNET_IMAGE)

docker: docker/build docker/publish

ci/coverage:
	go test $(cd src && go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt

ci/test:
	go test -v ./... -cover

ci/lint:
	golangci-lint --config .golangci.yml run --print-issued-lines --out-format=github-actions ./...

proto: proto/format proto/gen

proto/format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find .  -name "*.proto" -not -path "./third_party/*" -exec clang-format -i {} \; ; fi

proto/gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

src/run:
	go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	go mod tidy && go mod vendor

mockgen:
	./scripts/mockgen.sh

swagger: proto-swagger-gen
.PHONY: swagger

clean:
	rm -rf $(BUILD_DIR)

bin:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR) $(CURDIR)/cmd/fullcored
