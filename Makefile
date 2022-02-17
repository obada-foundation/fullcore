PROJECT = obada/fullcore
CI_COMMIT_REF_SLUG ?= develop
CONTAINER_IMAGE = $(PROJECT):$(CI_COMMIT_REF_SLUG)
CONTAINER_TAG_IMAGE = $(PROJECT):$(CI_COMMIT_TAG)
UI_CONTAINER_IMAGE =  $(PROJECT)-ui:$(CI_COMMIT_REF_SLUG)
UI_CONTAINER_TAG_IMAGE = $(PROJECT)-ui:$(CI_COMMIT_TAG)
COMMIT_HASH = $(CI_COMMIT_SHA)
SHELL := /bin/sh
.DEFAULT_GOAL := help


docker/node/build:
	docker build -t $(CONTAINER_IMAGE) -f docker/node/Dockerfile .

docker/node/publish:
	docker push $(CONTAINER_IMAGE)

docker/node: docker/node/build docker/node/publish

docker/ui/build:
	docker build -t $(UI_CONTAINER_IMAGE) -f docker/ui/Dockerfile .

docker/ui/publish:
	docker push $(UI_CONTAINER_IMAGE)

docker/ui: docker/ui/build docker/ui/publish

docker: docker/node docker/ui

src/protogen:
	cd src && starport generate proto-go

src/run:
	cd src && go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	cd src && go mod tidy && go mod vendor

