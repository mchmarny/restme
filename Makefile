SERVICE_NAME     ?=restme
RELEASE_VERSION  ?=v0.6.28
SERVICE_URL      ?=https://restme.cloudylab.dev
KO_DOCKER_REPO   ?=gcr.io/cloudy-labs

all: help

name: ## Outputs service name
	@echo $(SERVICE_NAME)
.PHONY: name

version: ## Outputs current verison
	@echo $(RELEASE_VERSION)
.PHONY: version

tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor
.PHONY: tidy

upgrade: ## Upgrades all dependancies 
	go get -u ./...
	go mod tidy
	go mod vendor
.PHONY: upgrade

test: ## Runs tests on the entire project 
	go test -count=1 -race -covermode=atomic -coverprofile=cover.out ./...
.PHONY: test

lint: ## Lints the entire project 
	golangci-lint -c .golangci.yaml run --timeout=3m
.PHONY: lint

run: ## Runs uncompiled Go service code
	LOG_LEVEL=debug ADDRESS=":8080" go run ./cmd/main.go
.PHONY: run

verify: ## Runs verification test against the running service
	test/version $(SERVICE_URL) $(RELEASE_VERSION)
	test/endpoints $(SERVICE_URL)
.PHONY: verify

image: ## Creates container image using ko
	KO_DOCKER_REPO=$(KO_DOCKER_REPO)/$(SERVICE_NAME) \
	GOFLAGS="-ldflags=-X=main.version=$(RELEASE_VERSION)" \
		ko publish ./cmd/ --bare --tags $(RELEASE_VERSION),latest
	COSIGN_PASSWORD="" cosign sign \
		--key cosign.key \
		-a tag=$(RELEASE_VERSION) \
		$(KO_DOCKER_REPO)/$(SERVICE_NAME)
	COSIGN_PASSWORD="" cosign verify \
		--key cosign.pub \
		$(KO_DOCKER_REPO)/$(SERVICE_NAME)
.PHONY: image

tag: ## Creates release tag 
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)
.PHONY: tag

clean: ## Cleans bin and temp directories
	go clean
	rm -fr ./vendor
	rm -fr ./bin
.PHONY: clean

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help