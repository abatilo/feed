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

tmp/asdf-installs: .tool-versions
	@-mkdir -p $(@D)
	@-asdf plugin-add golang  || asdf install golang
	@-touch $@

tmp/bootstrap: tmp/asdf-installs
	@-mkdir -p $(@D)
	@-touch $@

.PHONY: clean
clean: ## Delete local dev environment
	@-rm -rf tmp

dist/feed: tmp/bootstrap go.mod go.sum
	@-mkdir -p $(@D)
	go mod vendor -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ -ldflags="-w -s" main.go

.PHONY: build
build: dist/feed ## Build the binary

.PHONY: run
run: ## Run the binary
	@-go mod vendor
	go run main.go
