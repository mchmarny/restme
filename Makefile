SERVICE_NAME     ?=restme
RELEASE_VERSION  ?=v0.5.6
KO_DOCKER_REPO   ?=ghcr.io/mchmarny
TEST_AUTH_TOKEN  ?=test/test.token

all: help

version: ## Outputs current verison
	@echo $(RELEASE_VERSION)
.PHONY: version

tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor
.PHONY: tidy

test: ## Runs tests on the entire project 
	go test -count=1 -race -covermode=atomic -coverprofile=cover.out ./...
.PHONY: test

lint: ## Lints the entire project 
	golangci-lint -c .golangci.yaml run --timeout=3m
.PHONY: lint

run: ## Runs uncompiled Go service code
	LOG_LEVEL=debug ADDRESS=":8080" KEY_FILE="test/test.key" go run ./cmd/service/main.go
.PHONY: run

cli: ## Compiles the CLI code.
	CGO_ENABLED=0 \
	GOFLAGS="-ldflags=-X=main.appVersion=$(RELEASE_VERSION)" \
		go build -a -mod vendor -o bin/restme-cli ./cmd/cli/
	bin/restme-cli --help
.PHONY: cli

token: ## Runs uncompiled Go CLI to generate token
	bin/restme-cli token create \
		--secret test/test.key \
		--issuer test \
		--email demo@domain.com \
		--ttl "30s"
.PHONY: token

verify: ## Runs verification test against the running service
	AUTH_TOKEN="$(shell cat $(TEST_AUTH_TOKEN))" test/endpoints "http://localhost:8080"
.PHONY: verify

curl-message: ## Invokes echo service 
	curl -i \
	     -H "Content-Type: application/json" \
	     -H "Authorization: Bearer $(shell cat $(TEST_AUTH_TOKEN))" \
		 http://localhost:8080/v1/echo/message \
		 -d '{ "on": $(shell date +%s), "msg": "hello?" }'
.PHONY: message

metrics: ## Collects metrics 
	curl http://localhost:8080/metrics
.PHONY: metrics

build: ## Compiles the Service code.
	CGO_ENABLED=0 go build -a -mod vendor -o bin/rester ./cmd/service/
	GIN_MODE=release LOG_JSON=true bin/restme
.PHONY: build

upgrade: ## Upgrades all dependancies 
	go get -u ./...
	go mod tidy 
.PHONY: upgrade

image: ## Creates container image using ko
	KO_DOCKER_REPO=$(KO_DOCKER_REPO)/$(SERVICE_NAME) \
	GOFLAGS="-ldflags=-X=main.version=$(RELEASE_VERSION)" \
		ko publish ./cmd/service/ --bare --tags $(RELEASE_VERSION),latest
.PHONY: tag

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