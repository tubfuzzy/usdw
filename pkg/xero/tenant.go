package xero

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type XeroConnection struct {
	TenantID string `json:"tenantId"`
}

var tenantIDCache string

// GetXeroTenantID fetches the tenant ID using the access token
func GetXeroTenantID(ctx context.Context) (string, error) {
	if tenantIDCache != "" {
		return tenantIDCache, nil
	}

	accessToken, err := GetAccessToken(ctx)
	if err != nil {
		return "", err
	}

	client := resty.New()
	url := "https://api.xero.com/connections"

	resp, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		return "", fmt.Errorf("failed to call Xero Connections API: %w", err)
	}

	var connections []XeroConnection
	err = json.Unmarshal(resp.Body(), &connections)
	if err != nil || len(connections) == 0 {
		return "", fmt.Errorf("no tenants found")
	}

	tenantIDCache = connections[0].TenantID
	return tenantIDCache, nil
}
