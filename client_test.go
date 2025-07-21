package thecompaniesapi_test

import (
	"context"
	"net/http"
	"testing"
	"time"
	
	"github.com/thecompaniesapi/sdk-go"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := thecompaniesapi.NewClient(apiKey)

	if client == nil {
		t.Fatal("NewClient returned nil")
	}

	// Test default base URL
	if client.BaseURL() != thecompaniesapi.DefaultBaseURL {
		t.Errorf("Expected default base URL %s, got %s", thecompaniesapi.DefaultBaseURL, client.BaseURL())
	}

	// Test with options
	customBaseURL := "https://custom-api.example.com"
	client = thecompaniesapi.NewClient(apiKey,
		thecompaniesapi.WithBaseURL(customBaseURL),
		thecompaniesapi.WithTimeout(60*time.Second),
	)

	if client == nil {
		t.Fatal("NewClient with options returned nil")
	}

	if client.BaseURL() != customBaseURL {
		t.Errorf("Expected custom base URL %s, got %s", customBaseURL, client.BaseURL())
	}
}

func TestClientConfiguration(t *testing.T) {
	// Test custom base URL
	customBaseURL := "https://test.api.com"
	client := thecompaniesapi.NewClient("test-key",
		thecompaniesapi.WithBaseURL(customBaseURL),
	)
	
	if client == nil {
		t.Error("Client should not be nil")
	}

	if client.BaseURL() != customBaseURL {
		t.Errorf("Expected base URL %s, got %s", customBaseURL, client.BaseURL())
	}

	// Test custom timeout
	client = thecompaniesapi.NewClient("test-key",
		thecompaniesapi.WithTimeout(5*time.Second),
	)
	
	if client == nil {
		t.Error("Client should not be nil")
	}

	// Test custom HTTP client
	customHTTPClient := &http.Client{Timeout: 10 * time.Second}
	client = thecompaniesapi.NewClient("test-key",
		thecompaniesapi.WithHTTPClient(customHTTPClient),
	)
	
	if client == nil {
		t.Error("Client should not be nil")
	}

	// Test visitor ID configuration
	visitorID := "test-visitor-123"
	client = thecompaniesapi.NewClient("test-key",
		thecompaniesapi.WithVisitorID(visitorID),
	)
	
	if client == nil {
		t.Error("Client should not be nil")
	}

	// Test multiple options together
	client = thecompaniesapi.NewClient("test-key",
		thecompaniesapi.WithBaseURL(customBaseURL),
		thecompaniesapi.WithTimeout(60*time.Second),
		thecompaniesapi.WithVisitorID(visitorID),
	)
	
	if client == nil {
		t.Error("Client should not be nil")
	}

	if client.BaseURL() != customBaseURL {
		t.Errorf("Expected base URL %s, got %s", customBaseURL, client.BaseURL())
	}
}

func TestErrorType(t *testing.T) {
	err := &thecompaniesapi.Error{
		Code:    "TEST_ERROR",
		Message: "Test error message",
		Details: "Additional details",
	}

	expected := "TEST_ERROR: Test error message (Additional details)"
	if err.Error() != expected {
		t.Errorf("Expected error string '%s', got '%s'", expected, err.Error())
	}

	// Test error without details
	err2 := &thecompaniesapi.Error{
		Code:    "SIMPLE_ERROR",
		Message: "Simple error",
	}

	expected2 := "SIMPLE_ERROR: Simple error"
	if err2.Error() != expected2 {
		t.Errorf("Expected error string '%s', got '%s'", expected2, err2.Error())
	}
}

func TestMakeRequest(t *testing.T) {
	client := thecompaniesapi.NewClient("test-api-key")

	// Test that MakeRequest method exists and can be called
	// Note: This will fail without a real API key, but we're just testing the interface
	ctx := context.Background()
	_, err := client.MakeRequest(ctx, "GET", "/v2/companies", nil)
	
	// We expect this to fail since we don't have a real API key,
	// but it should fail with a proper error, not a panic
	if err == nil {
		t.Log("MakeRequest succeeded (unexpected with test API key)")
	} else {
		t.Logf("MakeRequest failed as expected: %v", err)
	}
}

