SERVICE_NAME     ?=restme
RELEASE_VERSION  ?=v0.0.4
KO_DOCKER_REPO   ?=ghcr.io/mchmarny

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
	golangci-lint -c .github/linters.yml run --timeout=3m
.PHONY: lint

run: ## Runs uncompiled Go code
	go run ./cmd/main.go
.PHONY: run

message: ## Invokes echo service 
	curl -i -H "Content-Type: application/json" \
		http://localhost:8080/v1/echo \
		-d '{ "on": 1620253683, "msg": "hellow" }'
.PHONY: message

upgrade: ## Upgrades all dependancies 
	go get -u ./...
	go mod tidy 
.PHONY: upgrade

image: ## Creates container image using ko
	KO_DOCKER_REPO=$(KO_DOCKER_REPO)/$(SERVICE_NAME) ko publish ./cmd/ --bare --tags $(RELEASE_VERSION),latest
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