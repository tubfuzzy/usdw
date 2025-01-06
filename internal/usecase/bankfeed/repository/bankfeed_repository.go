package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"usdw/config"
	"usdw/internal/domain"
	"usdw/internal/domain/entity"
	"usdw/pkg/xero"

	"github.com/go-resty/resty/v2"
)

const bankFeedsBaseURL = "https://api.xero.com/bankfeeds.xro/1.0"

type bankFeedRepository struct {
	Config *config.Configuration
	Client *resty.Client
}

func NewBankFeedRepository(client *resty.Client, config *config.Configuration) domain.BankFeedRepository {
	return &bankFeedRepository{
		Config: config,
		Client: client,
	}
}

func (r *bankFeedRepository) authenticatedRequest(ctx context.Context) (*resty.Request, error) {
	accessToken, err := xero.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	tenantID, err := xero.GetXeroTenantID(ctx)
	if err != nil {
		return nil, err
	}

	return r.Client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("xero-tenant-id", tenantID).
		SetHeader("Content-Type", "application/json"), nil
}

func (r *bankFeedRepository) CreateConnections(ctx context.Context, request entity.FeedConnectionRequest) (*entity.FeedConnectionResponse, error) {
	url := fmt.Sprintf("%s/FeedConnections", bankFeedsBaseURL)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.SetBody(request).Post(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var response entity.FeedConnectionResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Xero response: %w", err)
	}

	return &response, nil
}

func (r *bankFeedRepository) FetchConnections(ctx context.Context, page, pageSize int) (*entity.FetchConnectionsResponse, error) {
	url := fmt.Sprintf("%s/FeedConnections?page=%d&pageSize=%d", bankFeedsBaseURL, page, pageSize)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var response entity.FetchConnectionsResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Xero response: %w", err)
	}

	return &response, nil
}

func (r *bankFeedRepository) FetchConnectionByID(ctx context.Context, feedConnectionID string) (*entity.FeedConnection, error) {
	url := fmt.Sprintf("%s/FeedConnections/%s", bankFeedsBaseURL, feedConnectionID)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("feed connection not found")
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var feedConnection entity.FeedConnection
	err = json.Unmarshal(resp.Body(), &feedConnection)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Xero response: %w", err)
	}

	return &feedConnection, nil
}

func (r *bankFeedRepository) DeleteConnection(ctx context.Context, request entity.DeleteRequest) (*entity.DeleteResponse, error) {
	url := fmt.Sprintf("%s/FeedConnections/DeleteRequests", bankFeedsBaseURL)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.SetBody(request).Post(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var response entity.DeleteResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Xero response: %w", err)
	}

	return &response, nil
}

func (r *bankFeedRepository) PostStatements(ctx context.Context, request entity.StatementRequest) (*entity.StatementResponse, error) {
	url := fmt.Sprintf("%s/Statements", bankFeedsBaseURL)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.SetBody(request).Post(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var response entity.StatementResponse
	err = json.Unmarshal(resp.Body(), &response)
	return &response, err
}

func (r *bankFeedRepository) GetStatements(ctx context.Context, page, pageSize int) (*entity.StatementResponse, error) {
	url := fmt.Sprintf("%s/Statements?page=%d&pageSize=%d", bankFeedsBaseURL, page, pageSize)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var response entity.StatementResponse
	err = json.Unmarshal(resp.Body(), &response)
	return &response, err
}

func (r *bankFeedRepository) GetStatementByID(ctx context.Context, statementID string) (*entity.StatementResult, error) {
	url := fmt.Sprintf("%s/Statements/%s", bankFeedsBaseURL, statementID)

	req, err := r.authenticatedRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Xero API: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("statement not found")
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("Xero API error [%d]: %s", resp.StatusCode(), resp.String())
	}

	var statement entity.StatementResult
	err = json.Unmarshal(resp.Body(), &statement)
	return &statement, err
}
