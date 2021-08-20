SERVICE_NAME     ?=restme
RELEASE_VERSION  ?=v0.5.11
KO_DOCKER_REPO   ?=gcr.io/cloudy-lab
TEST_AUTH_TOKEN  ?=test/test.token
SERVICE_URL      :=$(shell gcloud run services describe restme --region us-west1 --format='value(status.url)')

all: help

version: ## Outputs current verison
	@echo $(RELEASE_VERSION)
.PHONY: version

url: ## Outputs service url
	@echo $(SERVICE_URL)
.PHONY: url

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

message: ## Invokes echo service 
	curl -i \
	     -H "Content-Type: application/json" \
	     -H "Authorization: Bearer $(shell cat $(TEST_AUTH_TOKEN))" \
		 $(SERVICE_URL)/v1/echo/message \
		 -d '{ "on": $(shell date +%s), "msg": "hello?" }'
.PHONY: message

request: ## Invokes request service 
	curl -i \
	     -H "Content-Type: application/json" \
	     -H "Authorization: Bearer $(shell cat $(TEST_AUTH_TOKEN))" \
		 $(SERVICE_URL)/v1/request/info
.PHONY: request

runtime: ## Invokes requesruntimet service 
	curl -i \
	     -H "Content-Type: application/json" \
	     -H "Authorization: Bearer $(shell cat $(TEST_AUTH_TOKEN))" \
		 $(SERVICE_URL)/v1/runtime/info
.PHONY: runtime

load: ## Invokes load service 
	curl -i \
	     -H "Content-Type: application/json" \
	     -H "Authorization: Bearer $(shell cat $(TEST_AUTH_TOKEN))" \
		 $(SERVICE_URL)/v1/load/cpu/30s
.PHONY: load

metrics: ## Collects metrics 
	curl $(SERVICE_URL)/metrics
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
	COSIGN_PASSWORD="" cosign sign \
		-key cosign.key \
		-a tag=$(RELEASE_VERSION) \
		$(KO_DOCKER_REPO)/$(SERVICE_NAME)
	COSIGN_PASSWORD="" cosign verify \
		-key cosign.pub \
		$(KO_DOCKER_REPO)/$(SERVICE_NAME)
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