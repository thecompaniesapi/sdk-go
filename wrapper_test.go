package thecompaniesapi_test

import (
	"context"
	"testing"
	
	"github.com/thecompaniesapi/sdk-go"
)

func TestCompaniesAPIClient(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	if client == nil {
		t.Fatal("NewClient returned nil")
	}

	// Test that the BaseURL method is still available (delegated to internal baseClient)
	if client.BaseURL() != thecompaniesapi.DefaultBaseURL {
		t.Errorf("Expected base URL %s, got %s", thecompaniesapi.DefaultBaseURL, client.BaseURL())
	}
}

func TestCompaniesAPIClientWithOptions(t *testing.T) {
	customBaseURL := "https://custom-api.example.com"
	visitorID := "test-visitor-123"
	
	client, err := thecompaniesapi.ApiClient("test-api-key",
		thecompaniesapi.WithCustomBaseURL(customBaseURL),
		thecompaniesapi.WithVisitorID(visitorID),
	)
	if err != nil {
		t.Fatalf("NewClient with options returned error: %v", err)
	}

	if client == nil {
		t.Fatal("NewClient with options returned nil")
	}

	if client.BaseURL() != customBaseURL {
		t.Errorf("Expected base URL %s, got %s", customBaseURL, client.BaseURL())
	}
}

func TestCleanAPIInterface(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	// Test that users have access to the clean API they need
	
	// ✅ Users can access BaseURL for debugging
	baseURL := client.BaseURL()
	if baseURL == "" {
		t.Error("BaseURL should be accessible")
	}

	// ✅ Users have access to all API operations
	ctx := context.Background()
	
	// Health check
	_, err = client.FetchApiHealth(ctx)
	if err != nil {
		t.Logf("FetchApiHealth failed as expected with test key: %v", err)
	}

	// Search operations
	params := &thecompaniesapi.SearchCompaniesParams{}
	_, err = client.SearchCompanies(ctx, params)
	if err != nil {
		t.Logf("SearchCompanies failed as expected with test key: %v", err)
	}

	// Note: Users can no longer access BaseClient internals like:
	// - client.HTTPClient() ❌ (properly hidden)
	// - client.MakeRequest() ❌ (properly hidden)
	// - client.BuildQueryString() ❌ (properly hidden)
	
	// This ensures a clean, focused API surface
}

func TestSearchCompaniesMethod(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	// Test with basic parameters
	size := float32(10)
	page := float32(1)
	search := "technology"
	
	params := &thecompaniesapi.SearchCompaniesParams{
		Page:   &page,
		Size:   &size,
		Search: &search,
	}

	// This will fail with the test API key, but we're testing the method signature and structure
	response, err := client.SearchCompanies(context.Background(), params)
	
	if err == nil {
		t.Log("SearchCompanies succeeded (unexpected with test API key)")
		if response != nil && response.JSON200 != nil {
			t.Logf("Response has %d companies", len(response.JSON200.Companies))
		}
	} else {
		t.Logf("SearchCompanies failed as expected: %v", err)
	}
}

func TestCountCompaniesMethod(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	params := &thecompaniesapi.CountCompaniesParams{}

	// This will fail with the test API key, but we're testing the method signature
	response, err := client.CountCompanies(context.Background(), params)
	
	if err == nil {
		t.Log("CountCompanies succeeded (unexpected with test API key)")
		if response != nil && response.JSON200 != nil {
			t.Logf("Count: %f", response.JSON200.Count)
		}
	} else {
		t.Logf("CountCompanies failed as expected: %v", err)
	}
}

func TestFetchCompanyByEmailMethod(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	params := &thecompaniesapi.FetchCompanyByEmailParams{
		Email: "test@example.com",
	}

	// This will fail with the test API key, but we're testing the method signature
	response, err := client.FetchCompanyByEmail(context.Background(), params)
	
	if err == nil {
		t.Log("FetchCompanyByEmail succeeded (unexpected with test API key)")
		if response != nil && response.JSON200 != nil {
			t.Logf("Email domain: %s", response.JSON200.Email.Domain)
		}
	} else {
		t.Logf("FetchCompanyByEmail failed as expected: %v", err)
	}
}

// Test a few more methods to demonstrate the comprehensive API coverage
func TestAdditionalMethods(t *testing.T) {
	client, err := thecompaniesapi.ApiClient("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	// Test FetchApiHealth
	_, err = client.FetchApiHealth(context.Background())
	if err != nil {
		t.Logf("FetchApiHealth failed as expected: %v", err)
	}

	// Test SearchIndustries
	params := &thecompaniesapi.SearchIndustriesParams{}
	_, err = client.SearchIndustries(context.Background(), params)
	if err != nil {
		t.Logf("SearchIndustries failed as expected: %v", err)
	}

	// Test FetchUser
	_, err = client.FetchUser(context.Background())
	if err != nil {
		t.Logf("FetchUser failed as expected: %v", err)
	}
} 
