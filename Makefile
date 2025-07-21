.PHONY: generate test test-integration test-unit clean help deps

# Default target
help:
	@echo "Available targets:"
	@echo "  generate         Generate Go models from OpenAPI specification"
	@echo "  test             Run all tests (unit + integration if API token is set)"
	@echo "  test-unit        Run only unit tests"
	@echo "  test-integration Run only integration tests (requires TCA_API_TOKEN)"
	@echo "  clean            Clean generated files"
	@echo "  deps             Install dependencies"
	@echo "  help             Show this help"

# Generate Go models from OpenAPI specification
generate:
	@echo "🚀 Generating models from The Companies API OpenAPI spec..."
	./generate.sh

# Run all tests
test:
	@echo "🧪 Running all tests..."
	go test -v .

# Run only unit tests (excluding integration tests)
test-unit:
	@echo "🧪 Running unit tests..."
	go test -v . -short

# Run only integration tests (requires API token)
test-integration:
	@echo "🌐 Running integration tests..."
	@if [ -z "$$TCA_API_TOKEN" ] && [ ! -f .env ]; then \
		echo "❌ TCA_API_TOKEN not found. Please:"; \
		echo "   1. Create a .env file with TCA_API_TOKEN=your_token"; \
		echo "   2. Or set TCA_API_TOKEN environment variable"; \
		echo "   Get your token from: https://www.thecompaniesapi.com/"; \
		exit 1; \
	fi
	go test -v . -run "TestIntegration"

# Clean generated files
clean:
	@echo "🧹 Cleaning generated files..."
	rm -f generated.go openapi-3.1.json openapi-3.0.json

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy 
