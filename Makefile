.PHONY: generate test clean help

# Default target
help:
	@echo "Available targets:"
	@echo "  generate    Generate Go models from OpenAPI specification"
	@echo "  test        Run tests"
	@echo "  clean       Clean generated files"
	@echo "  help        Show this help"

# Generate Go models from OpenAPI specification
generate:
	@echo "🚀 Generating models from The Companies API OpenAPI spec..."
	./generate.sh

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v .

# Clean generated files
clean:
	@echo "🧹 Cleaning generated files..."
	rm -f models.go openapi-3.1.json openapi-3.0.json

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy 
