SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

PROJECT_NAME = feed
DIST_DIR = "dist"

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Delete local dev environment
	@-rm -rf tmp

tmp/asdf-installs: .tool-versions
	@-mkdir -p $(@D)
	@-asdf plugin-add golang  || asdf install golang
	@-touch $@

tmp/go-installs: tools/tools.go
	@-mkdir -p $(@D)
	./tools/install.sh
	@-touch $@

tmp/bootstrap: tmp/asdf-installs tmp/go-installs
	@-mkdir -p $(@D)
	@-touch $@

.PHONY: generate
generate: tmp/go-installs
	go mod vendor
	go generate ./...

dist/feed: tmp/bootstrap generate
	@-mkdir -p $(@D)
	go mod vendor
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ -ldflags="-w -s" main.go

.PHONY: build
build: dist/feed ## Build the binary
	docker build -t $(PROJECT_NAME) -f build/Dockerfile .

.PHONY: test
test: dist/feed ## Run tests
	go test -race -covermode=atomic -coverpkg ./... ./...

.PHONY: run
run: build ## Run the binary
	docker run --rm -it $(PROJECT_NAME)
