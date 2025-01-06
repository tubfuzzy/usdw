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

	//logger.Info("Requesting new access token from Xero...")

	resp, err := client.R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     config.XeroOAuthConfig.ClientID,
			"client_secret": config.XeroOAuthConfig.ClientSecret,
		}).
		Post(config.XeroOAuthConfig.Endpoint.TokenURL)

	if err != nil {
		return nil, fmt.Errorf("failed to refresh access token: %w", err)
	}

	// Parse the response
	var token oauth2.Token
	err = json.Unmarshal(resp.Body(), &token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	token.Expiry = time.Now().Add(time.Duration(token.Expiry.Unix()) * time.Second)
	//logger.Info("New Xero access token retrieved successfully.")

	return &token, nil
}
