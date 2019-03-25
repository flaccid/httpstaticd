DOCKER_REGISTRY = index.docker.io
IMAGE_NAME = httpstaticd
IMAGE_VERSION = latest
IMAGE_ORG = flaccid
IMAGE_TAG = $(DOCKER_REGISTRY)/$(IMAGE_ORG)/$(IMAGE_NAME):$(IMAGE_VERSION)

WORKING_DIR := $(shell pwd)

.DEFAULT_GOAL := docker-build

.PHONY: docker-release

run:: ## runs the main program with go
		@go run cmd/httpstaticd/httpstaticd.go

build:: ## builds the main program with go
		@go build -o bin/httpstaticd cmd/httpstaticd/httpstaticd.go 

docker-build:: ## builds the docker image locally
		@docker build --pull \
		-t $(IMAGE_TAG) $(WORKING_DIR)

docker-run:: ## runs the docker image locally
		@docker run \
			-it \
			$(DOCKER_REGISTRY)/$(IMAGE_ORG)/$(IMAGE_NAME):$(IMAGE_VERSION)

docker-push:: ## pushes the docker image to the registry
		@docker push $(IMAGE_TAG)

docker-release:: docker-build docker-push ## builds and pushes the docker image to the registry

# A help target including self-documenting targets (see the awk statement)
define HELP_TEXT
Usage: make [TARGET]... [MAKEVAR1=SOMETHING]...

Available targets:
endef
export HELP_TEXT
help: ## this help target
	@echo httpsr
	@echo
	@echo "$$HELP_TEXT"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / \
		{printf "\033[36m%-30s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)
