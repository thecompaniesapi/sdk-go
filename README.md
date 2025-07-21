# The Companies API SDK for Go

[![GoDoc](https://godoc.org/github.com/thecompaniesapi/sdk-go?status.svg)](https://godoc.org/github.com/thecompaniesapi/sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/thecompaniesapi/sdk-go)](https://goreportcard.com/report/github.com/thecompaniesapi/sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fully-featured Go SDK for [The Companies API](https://www.thecompaniesapi.com), providing type-safe access to company data, locations, industries, technologies, job titles, lists, and more.

If you need more details about a specific endpoint, you can find the corresponding documentation in [the API reference](https://www.thecompaniesapi.com/api).

You can also contact us on our [livechat](https://www.thecompaniesapi.com/) if you have any questions.

## 🚀 Features

- **Type-safe API client** with full model types from our [OpenAPI](https://api.thecompaniesapi.com/v2/openapi) schema
- **Powerful search capabilities** with filters, sorting and pagination
- **Real-time company enrichment** with both synchronous and asynchronous options
- **Create and manage** your company lists
- **Track and monitor** enrichment actions and requests
- **Generate detailed analytics** and insights for searches and lists
- **Natural language querying** for structured company information
- **Lightweight** with minimal dependencies
- **Context support** for all operations
- **Comprehensive error handling** with detailed error responses

## 📦 Installation

```bash
go get github.com/thecompaniesapi/sdk-go
```

## 🔑 Prerequisites

- Go 1.19 or higher
- Valid API key from [The Companies API](https://www.thecompaniesapi.com/)

## 🚀 Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    tca "github.com/thecompaniesapi/sdk-go"
)

func main() {
    // Initialize the client
    client, err := tca.ApiClient("your-api-key-here")
    if err != nil {
        log.Fatal(err)
    }
    
    // Search for technology companies
    page := float32(1)
    size := float32(10)
    search := "technology"
    
    response, err := client.SearchCompanies(context.Background(), &tca.SearchCompaniesParams{
        Page:   &page,
        Size:   &size,
        Search: &search,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    if response.JSON200 != nil {
        fmt.Printf("Found %d companies\n", response.JSON200.Meta.Total)
        for _, company := range response.JSON200.Companies {
            if company.About != nil && company.About.Name != nil {
                fmt.Printf("- %s\n", *company.About.Name)
            }
        }
    }
}
```

## 🏢 Companies

### Search companies

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/searchCompanies)

```go
// Search companies with basic filters
page := float32(1)
size := float32(10)
search := "artificial intelligence"

response, err := client.SearchCompanies(ctx, &tca.SearchCompaniesParams{
    Page:   &page,
    Size:   &size,
    Search: &search,
})

companies := response.JSON200.Companies // Companies that match the search
meta := response.JSON200.Meta          // Meta information (pagination, etc.)
```

```go
// Advanced search with complex query conditions
response, err := client.SearchCompaniesPost(ctx, tca.SearchCompaniesPostJSONRequestBody{
    Query: &[]tca.QueryCondition{
        {
            Attribute: "about.industries",
            Operator:  "or",
            Sign:      "equals",
            Values:    []string{"artificial-intelligence", "machine-learning"},
        },
        {
            Attribute: "locations.headquarters.country.code",
            Operator:  "and",
            Sign:      "equals",
            Values:    []string{"us"},
        },
    },
    Page: tca.Float32Ptr(1),
    Size: tca.Float32Ptr(20),
})

companies := response.JSON200.Companies // Companies matching the query
meta := response.JSON200.Meta          // Meta information
```

### Search companies by name

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/searchCompaniesByName)

```go
// Search companies by their name
name := "microsoft"
response, err := client.SearchCompaniesByName(ctx, &tca.SearchCompaniesByNameParams{
    Name: &name,
})

companies := response.JSON200.Companies // Companies with matching names
```

### Search companies by prompt

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/searchCompaniesByPrompt)

```go
// Use natural language to find companies
prompt := "Find me SaaS companies in San Francisco with more than 100 employees"
response, err := client.SearchCompaniesByPrompt(ctx, &tca.SearchCompaniesByPromptParams{
    Prompt: &prompt,
})

companies := response.JSON200.Companies // Companies matching the prompt
```

### Find similar companies

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/searchSimilarCompanies)

```go
// Find companies similar to given domains
domains := "apple.com,microsoft.com"
response, err := client.SearchSimilarCompanies(ctx, &tca.SearchSimilarCompaniesParams{
    Domains: &domains,
})

companies := response.JSON200.Companies // Similar companies
```

### Count companies matching your query

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/countCompanies)

```go
// Count how many companies are in the computer-software industry
response, err := client.CountCompaniesPost(ctx, tca.CountCompaniesPostJSONRequestBody{
    Query: &[]tca.QueryCondition{
        {
            Attribute: "about.industries",
            Operator:  "or",
            Sign:      "equals",
            Values:    []string{"computer-software"},
        },
    },
})

count := response.JSON200.Data // Number of companies that match the query
```

### Enrich a company from a domain name

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompany)

```go
// Fetch company data from our database without enrichment (faster response)
response, err := client.FetchCompany(ctx, "microsoft.com", &tca.FetchCompanyParams{})

company := response.JSON200.Data // The company profile
```

```go
// Fetch company data and re-analyze it in real-time to get fresh, up-to-date information
refresh := true
response, err := client.FetchCompany(ctx, "microsoft.com", &tca.FetchCompanyParams{
    Refresh: &refresh,
})

company := response.JSON200.Data // The company profile (refreshed)
```

### Enrich a company from an email

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompanyByEmail)

🕹️ Enrich your users at signup with the latest information about their company

```go
// Fetch the company profile behind a professional email address
email := "jack@openai.com"
response, err := client.FetchCompanyByEmail(ctx, &tca.FetchCompanyByEmailParams{
    Email: &email,
})

company := response.JSON200.Data // The company profile
```

### Enrich a company from a social network URL

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompanyBySocial)

```go
// Fetch the company profile behind a social network URL
linkedin := "https://www.linkedin.com/company/apple"
response, err := client.FetchCompanyBySocial(ctx, &tca.FetchCompanyBySocialParams{
    Linkedin: &linkedin,
})

company := response.JSON200.Data // The company profile
```

### Find a company email patterns

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompanyEmailPatterns)

```go
// Fetch the company email patterns for a specific domain
response, err := client.FetchCompanyEmailPatterns(ctx, "apple.com", &tca.FetchCompanyEmailPatternsParams{})

patterns := response.JSON200.Data // The company email patterns
```

### Ask a question about a company

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/askCompany)

```go
// Ask what products a company offers using its domain
response, err := client.AskCompany(ctx, "microsoft.com", tca.AskCompanyJSONRequestBody{
    Question: "What products does this company offer?",
    Model:    tca.StringPtr("large"), // 'small' is also available
    Fields: &[]tca.QuestionField{
        {
            Key:         "products",
            Type:        "array|string",
            Description: tca.StringPtr("The products that the company offers"),
        },
    },
})

answer := response.JSON200.Data.Answer // Structured AI response
meta := response.JSON200.Data.Meta     // Meta information
```

### Fetch the context of a company

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompanyContext)

```go
// Get AI-generated strategic insights about a company
response, err := client.FetchCompanyContext(ctx, "microsoft.com")

context := response.JSON200.Data.Context // Includes market, model, differentiators, etc.
meta := response.JSON200.Data.Meta       // Meta information
```

### Fetch analytics data for a query or your lists

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/fetchCompaniesAnalytics)

```go
// Analyze company distribution by business type
attribute := "about.businessType"
response, err := client.FetchCompaniesAnalytics(ctx, &tca.FetchCompaniesAnalyticsParams{
    Attribute: &attribute,
    Query: &[]tca.QueryCondition{
        {
            Attribute: "locations.headquarters.country.code",
            Operator:  "or",
            Sign:      "equals",
            Values:    []string{"us", "gb", "fr"},
        },
    },
})

analytics := response.JSON200.Data.Data // Aggregated values
meta := response.JSON200.Data.Meta      // Meta information
```

### Export analytics data in multiple formats for a search

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Companies/operation/exportCompaniesAnalytics)

```go
// Export analytics to CSV
response, err := client.ExportCompaniesAnalytics(ctx, tca.ExportCompaniesAnalyticsJSONRequestBody{
    Format:     "csv",
    Attributes: []string{"about.industries", "about.totalEmployees"},
    Query: &[]tca.QueryCondition{
        {
            Attribute: "technologies.active",
            Operator:  "or",
            Sign:      "equals",
            Values:    []string{"shopify"},
        },
    },
})

analytics := response.JSON200.Data.Data // Aggregated values
meta := response.JSON200.Data.Meta      // Meta information
```

## 🎯 Actions

### Request an action on one or more companies

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Actions/operation/requestAction)

```go
// Request an enrichment job on multiple companies
estimate := false
response, err := client.RequestAction(ctx, tca.RequestActionJSONRequestBody{
    Domains:  []string{"microsoft.com", "apple.com"},
    Job:      "enrich-companies",
    Estimate: &estimate,
})

actions := response.JSON200.Data.Actions // Track this via FetchActions
meta := response.JSON200.Data.Meta       // Meta information
```

### Fetch the actions for your actions

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Actions/operation/fetchActions)

```go
// Fetch recent actions
status := "completed"
page := float32(1)
size := float32(5)

response, err := client.FetchActions(ctx, &tca.FetchActionsParams{
    Status: &status,
    Page:   &page,
    Size:   &size,
})

actions := response.JSON200.Data.Actions // Actions that match the query
meta := response.JSON200.Data.Meta       // Meta information
```

## 🏭 Industries

### Search industries

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Industries/operation/searchIndustries)

```go
// Search industries by keyword
search := "software"
size := float32(10)

response, err := client.SearchIndustries(ctx, &tca.SearchIndustriesParams{
    Search: &search,
    Size:   &size,
})

industries := response.JSON200.Data.Industries // Industries that match the keyword
meta := response.JSON200.Data.Meta             // Meta information
```

### Find similar industries

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Industries/operation/searchIndustriesSimilar)

```go
// Find industries similar to given ones
industries := "saas,fintech"
response, err := client.SearchIndustriesSimilar(ctx, &tca.SearchIndustriesSimilarParams{
    Industries: &industries,
})

similar := response.JSON200.Data.Industries // Industries that are similar to the given ones
meta := response.JSON200.Data.Meta          // Meta information
```

## ⚛️ Technologies

### Search technologies

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Technologies/operation/searchTechnologies)

```go
// Search technologies by keyword
search := "shopify"
size := float32(10)

response, err := client.SearchTechnologies(ctx, &tca.SearchTechnologiesParams{
    Search: &search,
    Size:   &size,
})

technologies := response.JSON200.Data.Technologies // Technologies that match the keyword
meta := response.JSON200.Data.Meta                 // Meta information
```

## 🌍 Locations

### Search cities

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Locations/operation/searchCities)

```go
// Search cities by name
search := "new york"
size := float32(5)

response, err := client.SearchCities(ctx, &tca.SearchCitiesParams{
    Search: &search,
    Size:   &size,
})

cities := response.JSON200.Data.Cities // Cities that match the name
meta := response.JSON200.Data.Meta     // Meta information
```

### Search counties

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Locations/operation/searchCounties)

```go
// Search counties by name
search := "orange"
size := float32(5)

response, err := client.SearchCounties(ctx, &tca.SearchCountiesParams{
    Search: &search,
    Size:   &size,
})

counties := response.JSON200.Data.Counties // Counties that match the name
meta := response.JSON200.Data.Meta         // Meta information
```

### Search states

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Locations/operation/searchStates)

```go
// Search states by name
search := "california"
size := float32(5)

response, err := client.SearchStates(ctx, &tca.SearchStatesParams{
    Search: &search,
    Size:   &size,
})

states := response.JSON200.Data.States // States that match the name
meta := response.JSON200.Data.Meta     // Meta information
```

### Search countries

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Locations/operation/searchCountries)

```go
// Search countries by name
search := "france"
size := float32(5)

response, err := client.SearchCountries(ctx, &tca.SearchCountriesParams{
    Search: &search,
    Size:   &size,
})

countries := response.JSON200.Data.Countries // Countries that match the name
meta := response.JSON200.Data.Meta           // Meta information
```

### Search continents

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Locations/operation/searchContinents)

```go
// Search continents by name
search := "asia"
size := float32(5)

response, err := client.SearchContinents(ctx, &tca.SearchContinentsParams{
    Search: &search,
    Size:   &size,
})

continents := response.JSON200.Data.Continents // Continents that match the name
meta := response.JSON200.Data.Meta             // Meta information
```

## 💼 Job titles

### Enrich a job title from its name

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Job-titles/operation/enrichJobTitles)

```go
// Enrich "chief marketing officer"
name := "chief marketing officer"
response, err := client.EnrichJobTitles(ctx, &tca.EnrichJobTitlesParams{
    Name: &name,
})

jobTitle := response.JSON200.Data // Contains department, seniority, etc.
```

## 📋 Lists

### Fetch your lists

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Lists/operation/fetchLists)

```go
// Fetch your lists
response, err := client.FetchLists(ctx, &tca.FetchListsParams{})

lists := response.JSON200.Data.Lists // Lists that match the query
meta := response.JSON200.Data.Meta   // Meta information
```

### Create a list of companies

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Lists/operation/createList)

```go
// Create a list of companies
response, err := client.CreateList(ctx, tca.CreateListJSONRequestBody{
    Name: "My SaaS List",
    Type: "companies",
})

newList := response.JSON200.Data // The new list
```

### Fetch companies in your list

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Lists/operation/fetchCompaniesInList)

```go
// Fetch companies in a list
listId := float32(1234)
response, err := client.FetchCompaniesInList(ctx, listId, &tca.FetchCompaniesInListParams{})

companies := response.JSON200.Data.Companies // Companies that match the list
meta := response.JSON200.Data.Meta           // Meta information
```

### Add or remove companies in your list

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Lists/operation/listsToggleCompanies)

```go
// Add companies to a list
listId := float32(1234)
response, err := client.ListsToggleCompanies(ctx, listId, tca.ListsToggleCompaniesJSONRequestBody{
    Companies: []string{"apple.com", "stripe.com"},
})

list := response.JSON200.Data // The updated list
```

## 👥 Teams

### Fetch your team

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Teams/operation/fetchTeam)

```go
// Fetch your team details
teamId := float32(1234)
response, err := client.FetchTeam(ctx, teamId)

team := response.JSON200.Data // Your team details
```

## 🔧 Utilities

### Fetch the health of the API

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Utilities/operation/fetchApiHealth)

```go
// Check API health status
response, err := client.FetchApiHealth(ctx)

health := response.JSON200.Data // The health of the API
```

### Fetch the OpenAPI schema

📖 [Documentation](https://www.thecompaniesapi.com/api#tag/Utilities/operation/fetchOpenApi)

```go
// Fetch OpenAPI schema
response, err := client.FetchOpenApi(ctx)

schema := response.JSON200 // The OpenAPI schema
```

## 🧪 Testing

### Unit Tests

Run the unit tests (no API key required):

```bash
make test-unit
# or
go test -v . -short
```

### Integration Tests

Integration tests verify the SDK works with the real API. They require a valid API token.

#### Setup for Integration Tests

1. **Get your API token** from [The Companies API](https://www.thecompaniesapi.com/)

2. **Create a `.env` file** in the project root:
   ```bash
   # Required for integration tests
   TCA_API_TOKEN=your_actual_api_token_here
   
   # Optional: Custom base URL
   # TCA_API_URL=https://api.thecompaniesapi.com
   
   # Optional: Visitor ID for analytics tracking
   # TCA_VISITOR_ID=your_visitor_id_here
   ```

3. **Run integration tests**:
   ```bash
   make test-integration
   # or
   go test -v . -run "TestIntegration"
   ```

#### Integration Test Coverage

The integration tests verify:
- ✅ **SearchCompanies**: Basic search and complex query conditions
- ✅ **CountCompanies**: Company counting with filters
- ✅ **FetchCompanyByEmail**: Email-based company lookup
- ✅ **Error Handling**: API error responses and validation
- ✅ **Query Serialization**: Complex nested parameters

#### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `TCA_API_TOKEN` | ✅ | Your API token from The Companies API |
| `TCA_API_URL` | ❌ | Custom API base URL (defaults to production) |
| `TCA_VISITOR_ID` | ❌ | Visitor ID for analytics tracking |

### Run All Tests

```bash
make test
# or
go test -v .
```

This runs both unit tests and integration tests (if `TCA_API_TOKEN` is available).

## 🛠️ Development

### Available Commands

```bash
make help              # Show all available commands
make generate          # Generate models from OpenAPI spec
make test              # Run all tests
make test-unit         # Run only unit tests
make test-integration  # Run only integration tests
make clean             # Clean generated files
make deps              # Install dependencies
```

### Architecture

- **`client.go`**: Base HTTP client with authentication and query serialization
- **`wrapper.go`**: High-level API methods using generated types
- **`generated.go`**: Auto-generated types and client code from OpenAPI spec
- **`*_test.go`**: Unit tests for client functionality
- **`integration_test.go`**: Integration tests with real API calls

### Helper Functions

The SDK provides several helper functions for working with pointer types:

```go
// Convert values to pointers for optional parameters
page := tca.Float32Ptr(1)
size := tca.Float32Ptr(10)
search := tca.StringPtr("technology")

response, err := client.SearchCompanies(ctx, &tca.SearchCompaniesParams{
    Page:   page,
    Size:   size,
    Search: search,
})
```

## 📄 License

This SDK is released under the MIT License. See [LICENSE](LICENSE) for details.

## 🔗 Links

- [The Companies API](https://www.thecompaniesapi.com)
- [API Documentation](https://www.thecompaniesapi.com/api)
- [TypeScript SDK](https://github.com/thecompaniesapi/sdk-typescript)
- [Support & Live Chat](https://www.thecompaniesapi.com/)
