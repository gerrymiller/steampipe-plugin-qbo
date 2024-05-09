package qbo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/oauth2"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Global variable to hold state
var globalState sync.Map

func getDiscoveryDocument(url string) (DiscoveryDocument, error) {
	var doc DiscoveryDocument
	resp, err := http.Get(url)
	if err != nil {
		return doc, fmt.Errorf("error fetching discovery document: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return doc, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return doc, fmt.Errorf("error decoding discovery document: %v", err)
	}

	return doc, nil
}

func getCacheValue[T any](key string) (T, bool) {
	value, exists := globalState.Load(key)
	if !exists {
		var empty T
		return empty, false
	}
	return value.(T), true
}

func setCacheValue[T any](key string, value T) {
	globalState.Store(key, value)
}

// TODO: figure out how to cache the token and discovery document (context?)
func qboApiCall[T any](apiResponse *T,
	urlQuery string,
	_ context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (*T, error) {

	config := GetConfig(d.Connection)

	var tokenSource oauth2.TokenSource
	if cached, exists := getCacheValue[oauth2.TokenSource]("tokenSource"); exists {
		tokenSource = cached
	} else {
		var oauth2Config oauth2.Config
		if cached, exists := getCacheValue[oauth2.Config]("oauth2Config"); exists {
			oauth2Config = cached
		} else {
			var discoveryDoc DiscoveryDocument
			if cached, exists := getCacheValue[DiscoveryDocument]("discoveryDoc"); exists {
				discoveryDoc = cached
			} else {
				returnedDoc, err := getDiscoveryDocument(*config.DiscoveryDocumentURL)
				if err != nil {
					return nil, fmt.Errorf("error getting discover doc: %v", err)
				}
				discoveryDoc = returnedDoc
				setCacheValue("discoveryDoc", discoveryDoc)
			}
			oauth2Config = oauth2.Config{
				ClientID:     *config.ClientId,
				ClientSecret: *config.ClientSecret,
				Endpoint: oauth2.Endpoint{
					TokenURL: discoveryDoc.TokenEndpoint, // Token endpoint for refresh
				},
				// Optionally include Scopes if required:
				// Scopes: []string{"scope1", "scope2"},
			}
			setCacheValue("oauth2Config", oauth2Config)
		}

		token := &oauth2.Token{
			AccessToken:  *config.AccessToken,
			RefreshToken: *config.RefreshToken,
			TokenType:    "Bearer",
			// Expiry is important for the client to know when to refresh the token
			Expiry: time.Now().Add(-time.Hour), // Set to past to trigger refresh immediately
		}

		// Create a token source from the token
		tokenSource = oauth2Config.TokenSource(context.Background(), token)

		setCacheValue("tokenSource", tokenSource)
	}

	// Force refresh if neccessary
	_, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("error getting token: %v", err)
	}
	// Update cache if tokenSource has been updated
	setCacheValue("tokenSource", tokenSource)

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
