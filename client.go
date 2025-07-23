package thecompaniesapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for The Companies API
	DefaultBaseURL = "https://api.thecompaniesapi.com"
	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 300 * time.Second
)

// BaseClient represents The Companies API client foundation
type BaseClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	visitorID  string // Added for visitor ID support
}

// BaseClientOption is a function type for configuring the client
type BaseClientOption func(*BaseClient)

// WithCustomBaseURL sets a custom base URL for the client
func WithCustomBaseURL(baseURL string) BaseClientOption {
	return func(c *BaseClient) {
		c.baseURL = baseURL
	}
}

// WithCustomHTTPClient sets a custom HTTP client
func WithCustomHTTPClient(httpClient *http.Client) BaseClientOption {
	return func(c *BaseClient) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets a custom timeout for HTTP requests
func WithTimeout(timeout time.Duration) BaseClientOption {
	return func(c *BaseClient) {
		c.httpClient.Timeout = timeout
	}
}

// WithVisitorID sets a custom visitor ID for the client
func WithVisitorID(visitorID string) BaseClientOption {
	return func(c *BaseClient) {
		c.visitorID = visitorID
	}
}

// NewBaseClient creates a new Companies API client
func NewBaseClient(apiKey string, options ...BaseClientOption) *BaseClient {
	client := &BaseClient{
		baseURL: DefaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// Error represents an API error response
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *Error) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// BuildQueryString serializes query parameters
// - Objects and arrays are JSON stringified then URL encoded
// - Primitives are converted to strings
func (c *BaseClient) BuildQueryString(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	var parts []string

	for key, value := range params {
		if value == nil {
			continue
		}

		encodedKey := url.QueryEscape(key)
		var encodedValue string

		// Use reflection to determine the type
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			// Objects and arrays: JSON stringify then URL encode
			jsonBytes, err := json.Marshal(value)
			if err != nil {
				// Fallback to string representation
				encodedValue = url.QueryEscape(fmt.Sprintf("%v", value))
			} else {
				encodedValue = url.QueryEscape(string(jsonBytes))
			}

		case reflect.Ptr:
			// Handle pointers by dereferencing
			if v.IsNil() {
				continue
			}
			elem := v.Elem()
			switch elem.Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
				jsonBytes, err := json.Marshal(elem.Interface())
				if err != nil {
					encodedValue = url.QueryEscape(fmt.Sprintf("%v", elem.Interface()))
				} else {
					encodedValue = url.QueryEscape(string(jsonBytes))
				}
			default:
				// Primitive pointer: convert to string (no additional encoding needed for primitives)
				encodedValue = url.QueryEscape(fmt.Sprintf("%v", elem.Interface()))
			}

		default:
			// Primitives: convert to string (no additional encoding needed)
			switch val := value.(type) {
			case string:
				encodedValue = url.QueryEscape(val)
			case int, int8, int16, int32, int64:
				encodedValue = url.QueryEscape(fmt.Sprintf("%d", val))
			case uint, uint8, uint16, uint32, uint64:
				encodedValue = url.QueryEscape(fmt.Sprintf("%d", val))
			case float32, float64:
				encodedValue = url.QueryEscape(fmt.Sprintf("%g", val))
			case bool:
				encodedValue = url.QueryEscape(strconv.FormatBool(val))
			default:
				encodedValue = url.QueryEscape(fmt.Sprintf("%v", val))
			}
		}

		parts = append(parts, encodedKey+"="+encodedValue)
	}

	// Sort to ensure consistent output (matches Go's url.Values behavior)
	// This helps with testing and debugging
	if len(parts) > 1 {
		// Simple sort by key name (extract key from "key=value")
		for i := 0; i < len(parts)-1; i++ {
			for j := i + 1; j < len(parts); j++ {
				keyI := strings.Split(parts[i], "=")[0]
				keyJ := strings.Split(parts[j], "=")[0]
				if keyI > keyJ {
					parts[i], parts[j] = parts[j], parts[i]
				}
			}
		}
	}

	return strings.Join(parts, "&")
}

// MakeRequestWithQuery performs an HTTP request with query parameters serialized
func (c *BaseClient) MakeRequestWithQuery(ctx context.Context, method, path string, queryParams map[string]interface{}, body any) ([]byte, error) {
	fullPath := path
	if len(queryParams) > 0 {
		queryString := c.BuildQueryString(queryParams)
		if queryString != "" {
			separator := "?"
			if strings.Contains(path, "?") {
				separator = "&"
			}
			fullPath = path + separator + queryString
		}
	}

	return c.MakeRequest(ctx, method, fullPath, body)
}

// MakeRequest performs an HTTP request with authentication and returns the response body
func (c *BaseClient) MakeRequest(ctx context.Context, method, path string, body any) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+c.apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.visitorID != "" {
		req.Header.Set("Tca-Visitor-Id", c.visitorID)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr Error
		if err := json.Unmarshal(responseBody, &apiErr); err != nil {
			return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(responseBody))
		}
		return nil, &apiErr
	}

	return responseBody, nil
}

// BaseURL returns the configured base URL
func (c *BaseClient) BaseURL() string {
	return c.baseURL
}

// HTTPClient returns the underlying HTTP client
func (c *BaseClient) HTTPClient() *http.Client {
	return c.httpClient
}
