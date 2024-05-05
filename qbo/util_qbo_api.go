package qbo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func attributesToJson(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	companyInfo, ok := d.Value.(*CompanyInfo)
	if !ok {
		return nil, nil // or error, depending on your error handling policy
	}

	jsonData, err := json.Marshal(companyInfo.Attributes)
	if err != nil {
		return nil, err
	}

	return string(jsonData), nil
}

func getDiscoveryDocument(url string) (*DiscoveryDocument, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching discovery document: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var doc DiscoveryDocument
	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("error decoding discovery document: %v", err)
	}

	return &doc, nil
}

// TODO: figure out how to cache the token and discovery document (context?)
func qboApiCall[T any](apiResponse *T,
	urlQuery string,
	_ context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (*T, error) {

	config := GetConfig(d.Connection)

	discoveryDoc, err := getDiscoveryDocument("https://developer.api.intuit.com/.well-known/openid_sandbox_configuration")
	if err != nil {
		return nil, fmt.Errorf("error getting discover doc: %v", err)
	}

	oauth2Config := &oauth2.Config{
		ClientID:     *config.ClientId,
		ClientSecret: *config.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: discoveryDoc.TokenEndpoint, // Token endpoint for refresh
		},
		// Optionally include Scopes if required:
		// Scopes: []string{"scope1", "scope2"},
	}

	token := &oauth2.Token{
		RefreshToken: *config.RefreshToken,
		TokenType:    "Bearer",
		// Expiry is important for the client to know when to refresh the token
		Expiry: time.Now().Add(-time.Hour), // Set to past to trigger refresh immediately
	}

	// Create a token source from the token
	tokenSource := oauth2Config.TokenSource(context.Background(), token)
	_, err = tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("error getting token: %v", err)
	}

	client := oauth2.NewClient(context.Background(), tokenSource)
	request, err := http.NewRequest("GET", fmt.Sprintf(urlQuery,
		*config.BaseURL, *config.RealmId, *config.RealmId), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create an https request: %v", err)
	}

	request.Header.Set("Accept", "application/json") // Requests JSON content

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error requesting content from server: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request rejected by server: %v", response)
	}

	err = json.NewDecoder(response.Body).Decode(apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding company info response: %v", err)
	}

	return apiResponse, nil
}
