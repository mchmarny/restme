RELEASE_VERSION ?=$(shell cat version)

all: help

version: ## Outputs current verison
	@echo $(RELEASE_VERSION)
.PHONY: version

tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor
.PHONY: tidy

upgrade: ## Upgrades all dependancies 
	go get -d -u ./...
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
	CONFIG=configs/dev.json go run ./cmd/
.PHONY: run

binrun: ## Compile local version of the binary 
	CGO_ENABLED=0 go build \
		-ldflags "-X main.version=$(RELEASE_VERSION)" \
		-mod vendor -o ./bin/restme ./cmd/
	CONFIG=configs/dev.json ./bin/restme
.PHONY: binrun

tests: ## Runs verification test against the running service
	tests/version http://127.0.0.1:8080 $(RELEASE_VERSION)
	tests/endpoints http://127.0.0.1:8080
.PHONY: tests

infra1: ## Sets up developer loop (gh, gcr, oidc)
	terraform -chdir=infra/1-dev-flow apply -var-file=terraform.tfvars
.PHONY: infra1

infra2: ## Sets up serving (run, secret, policy, monitorng)
	terraform -chdir=infra/2-service apply -var-file=terraform.tfvars
.PHONY: infra2

token: ## Prints new dev token 
	tools/make-token dev-user
.PHONY: token

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