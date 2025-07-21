#!/bin/bash

# Generate models from The Companies API OpenAPI specification
# This script handles the OpenAPI 3.1 -> 3.0 downgrade and generates Go models

set -e

echo "🔄 Downloading OpenAPI specification..."
curl -s https://api.thecompaniesapi.com/v2/openapi -o openapi-3.1.json

echo "📦 Installing openapi-down-convert..."
if ! command -v npx &> /dev/null; then
    echo "❌ Error: npx is required but not installed. Please install Node.js and npm."
    exit 1
fi

echo "⬇️  Converting OpenAPI 3.1 to 3.0..."
npx --yes @apiture/openapi-down-convert@latest --input openapi-3.1.json --output openapi-3.0.json

echo "🔧 Generating Go client and models..."
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen.yaml openapi-3.0.json

echo "🧹 Cleaning up temporary files..."
rm openapi-3.1.json openapi-3.0.json

echo "✅ Successfully generated client and models!"
echo "💡 The generated code includes both types and client methods." 

