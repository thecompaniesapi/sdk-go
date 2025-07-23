package thecompaniesapi_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/thecompaniesapi/sdk-go"
)

// loadEnvForTesting loads .env file if it exists (for local testing)
func loadEnvForTesting() {
	// Try to load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()
}

// getAPIToken gets the API token from environment variables
func getAPIToken(t *testing.T) string {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	
	loadEnvForTesting()
	
	token := os.Getenv("TCA_API_TOKEN")
	if token == "" {
		t.Skip("TCA_API_TOKEN not set, skipping integration tests. Set TCA_API_TOKEN in .env file or environment.")
	}
	return token
}

// setupIntegrationClient creates a client configured for integration testing
func setupIntegrationClient(t *testing.T) *thecompaniesapi.CompaniesAPIClient {
	token := getAPIToken(t)
	
	options := []thecompaniesapi.BaseClientOption{
		thecompaniesapi.WithTimeout(30 * time.Second), // Reasonable timeout for tests
	}
	
	// Optional: Custom base URL from environment
	if baseURL := os.Getenv("TCA_API_URL"); baseURL != "" {
		options = append(options, thecompaniesapi.WithCustomBaseURL(baseURL))
	}
	
	// Optional: Visitor ID from environment  
	if visitorID := os.Getenv("TCA_VISITOR_ID"); visitorID != "" {
		options = append(options, thecompaniesapi.WithVisitorID(visitorID))
	}
	
	client, err := thecompaniesapi.ApiClient(token, options...)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func TestIntegration_SearchCompanies(t *testing.T) {
	client := setupIntegrationClient(t)
	ctx := context.Background()

	// Test basic search
	t.Run("basic_search", func(t *testing.T) {
		page := float32(1)
		size := float32(5) // Small size for faster tests
		search := "technology"
		
		params := &thecompaniesapi.SearchCompaniesParams{
			Page:   &page,
			Size:   &size,
			Search: &search,
		}

		response, err := client.SearchCompanies(ctx, params)
		if err != nil {
			t.Fatalf("SearchCompanies failed: %v", err)
		}

		if response.JSON200 == nil {
			t.Fatal("Expected JSON200 response, got nil")
		}

		t.Logf("Search returned %d companies out of %f total", 
			len(response.JSON200.Companies), 
			response.JSON200.Meta.Total)

		if len(response.JSON200.Companies) == 0 {
			t.Log("No companies found, but request succeeded")
		} else {
			// Verify first company has expected data
			company := response.JSON200.Companies[0]
			if company.About != nil && company.About.Name != nil {
				t.Logf("First company: %s", *company.About.Name)
			}
		}
	})

	// Test search with query conditions
	t.Run("search_with_query", func(t *testing.T) {
		page := float32(1)
		size := float32(3)
		
		// Create a segmentation condition for companies in technology industry
		var technologyValueItem thecompaniesapi.SegmentationCondition_Values_Item
		err := technologyValueItem.FromSegmentationConditionValues0("technology")
		if err != nil {
			t.Fatalf("Failed to create union value: %v", err)
		}
		
		query := []thecompaniesapi.SegmentationCondition{
			{
				Attribute: thecompaniesapi.SegmentationConditionAttributeAboutIndustries,
				Operator:  thecompaniesapi.Or, // Add required operator field
				Sign:      thecompaniesapi.Equals,
				Values: []thecompaniesapi.SegmentationCondition_Values_Item{
					technologyValueItem,
				},
			},
		}
		
		// Use POST version for complex queries
		body := thecompaniesapi.SearchCompaniesPostJSONRequestBody{
			Page:  &page,
			Size:  &size,
			Query: &query,
		}

		response, err := client.SearchCompaniesPost(ctx, body)
		if err != nil {
			t.Fatalf("SearchCompanies with query failed: %v", err)
		}

		if response.JSON200 == nil {
			t.Fatal("Expected JSON200 response, got nil")
		}

		t.Logf("Query search returned %d companies", len(response.JSON200.Companies))
	})
}

func TestIntegration_CountCompanies(t *testing.T) {
	client := setupIntegrationClient(t)
	ctx := context.Background()

	t.Run("basic_count", func(t *testing.T) {
		search := "software"
		params := &thecompaniesapi.CountCompaniesParams{
			Search: &search,
		}

		response, err := client.CountCompanies(ctx, params)
		if err != nil {
			t.Fatalf("CountCompanies failed: %v", err)
		}

		if response.JSON200 == nil {
			t.Fatal("Expected JSON200 response, got nil")
		}

		t.Logf("Total companies matching 'software': %f", response.JSON200.Count)

		if response.JSON200.Count < 0 {
			t.Error("Expected non-negative count")
		}
	})

	t.Run("count_with_query", func(t *testing.T) {
		var saasValueItem thecompaniesapi.SegmentationCondition_Values_Item
		err := saasValueItem.FromSegmentationConditionValues0("saas")
		if err != nil {
			t.Fatalf("Failed to create union value: %v", err)
		}
		
		query := []thecompaniesapi.SegmentationCondition{
			{
				Attribute: thecompaniesapi.SegmentationConditionAttributeAboutIndustries,
				Operator:  thecompaniesapi.Or, // Add required operator field
				Sign:      thecompaniesapi.Equals,
				Values: []thecompaniesapi.SegmentationCondition_Values_Item{
					saasValueItem,
				},
			},
		}
		
		// Use POST version for complex queries
		body := thecompaniesapi.CountCompaniesPostJSONRequestBody{
			Query: &query,
		}

		response, err := client.CountCompaniesPost(ctx, body)
		if err != nil {
			t.Fatalf("CountCompanies with query failed: %v", err)
		}

		if response.JSON200 == nil {
			t.Fatal("Expected JSON200 response, got nil")
		}

		t.Logf("Total SaaS companies: %f", response.JSON200.Count)
	})
}

func TestIntegration_FetchCompanyByEmail(t *testing.T) {
	client := setupIntegrationClient(t)
	ctx := context.Background()

	// Test with a well-known company email
	testCases := []struct {
		name  string
		email string
	}{
		{"openai_email", "contact@openai.com"},
		{"microsoft_email", "info@microsoft.com"},
		{"google_email", "press@google.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := &thecompaniesapi.FetchCompanyByEmailParams{
				Email: tc.email,
			}

			response, err := client.FetchCompanyByEmail(ctx, params)
			if err != nil {
				t.Logf("FetchCompanyByEmail failed for %s: %v", tc.email, err)
				// Don't fail the test - the email might not be in the database
				return
			}

			if response.JSON200 == nil {
				t.Logf("No company found for email %s", tc.email)
				return
			}

			company := response.JSON200.Company
			if company.About != nil && company.About.Name != nil {
				t.Logf("Found company for %s: %s", tc.email, *company.About.Name)
			}

			if company.Domain != nil {
				t.Logf("Company domain: %s", company.Domain.Domain)
			}
		})
	}
}

