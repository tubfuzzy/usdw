package xero

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"usdw/config"
)

// TokenManager struct for caching token
type TokenManager struct {
	Token       *oauth2.Token
	LastRefresh time.Time
	Mutex       sync.Mutex
	Client      *resty.Client
}

// Global TokenManager instance
var tokenManager *TokenManager

// GetAccessToken retrieves a valid token, refreshing if needed
func GetAccessToken(ctx context.Context) (string, error) {
	// Initialize token manager if needed
	if tokenManager == nil {
		tokenManager = &TokenManager{
			Client: resty.New(),
		}
	}

	tokenManager.Mutex.Lock()
	defer tokenManager.Mutex.Unlock()

	// Refresh if token is missing or expired
	if tokenManager.Token == nil || tokenManager.Token.Expiry.Before(time.Now().Add(-5*time.Minute)) {
		//logger.Info("Fetching new Xero Access Token...")
		newToken, err := refreshAccessToken(ctx)
		if err != nil {
			//logger.Error("Failed to get Xero access token:", err)
			return "", err
		}
		tokenManager.Token = newToken
	}

	return tokenManager.Token.AccessToken, nil
}

// refreshAccessToken requests a new token from Xero
func refreshAccessToken(ctx context.Context) (*oauth2.Token, error) {
	client := resty.New()

	resp, err := client.R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     config.XeroOAuthConfig.ClientID,
			"client_secret": config.XeroOAuthConfig.ClientSecret,
			//"scope":         "bankfeeds offline_access",
		}).
		Post(config.XeroOAuthConfig.Endpoint.TokenURL)

	if err != nil {
		return nil, fmt.Errorf("failed to refresh access token: %w", err)
	}

	// Debugging: Print the raw response from Xero
	fmt.Println("Xero Token Response:", string(resp.Body()))

	// Create a struct to match the response
	var rawToken struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Parse the response correctly
	err = json.Unmarshal(resp.Body(), &rawToken)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	// ✅ Correctly set expiry time based on `expires_in`
	token := &oauth2.Token{
		AccessToken: rawToken.AccessToken,
		TokenType:   rawToken.TokenType,
		Expiry:      time.Now().Add(time.Duration(rawToken.ExpiresIn) * time.Second),
	}

	// Debugging: Print new expiry time and scope
	fmt.Println("New Access Token Expiry:", token.Expiry)
	fmt.Println("New Access Token Scope:", rawToken.Scope)

	return token, nil
}
