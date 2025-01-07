package xero

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
)

// XeroConnection represents the tenant ID response
type XeroConnection struct {
	TenantID string `json:"tenantId"`
}

// Singleton Tenant ID cache
var tenantIDCache string
var tenantMutex sync.Mutex

// GetXeroTenantID fetches the tenant ID using the access token
func GetXeroTenantID(ctx context.Context) (string, error) {
	tenantMutex.Lock()
	defer tenantMutex.Unlock()

	// Return cached tenant ID if available
	if tenantIDCache != "" {
		return tenantIDCache, nil
	}

	// Get access token
	accessToken, err := GetAccessToken(ctx)
	if err != nil {
		return "", err
	}

	client := resty.New()
	url := "https://api.xero.com/connections"

	//logger.Info("Fetching Xero tenant ID...")
	resp, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		return "", fmt.Errorf("failed to call Xero Connections API: %w", err)
	}

	// Parse response
	var connections []XeroConnection
	err = json.Unmarshal(resp.Body(), &connections)
	if err != nil {
		return "", fmt.Errorf("failed to parse tenant ID response: %w", err)
	}

	// Ensure at least one tenant exists
	if len(connections) == 0 {
		return "", fmt.Errorf("no tenants found for this Xero account")
	}

	// Cache and return the first tenant ID
	tenantIDCache = connections[0].TenantID
	//logger.Info("Xero Tenant ID:", tenantIDCache)
	return tenantIDCache, nil
}