func TestIntegration_ErrorHandling(t *testing.T) {
	client := setupIntegrationClient(t)
	ctx := context.Background()

	t.Run("invalid_email_format", func(t *testing.T) {
		params := &thecompaniesapi.FetchCompanyByEmailParams{
			Email: "invalid-email-format",
		}

		response, err := client.FetchCompanyByEmail(ctx, params)
		if err != nil {
			t.Logf("Expected error for invalid email format: %v", err)
			// Check if it's our custom Error type
			if apiErr, ok := err.(*thecompaniesapi.Error); ok {
				t.Logf("API Error Code: %s, Message: %s", apiErr.Code, apiErr.Message)
			}
		} else if response.JSON401 != nil {
			t.Logf("Got 401 response for invalid email: %+v", response.JSON401)
		} else if response.JSON403 != nil {
			t.Logf("Got 403 response for invalid email: %+v", response.JSON403)
		} else {
			t.Log("Unexpectedly got successful response for invalid email")
		}
	})
}

func TestIntegration_QuerySerialization(t *testing.T) {
	client := setupIntegrationClient(t)
	ctx := context.Background()

	t.Run("complex_query_serialization", func(t *testing.T) {
		page := float32(1)
		size := float32(2)
		
		// Test single condition with multiple values (API doesn't support multiple conditions)
		var techValueItem thecompaniesapi.SegmentationCondition_Values_Item
		err := techValueItem.FromSegmentationConditionValues0("technology")
		if err != nil {
			t.Fatalf("Failed to create tech union value: %v", err)
		}
		
		var saasValueItem thecompaniesapi.SegmentationCondition_Values_Item
		err = saasValueItem.FromSegmentationConditionValues0("saas")
		if err != nil {
			t.Fatalf("Failed to create saas union value: %v", err)
		}
		
		// Test with one condition having multiple values
		query := []thecompaniesapi.SegmentationCondition{
			{
				Attribute: thecompaniesapi.SegmentationConditionAttributeAboutIndustries,
				Operator:  thecompaniesapi.Or,
				Sign:      thecompaniesapi.Equals,
				Values: []thecompaniesapi.SegmentationCondition_Values_Item{
					techValueItem,
					saasValueItem,
				},
			},
		}
		
		searchFields := []thecompaniesapi.SearchCompaniesPostJSONBodySearchFields{
			"about.name",
			"domain.domain",
		}
		
		// Use POST version for complex queries
		body := thecompaniesapi.SearchCompaniesPostJSONRequestBody{
			Page:         &page,
			Size:         &size,
			Query:        &query,
			SearchFields: &searchFields,
		}

		response, err := client.SearchCompaniesPost(ctx, body)
		if err != nil {
			t.Fatalf("Complex query failed: %v", err)
		}

		if response.JSON200 == nil {
			t.Fatal("Expected JSON200 response, got nil")
		}

		t.Logf("Complex query returned %d companies", len(response.JSON200.Companies))
		
		// This test verifies that our query serialization works correctly
		// with complex nested parameters
	})
}

// Helper functions for creating pointers to basic types
func stringPtr(s string) *string {
	return &s
}

func float32Ptr(f float32) *float32 {
	return &f
} 