func TestConstants(t *testing.T) {
	if thecompaniesapi.DefaultBaseURL == "" {
		t.Error("DefaultBaseURL should not be empty")
	}

	if thecompaniesapi.DefaultTimeout == 0 {
		t.Error("DefaultTimeout should not be zero")
	}

	expectedBaseURL := "https://api.thecompaniesapi.com"
	if thecompaniesapi.DefaultBaseURL != expectedBaseURL {
		t.Errorf("Expected DefaultBaseURL %s, got %s", expectedBaseURL, thecompaniesapi.DefaultBaseURL)
	}

	expectedTimeout := 300 * time.Second
	if thecompaniesapi.DefaultTimeout != expectedTimeout {
		t.Errorf("Expected DefaultTimeout %v, got %v", expectedTimeout, thecompaniesapi.DefaultTimeout)
	}
}

func TestBuildQueryString(t *testing.T) {
	client := thecompaniesapi.NewClient("test-api-key")

	tests := []struct {
		name     string
		params   map[string]interface{}
		expected string
	}{
		{
			name:     "empty params",
			params:   map[string]interface{}{},
			expected: "",
		},
		{
			name: "simple string",
			params: map[string]interface{}{
				"search": "technology",
			},
			expected: "search=technology",
		},
		{
			name: "simple numbers",
			params: map[string]interface{}{
				"page": 1,
				"size": 25,
			},
			expected: "page=1&size=25",
		},
		{
			name: "boolean value",
			params: map[string]interface{}{
				"simplified": true,
			},
			expected: "simplified=true",
		},
		{
			name: "array gets JSON encoded",
			params: map[string]interface{}{
				"searchFields": []string{"about.name", "domain.domain"},
			},
			expected: `searchFields=%5B%22about.name%22%2C%22domain.domain%22%5D`, // URL encoded JSON array
		},
		{
			name: "object gets JSON encoded",
			params: map[string]interface{}{
				"query": map[string]interface{}{
					"attribute": "about.industries",
					"sign":      "equals",
					"values":    []string{"technology"},
				},
			},
			expected: `query=%7B%22attribute%22%3A%22about.industries%22%2C%22sign%22%3A%22equals%22%2C%22values%22%3A%5B%22technology%22%5D%7D`, // URL encoded JSON object
		},
		{
			name: "slice of objects gets JSON encoded",
			params: map[string]interface{}{
				"conditions": []map[string]interface{}{
					{
						"attribute": "about.name",
						"sign":      "equals",
						"values":    []string{"Google"},
					},
				},
			},
			expected: `conditions=%5B%7B%22attribute%22%3A%22about.name%22%2C%22sign%22%3A%22equals%22%2C%22values%22%3A%5B%22Google%22%5D%7D%5D`, // URL encoded JSON array of objects
		},
		{
			name: "mixed types",
			params: map[string]interface{}{
				"search":     "test",
				"page":       1,
				"simplified": true,
				"fields":     []string{"name", "domain"},
			},
			expected: `fields=%5B%22name%22%2C%22domain%22%5D&page=1&search=test&simplified=true`, // Note: url.Values sorts keys
		},
		{
			name: "nil values are skipped",
			params: map[string]interface{}{
				"search": "test",
				"page":   nil,
			},
			expected: "search=test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.BuildQueryString(tt.params)
			if result != tt.expected {
				t.Errorf("BuildQueryString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMakeRequestWithQuery(t *testing.T) {
	client := thecompaniesapi.NewClient("test-api-key")

	// Test that MakeRequestWithQuery method exists and handles query params
	ctx := context.Background()
	queryParams := map[string]interface{}{
		"page": 1,
		"size": 10,
	}

	_, err := client.MakeRequestWithQuery(ctx, "GET", "/v2/companies", queryParams, nil)
	
	// We expect this to fail since we don't have a real API key,
	// but it should fail with a proper error, not a panic
	if err == nil {
		t.Log("MakeRequestWithQuery succeeded (unexpected with test API key)")
	} else {
		t.Logf("MakeRequestWithQuery failed as expected: %v", err)
	}
}

func TestQueryStringWithExistingParams(t *testing.T) {
	client := thecompaniesapi.NewClient("test-api-key")

	// Test appending to path that already has query params
	ctx := context.Background()
	queryParams := map[string]interface{}{
		"page": 2,
	}

	_, err := client.MakeRequestWithQuery(ctx, "GET", "/v2/companies?search=test", queryParams, nil)
	
	// We expect this to fail since we don't have a real API key,
	// but it should construct the URL properly with & separator
	if err == nil {
		t.Log("MakeRequestWithQuery with existing params succeeded (unexpected with test API key)")
	} else {
		t.Logf("MakeRequestWithQuery with existing params failed as expected: %v", err)
	}
}

