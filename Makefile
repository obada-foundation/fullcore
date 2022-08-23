PROJECT = obada/fullcore
CI_COMMIT_REF_SLUG ?= develop
CONTAINER_IMAGE = $(PROJECT):$(CI_COMMIT_REF_SLUG)
CONTAINER_TAG_IMAGE = $(PROJECT):$(CI_COMMIT_TAG)
COMMIT_HASH = $(CI_COMMIT_SHA)
SHELL := /bin/sh
SRC_DIR = $(CURDIR)/src
protoVer = v0.7
protoImageName = tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen = cosmos-sdk-proto-gen-$(protoVer)
.DEFAULT_GOAL := help

docker/build:
	docker build -t $(CONTAINER_IMAGE) -f docker/Dockerfile .

docker/publish:
	docker push $(CONTAINER_IMAGE)

docker: docker/build docker/publish

proto: proto/format proto/gen

proto/format:

proto/gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(SRC_DIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

src/protogen:
	cd src && starport generate proto-go

src/run:
	cd src && go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	cd src && go mod tidy && go mod vendor

