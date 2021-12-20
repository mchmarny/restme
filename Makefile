SERVICE_NAME     ?=restme
RELEASE_VERSION  ?=v0.6.14
TARGET_REGISTRY  ?=gcr.io/cloudy-lab
SERVICE_URL      ?=https://restme.cloudylab.dev

all: help

version: ## Outputs current verison
	@echo $(RELEASE_VERSION)
.PHONY: version

url: ## Outputs service url
	@echo $(shell gcloud run services describe restme --region us-west1 --format='value(status.url)')
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
	LOG_LEVEL=debug ADDRESS=":8080" go run ./cmd/main.go
.PHONY: run

verify: ## Runs verification test against the running service
	test/version $(SERVICE_URL) $(RELEASE_VERSION)
	test/endpoints $(SERVICE_URL)
.PHONY: verify

build: ## Compiles the Service code.
	CGO_ENABLED=0 go build -a -mod vendor -o bin/rester ./cmd/
.PHONY: build

upgrade: ## Upgrades all dependancies 
	go get -u ./...
	go mod tidy
	go mod vendor
.PHONY: upgrade

image: ## Creates container image using ko
	gcloud builds submit \
		--project $(PROJECT_ID) \
		--tag "$(TARGET_REGISTRY)/$(SERVICE_NAME):$(RELEASE_VERSION)"
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