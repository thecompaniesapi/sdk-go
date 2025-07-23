package thecompaniesapi

import (
	"context"
	"fmt"
	"net/http"
)

// CompaniesAPIClient is the main client for interacting with The Companies API
// It provides access to all API operations with proper type safety and authentication
type CompaniesAPIClient struct {
	*ClientWithResponses // Generated operations with proper types
	baseClient *BaseClient // Internal HTTP client (not exposed to users)
}

// New creates the main client for The Companies API
// This is the primary entry point that users should use
func ApiClient(apiKey string, options ...BaseClientOption) (*CompaniesAPIClient, error) {
	baseClient := NewBaseClient(apiKey, options...)
	
	// Create the generated client using the same base URL and HTTP client with authentication
	generatedClient, err := NewClientWithResponses(
		baseClient.BaseURL(),
		WithHTTPClient(baseClient.HTTPClient()),
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Basic "+apiKey)
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create generated client: %w", err)
	}

	return &CompaniesAPIClient{
		ClientWithResponses: generatedClient,
		baseClient: baseClient,
	}, nil
}

// BaseURL returns the base URL being used by the client
// This is useful for debugging and logging purposes
func (c *CompaniesAPIClient) BaseURL() string {
	return c.baseClient.BaseURL()
}

// === API Health ===

func (c *CompaniesAPIClient) FetchApiHealth(ctx context.Context) (*FetchApiHealthResponse, error) {
	return c.ClientWithResponses.FetchApiHealthWithResponse(ctx)
}

// === Actions ===

func (c *CompaniesAPIClient) FetchActions(ctx context.Context, params *FetchActionsParams) (*FetchActionsResponse, error) {
	return c.ClientWithResponses.FetchActionsWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) RequestAction(ctx context.Context, body RequestActionJSONRequestBody) (*RequestActionResponse, error) {
	return c.ClientWithResponses.RequestActionWithResponse(ctx, body)
}

func (c *CompaniesAPIClient) RetryAction(ctx context.Context, actionId float32, body RetryActionJSONRequestBody) (*RetryActionResponse, error) {
	return c.ClientWithResponses.RetryActionWithResponse(ctx, actionId, body)
}

// === Companies Search ===

