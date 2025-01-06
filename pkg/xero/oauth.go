package xero

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"usdw/config"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type TokenManager struct {
	Token       *oauth2.Token
	LastRefresh time.Time
	Config      *oauth2.Config
	Client      *resty.Client
}

var tokenCache *TokenManager

// GetAccessToken retrieves a valid access token, refreshing if needed
func GetAccessToken(ctx context.Context) (string, error) {
	if tokenCache == nil || tokenCache.Token.Expiry.Before(time.Now()) {
		newToken, err := refreshAccessToken(ctx)
		if err != nil {
			return "", err
		}
		tokenCache = &TokenManager{
			Token:       newToken,
			LastRefresh: time.Now(),
			Config:      config.XeroOAuthConfig,
			Client:      resty.New(),
		}
	}

	return tokenCache.Token.AccessToken, nil
}

// refreshAccessToken refreshes the OAuth access token
func refreshAccessToken(ctx context.Context) (*oauth2.Token, error) {
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     config.XeroOAuthConfig.ClientID,
			"client_secret": config.XeroOAuthConfig.ClientSecret,
		}).
		Post(config.XeroOAuthConfig.Endpoint.TokenURL)

	if err != nil {
		return nil, fmt.Errorf("failed to refresh access token: %w", err)
	}

	var token oauth2.Token
	err = json.Unmarshal(resp.Body(), &token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	return &token, nil
}
