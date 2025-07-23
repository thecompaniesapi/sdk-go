package thecompaniesapi_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thecompaniesapi/sdk-go"
)

// Example_basicUsage demonstrates basic SDK usage
func Example_basicUsage() {
	// Initialize client
	client, err := thecompaniesapi.ApiClient("your-api-key",
		thecompaniesapi.WithVisitorID("demo-visitor-123"), // Analytics tracking
		thecompaniesapi.WithTimeout(60*time.Second),       // Custom timeout
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Search companies with complex query
	searchExample(ctx, client)

	// Example 2: Count companies 
	countExample(ctx, client)

	// Example 3: Fetch company by email
	emailExample(ctx, client)
}

func searchExample(ctx context.Context, client *thecompaniesapi.CompaniesAPIClient) {
	// Complex search with generated types
	page := float32(1)
	size := float32(25)
	search := "technology"
	simplified := true

	// Generate query conditions using generated types
	var techValueItem thecompaniesapi.SegmentationCondition_Values_Item
	err := techValueItem.FromSegmentationConditionValues0("technology")
	if err != nil {
		log.Printf("Failed to create tech union value: %v", err)
		return
	}
	
	var employeeValueItem thecompaniesapi.SegmentationCondition_Values_Item
	err = employeeValueItem.FromSegmentationConditionValues1(float32(100))
	if err != nil {
		log.Printf("Failed to create employee union value: %v", err)
		return
	}

	query := []thecompaniesapi.SegmentationCondition{
		{
			Attribute: thecompaniesapi.SegmentationConditionAttributeAboutIndustries,
			Operator:  thecompaniesapi.Or, // Add required operator field
			Sign:      thecompaniesapi.Equals,
			Values: []thecompaniesapi.SegmentationCondition_Values_Item{
				techValueItem,
			},
		},
		{
			Attribute: thecompaniesapi.SegmentationConditionAttributeAboutTotalEmployees,
			Operator:  thecompaniesapi.And, // Add required operator field
			Sign:      thecompaniesapi.Greater,
			Values: []thecompaniesapi.SegmentationCondition_Values_Item{
				employeeValueItem,
			},
		},
	}

	searchFields := []thecompaniesapi.SearchCompaniesParamsSearchFields{
		"about.name",
		"domain.domain",
	}

	params := &thecompaniesapi.SearchCompaniesParams{
		Page:         &page,
		Size:         &size,
		Search:       &search,
		Simplified:   &simplified,
		Query:        &query,
		SearchFields: &searchFields,
	}

	// Call with sophisticated query serialization
	response, err := client.SearchCompanies(ctx, params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	// Use generated response types
	if response.JSON200 != nil {
		fmt.Printf("Found %f companies\n", response.JSON200.Meta.Total)
		
		for _, company := range response.JSON200.Companies {
			if company.About != nil && company.About.Name != nil {
				fmt.Printf("- %s", *company.About.Name)
				if company.Domain != nil {
					fmt.Printf(" (%s)", company.Domain.Domain)
				}
				fmt.Println()
			}
		}
	}
}

func countExample(ctx context.Context, client *thecompaniesapi.CompaniesAPIClient) {
	search := "saas"
	params := &thecompaniesapi.CountCompaniesParams{
		Search: &search,
	}

	response, err := client.CountCompanies(ctx, params)
	if err != nil {
		log.Printf("Count failed: %v", err)
		return
	}

	if response.JSON200 != nil {
		fmt.Printf("Total SaaS companies: %f\n", response.JSON200.Count)
	}
}

func emailExample(ctx context.Context, client *thecompaniesapi.CompaniesAPIClient) {
	params := &thecompaniesapi.FetchCompanyByEmailParams{
		Email: "contact@openai.com",
	}

	response, err := client.FetchCompanyByEmail(ctx, params)
	if err != nil {
		log.Printf("Email lookup failed: %v", err)
		return
	}

	if response.JSON200 != nil && response.JSON200.Company.About != nil {
		if response.JSON200.Company.About.Name != nil {
			fmt.Printf("Company from email: %s\n", *response.JSON200.Company.About.Name)
		}
	}
} 
