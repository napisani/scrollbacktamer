VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -X 'main.Version=$(VERSION)' \
           -X 'main.GitCommit=$(GIT_COMMIT)' \
           -X 'main.BuildDate=$(BUILD_DATE)'

all: build

build:
	go build -ldflags "$(LDFLAGS)" -o scrollbacktamer .

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -f scrollbacktamer coverage.out coverage.html

install:
	go install -ldflags "$(LDFLAGS)" .

# Recompute flake.nix's vendorHash after a go.mod/go.sum change (see
# scripts/update-vendor-hash.sh for how).
update-vendor-hash:
	./scripts/update-vendor-hash.sh

.PHONY: all build test test-coverage clean install update-vendor-hash