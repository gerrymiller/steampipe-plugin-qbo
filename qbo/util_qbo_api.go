package qbo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

/*
 * TODO: refactor all of this to different files separation of concerns, etc.
 */

// Global variable to hold state
// TODO: Refactor to structure
// TODO: Refactor to folder structure (e.g., util/cache)
var globalState sync.Map

func getCacheValue[T any](key string) (*T, bool) {
	value, exists := globalState.Load(key)
	if !exists {
		return nil, false
	}
	return value.(*T), true
}

func setCacheValue[T any](key string, value *T) {
	globalState.Store(key, value)
}

/*
 * Custom oauth2 object to pass client id and secret as basic auth
 */
type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

func refreshToken(clientID, clientSecret, refreshToken, endpoint string) (*Token, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	auth := clientID + ":" + clientSecret
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to refresh token: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	// Set the expiry time (assuming the token expires in 1 hour)
	//   Hardcoded for now, the right way to do this is to unmarshall
	//   the "expires_in" field from the response and use it to calculate
	//   the expiry time ("expires_in" is the number of seconds from now
	//   until the token expires)
	token.Expiry = time.Now().Add(1 * time.Hour)

	return &token, nil
}

type AuthenticatedClient struct {
	clientID     string
	clientSecret string
	token        *Token
	endpoint     string
}

// TODO: check return for token validation failure and trigger a refresh there
// TODO: cache the new token
func (ac *AuthenticatedClient) Do(req *http.Request) (*http.Response, error) {
	// Refresh the token if it is expired or about to expire
	if time.Now().After(ac.token.Expiry) {
		newToken, err := refreshToken(ac.clientID, ac.clientSecret, ac.token.RefreshToken, ac.endpoint)
		if err != nil {
			return nil, err
		}
		ac.token = newToken
		setCacheValue("token", ac.token)
	}

	// Set the Authorization header with the current access token
	req.Header.Set("Authorization", "Bearer "+ac.token.AccessToken)
	client := &http.Client{}
	return client.Do(req)
}

/*
 * Discovery Document
 */
func getDiscoveryDocument(url string) (*DiscoveryDocument, error) {
	var doc DiscoveryDocument
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching discovery document: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("error decoding discovery document: %v", err)
	}

	return &doc, nil
}

/*
 * QBO API Call
 */
func qboApiCall[T any](apiResponse *T,
	urlQuery string,
	_ context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (*T, error) {

	config := GetConfig(d.Connection)

	var token *Token
	var discoveryDoc *DiscoveryDocument

	if cached, exists := getCacheValue[Token]("token"); exists {
		token = cached
	} else {
		token = &Token{
			AccessToken:  *config.AccessToken,
			RefreshToken: *config.RefreshToken,
			Expiry:       time.Now().Add(-time.Hour), // Set to past to trigger refresh immediately
		}
	}

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

	authClient := &AuthenticatedClient{
		clientID:     *config.ClientId,
		clientSecret: *config.ClientSecret,
		token:        token,
		endpoint:     discoveryDoc.TokenEndpoint,
	}

	request, err := http.NewRequest("GET", fmt.Sprintf(urlQuery,
		*config.BaseURL, *config.RealmId, *config.RealmId), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create an https request: %v", err)
	}

	request.Header.Set("Accept", "application/json") // Requests JSON content

	response, err := authClient.Do(request)
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