func (c *CompaniesAPIClient) SearchCompanies(ctx context.Context, params *SearchCompaniesParams) (*SearchCompaniesResponse, error) {
	return c.ClientWithResponses.SearchCompaniesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchCompaniesPost(ctx context.Context, body SearchCompaniesPostJSONRequestBody) (*SearchCompaniesPostResponse, error) {
	return c.ClientWithResponses.SearchCompaniesPostWithResponse(ctx, body)
}

func (c *CompaniesAPIClient) SearchCompaniesByName(ctx context.Context, params *SearchCompaniesByNameParams) (*SearchCompaniesByNameResponse, error) {
	return c.ClientWithResponses.SearchCompaniesByNameWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchCompaniesByPrompt(ctx context.Context, params *SearchCompaniesByPromptParams) (*SearchCompaniesByPromptResponse, error) {
	return c.ClientWithResponses.SearchCompaniesByPromptWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchSimilarCompanies(ctx context.Context, params *SearchSimilarCompaniesParams) (*SearchSimilarCompaniesResponse, error) {
	return c.ClientWithResponses.SearchSimilarCompaniesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) CountCompanies(ctx context.Context, params *CountCompaniesParams) (*CountCompaniesResponse, error) {
	return c.ClientWithResponses.CountCompaniesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) CountCompaniesPost(ctx context.Context, body CountCompaniesPostJSONRequestBody) (*CountCompaniesPostResponse, error) {
	return c.ClientWithResponses.CountCompaniesPostWithResponse(ctx, body)
}

// === Companies Analytics ===

func (c *CompaniesAPIClient) FetchCompaniesAnalytics(ctx context.Context, params *FetchCompaniesAnalyticsParams) (*FetchCompaniesAnalyticsResponse, error) {
	return c.ClientWithResponses.FetchCompaniesAnalyticsWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) ExportCompaniesAnalytics(ctx context.Context, body ExportCompaniesAnalyticsJSONRequestBody) (*ExportCompaniesAnalyticsResponse, error) {
	return c.ClientWithResponses.ExportCompaniesAnalyticsWithResponse(ctx, body)
}

// === Company Operations ===

func (c *CompaniesAPIClient) FetchCompany(ctx context.Context, domain string, params *FetchCompanyParams) (*FetchCompanyResponse, error) {
	return c.ClientWithResponses.FetchCompanyWithResponse(ctx, domain, params)
}

func (c *CompaniesAPIClient) FetchCompanyByEmail(ctx context.Context, params *FetchCompanyByEmailParams) (*FetchCompanyByEmailResponse, error) {
	return c.ClientWithResponses.FetchCompanyByEmailWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) FetchCompanyBySocial(ctx context.Context, params *FetchCompanyBySocialParams) (*FetchCompanyBySocialResponse, error) {
	return c.ClientWithResponses.FetchCompanyBySocialWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) FetchCompanyContext(ctx context.Context, domain string) (*FetchCompanyContextResponse, error) {
	return c.ClientWithResponses.FetchCompanyContextWithResponse(ctx, domain)
}

func (c *CompaniesAPIClient) FetchCompanyEmailPatterns(ctx context.Context, domain string, params *FetchCompanyEmailPatternsParams) (*FetchCompanyEmailPatternsResponse, error) {
	return c.ClientWithResponses.FetchCompanyEmailPatternsWithResponse(ctx, domain, params)
}

func (c *CompaniesAPIClient) AskCompany(ctx context.Context, domain string, body AskCompanyJSONRequestBody) (*AskCompanyResponse, error) {
	return c.ClientWithResponses.AskCompanyWithResponse(ctx, domain, body)
}

// === Industries ===

func (c *CompaniesAPIClient) SearchIndustries(ctx context.Context, params *SearchIndustriesParams) (*SearchIndustriesResponse, error) {
	return c.ClientWithResponses.SearchIndustriesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchIndustriesSimilar(ctx context.Context, params *SearchIndustriesSimilarParams) (*SearchIndustriesSimilarResponse, error) {
	return c.ClientWithResponses.SearchIndustriesSimilarWithResponse(ctx, params)
}

// === Job Titles ===

func (c *CompaniesAPIClient) EnrichJobTitles(ctx context.Context, params *EnrichJobTitlesParams) (*EnrichJobTitlesResponse, error) {
	return c.ClientWithResponses.EnrichJobTitlesWithResponse(ctx, params)
}

// === Lists ===

func (c *CompaniesAPIClient) FetchLists(ctx context.Context, params *FetchListsParams) (*FetchListsResponse, error) {
	return c.ClientWithResponses.FetchListsWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) CreateList(ctx context.Context, body CreateListJSONRequestBody) (*CreateListResponse, error) {
	return c.ClientWithResponses.CreateListWithResponse(ctx, body)
}

func (c *CompaniesAPIClient) DeleteList(ctx context.Context, listId float32) (*DeleteListResponse, error) {
	return c.ClientWithResponses.DeleteListWithResponse(ctx, listId)
}

func (c *CompaniesAPIClient) UpdateList(ctx context.Context, listId float32, body UpdateListJSONRequestBody) (*UpdateListResponse, error) {
	return c.ClientWithResponses.UpdateListWithResponse(ctx, listId, body)
}

func (c *CompaniesAPIClient) FetchCompaniesInList(ctx context.Context, listId float32, params *FetchCompaniesInListParams) (*FetchCompaniesInListResponse, error) {
	return c.ClientWithResponses.FetchCompaniesInListWithResponse(ctx, listId, params)
}

func (c *CompaniesAPIClient) FetchCompaniesInListPost(ctx context.Context, listId float32, body FetchCompaniesInListPostJSONRequestBody) (*FetchCompaniesInListPostResponse, error) {
	return c.ClientWithResponses.FetchCompaniesInListPostWithResponse(ctx, listId, body)
}

func (c *CompaniesAPIClient) ToggleCompaniesInList(ctx context.Context, listId float32, body ToggleCompaniesInListJSONRequestBody) (*ToggleCompaniesInListResponse, error) {
	return c.ClientWithResponses.ToggleCompaniesInListWithResponse(ctx, listId, body)
}

func (c *CompaniesAPIClient) FetchCompanyInList(ctx context.Context, listId float32, domain string) (*FetchCompanyInListResponse, error) {
	return c.ClientWithResponses.FetchCompanyInListWithResponse(ctx, listId, domain)
}

// === Locations ===

func (c *CompaniesAPIClient) SearchCities(ctx context.Context, params *SearchCitiesParams) (*SearchCitiesResponse, error) {
	return c.ClientWithResponses.SearchCitiesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchContinents(ctx context.Context, params *SearchContinentsParams) (*SearchContinentsResponse, error) {
	return c.ClientWithResponses.SearchContinentsWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchCounties(ctx context.Context, params *SearchCountiesParams) (*SearchCountiesResponse, error) {
	return c.ClientWithResponses.SearchCountiesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchCountries(ctx context.Context, params *SearchCountriesParams) (*SearchCountriesResponse, error) {
	return c.ClientWithResponses.SearchCountriesWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) SearchStates(ctx context.Context, params *SearchStatesParams) (*SearchStatesResponse, error) {
	return c.ClientWithResponses.SearchStatesWithResponse(ctx, params)
}

// === OpenAPI ===

func (c *CompaniesAPIClient) FetchOpenApi(ctx context.Context) (*FetchOpenApiResponse, error) {
	return c.ClientWithResponses.FetchOpenApiWithResponse(ctx)
}

// === Prompts ===

func (c *CompaniesAPIClient) FetchPrompts(ctx context.Context, params *FetchPromptsParams) (*FetchPromptsResponse, error) {
	return c.ClientWithResponses.FetchPromptsWithResponse(ctx, params)
}

func (c *CompaniesAPIClient) ProductPrompt(ctx context.Context, body ProductPromptJSONRequestBody) (*ProductPromptResponse, error) {
	return c.ClientWithResponses.ProductPromptWithResponse(ctx, body)
}

func (c *CompaniesAPIClient) PromptToSegmentation(ctx context.Context, body PromptToSegmentationJSONRequestBody) (*PromptToSegmentationResponse, error) {
	return c.ClientWithResponses.PromptToSegmentationWithResponse(ctx, body)
}

func (c *CompaniesAPIClient) DeletePrompt(ctx context.Context, promptId float32) (*DeletePromptResponse, error) {
	return c.ClientWithResponses.DeletePromptWithResponse(ctx, promptId)
}

// === Teams ===

func (c *CompaniesAPIClient) FetchTeam(ctx context.Context, teamId float32) (*FetchTeamResponse, error) {
	return c.ClientWithResponses.FetchTeamWithResponse(ctx, teamId)
}

func (c *CompaniesAPIClient) UpdateTeam(ctx context.Context, teamId float32, body UpdateTeamJSONRequestBody) (*UpdateTeamResponse, error) {
	return c.ClientWithResponses.UpdateTeamWithResponse(ctx, teamId, body)
}

// === Technologies ===

func (c *CompaniesAPIClient) SearchTechnologies(ctx context.Context, params *SearchTechnologiesParams) (*SearchTechnologiesResponse, error) {
	return c.ClientWithResponses.SearchTechnologiesWithResponse(ctx, params)
}

// === Users ===

func (c *CompaniesAPIClient) FetchUser(ctx context.Context) (*FetchUserResponse, error) {
	return c.ClientWithResponses.FetchUserWithResponse(ctx)
} 
