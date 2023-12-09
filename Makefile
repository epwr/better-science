# Go variables
GO      := go
GOFMT   := gofmt
GOLINT  := golint
GOFILES := $(shell find . -name "*.go")
OUTPUT  := bin
DOCS	:= docs

# Main build target
all: build run

# Build the Go application
build:
	$(GO) build -o $(OUTPUT)/better_science ./cmd/better_science.go

# Run the Go application
run:
	./$(OUTPUT)/better_science

# Clean the build artifacts
clean:
	rm -f $(OUTPUT)/better_science

# Format the code using gofmt
fmt:
	$(GOFMT) -w $(GOFILES)

# Lint the code using golint
lint:
	$(GOLINT) ./...

# Run tests
test:
	$(GO) test ./...

# Run tests with verbose output
testv:
	$(GO) test -v ./...

# Run tests with coverage report
coverage:
	$(GO) test -cover ./...

# Install dependencies
deps:
	$(GO) mod download

# Perform a full check (format, lint, and test)
check: fmt lint test

# Build and run the application
build-run: build run

# Set up and run the local auto-doc server.
docs:
	golds ./...

# Re-create the docs 
gen-docs:
	golds -gen -dir=$(DOCS) -nouses ./...

.PHONY: all build run clean fmt lint test deps check build-run docs
