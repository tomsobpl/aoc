SHELL := /bin/bash

# Project settings
GO ?= go
OUT_DIR := out
LIB_DIR := lib
PLUGINS_DIR := $(LIB_DIR)/plugins
CLI_PKG := ./cmd/aoccli
CLI_NAME := aoccli

# Detect local OS/ARCH for suffix
UNAME_S := $(shell uname -s | tr '[:upper:]' '[:lower:]')
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
  ARCH := amd64
else ifeq ($(UNAME_M),aarch64)
  ARCH := arm64
else ifeq ($(UNAME_M),arm64)
  ARCH := arm64
else
  ARCH := $(UNAME_M)
endif
OS := $(UNAME_S)
SUFFIX := -$(OS)-$(ARCH)

# Versioning (can be overridden: make build VERSION=1.2.3)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE    ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Read minimum Go version from go.mod (informational)
GO_MIN_VERSION := $(shell sed -n 's/^go //p' go.mod)

# ldflags injected into cmd/aoccli (variables defined in cmd/aoccli/version.go)
LDFLAGS := -X 'main.Version=$(VERSION)' -X 'main.Commit=$(COMMIT)' -X 'main.Date=$(DATE)'

CLI_OUT := $(OUT_DIR)/$(CLI_NAME)$(SUFFIX)

.PHONY: all build plugins plugins-test test clean info

all: build

info:
	@echo "Go minimum version: $(GO_MIN_VERSION)"
	@echo "Building for: $(OS)/$(ARCH)"
	@echo "Version: $(VERSION) Commit: $(COMMIT) Date: $(DATE)"

$(OUT_DIR):
	@mkdir -p $(OUT_DIR)

$(PLUGINS_DIR):
	@mkdir -p $(PLUGINS_DIR)

build: info $(OUT_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(CLI_OUT) $(CLI_PKG)
	@echo "Built $(CLI_OUT)"

# Build all solution plugins found under plugins/solution/year*/day*/ using go plugin mode
plugins: $(PLUGINS_DIR)
	@set -euo pipefail; \
	shopt -s nullglob; \
  for dir in plugins/solution/*/day*/; do \
    if [[ -d "$$dir" ]]; then \
      year=$$(basename $$(dirname "$$dir")); \
      day=$$(basename "$$dir"); \
      out="$(PLUGINS_DIR)/year$${year#year}-day$${day#day}-plugin.so"; \
      echo "Building plugin $$dir -> $$out"; \
      $(GO) build -buildmode=plugin -o "$$out" "./$$dir"; \
    fi; \
  done
	@echo "Plugins built into $(PLUGINS_DIR)"

# Test all plugins by running `go test` in each plugin directory and aggregate results
plugins-test:
	@set -euo pipefail; \
	shopt -s nullglob; \
	fail=0; pass=0; total=0; \
	for dir in plugins/solution/year*/day*/; do \
	  if [[ -d "$$dir" ]]; then \
	    echo "===> Testing $$dir"; \
	    if (cd "$$dir" && $(GO) test ./...); then \
	      echo "PASS $$dir"; pass=$$((pass+1)); \
	    else \
	      echo "FAIL $$dir"; fail=$$((fail+1)); \
	    fi; \
	    total=$$((total+1)); \
	  fi; \
	done; \
	if [[ $$total -eq 0 ]]; then echo "No plugins found under plugins/solution"; fi; \
	echo "--- Plugin tests summary: total=$$total, pass=$$pass, fail=$$fail ---"; \
	if [[ $$fail -gt 0 ]]; then exit 1; fi

test:
	$(GO) test ./...

clean:
	rm -rf $(OUT_DIR) $(LIB_DIR)
	@echo "Cleaned $(OUT_DIR) and $(LIB_DIR)"
