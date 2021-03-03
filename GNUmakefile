GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
USER := retr0h
HOSTNAME := github.com
NAME := terrable
#TAG := $(shell git describe --abbrev=0 --tags)
VERSION := 1.0
PLUGIN_NAME := terraform-provider-$(NAME)
TF_PLUGIN_PATH := $(HOME)/.terraform.d/plugins/$(HOSTNAME)/$(USER)/$(NAME)/$(VERSION)/$(GOOS)_$(GOARCH)

default: build

.PHONY: build
build: mod
	@go build -o build/$(GOOS)_$(GOARCH)/$(PLUGIN_NAME)_v$(VERSION)

.PHONY: install
install: build
	install -d $(TF_PLUGIN_PATH) && \
	install build/$(GOOS)_$(GOARCH)/$(PLUGIN_NAME)_v$(VERSION) $(TF_PLUGIN_PATH)

.PHONY: test
test:
	@go test -v -cover ./...

.PHONY: clean
clean:
	@rm -rf build/

.PHONY: mod
mod:
	@go mod tidy
	@go mod vendor
