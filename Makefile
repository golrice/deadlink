# Makefile for DeadLink project

# Variables
BINARY_NAME := deadlink
GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test
GORUN := $(GO) run

# Default target
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/client/main.go

# Clean built binary
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Run the application directly
run:
	$(GORUN) cmd/client/main.go

# Install dependencies (if needed)
deps:
	$(GO) mod download
	$(GO) mod tidy

# Install the binary to GOPATH/bin
install:
	$(GO) install cmd/client/main.go

# Lint and vet code
lint:
	$(GO) vet ./...

.PHONY: all build clean test run deps install lint