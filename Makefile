.PHONY: help test test-unit test-coverage lint fmt vet validate build clean

help: ## Show available tasks
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests
	go test ./... -v

test-unit: ## Run unit tests (short mode)
	go test ./... -v -short

test-coverage: ## Run tests with coverage report
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run golangci-lint
	golangci-lint run

fmt: ## Format code with gofmt
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

validate: fmt vet lint test ## Run all checks (pre-commit)

build: ## Build the project
	go build ./...

clean: ## Clean build artifacts
	rm -f coverage.out coverage.html

