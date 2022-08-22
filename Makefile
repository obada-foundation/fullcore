PROJECT = obada/fullcore
CI_COMMIT_REF_SLUG ?= develop
CONTAINER_IMAGE = $(PROJECT):$(CI_COMMIT_REF_SLUG)
CONTAINER_TAG_IMAGE = $(PROJECT):$(CI_COMMIT_TAG)
COMMIT_HASH = $(CI_COMMIT_SHA)
SHELL := /bin/sh
.DEFAULT_GOAL := help

docker/build:
	docker build -t $(CONTAINER_IMAGE) -f docker/node/Dockerfile .

docker/publish:
	docker push $(CONTAINER_IMAGE)

docker: docker/build docker/publish

src/protogen:
	cd src && starport generate proto-go

src/run:
	cd src && go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	cd src && go mod tidy && go mod vendor

