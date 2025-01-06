package service

import (
	"context"
	"fmt"
	"usdw/config"
	"usdw/internal/domain"
	"usdw/internal/domain/entity"
	"usdw/pkg/cache"
	"usdw/pkg/logger"
)

type bankFeedService struct {
	domain.BankFeedRepository
	*config.Configuration
}

func NewBankFeedService(bankFeedRepository domain.BankFeedRepository, config *config.Configuration, cache cache.Engine, logger logger.Logger) domain.BankFeedService {
	return &bankFeedService{
		BankFeedRepository: bankFeedRepository,
		Configuration:      config,
	}
}

func (s *bankFeedService) CreateConnections(ctx context.Context, request domain.CreateConnectionsRequest) (*domain.CreateConnectionsResponse, error) {
	// Convert domain request to entity request
	entityRequest := entity.FeedConnectionRequest{
		Items: make([]entity.FeedConnectionItem, len(request.Items)),
	}

	for i, item := range request.Items {
		entityRequest.Items[i] = entity.FeedConnectionItem(item)
	}

	// Call repository
	entityResponse, err := s.BankFeedRepository.CreateConnections(ctx, entityRequest)
	if err != nil {
		return nil, err
	}

	// Convert entity response to domain response
	domainResponse := &domain.CreateConnectionsResponse{
		Items: make([]domain.CreateConnectionResult, len(entityResponse.Items)),
	}

	for i, item := range entityResponse.Items {
		domainResponse.Items[i] = domain.CreateConnectionResult{
			ID:           item.ID,
			AccountToken: item.AccountToken,
			Status:       item.Status,
			Error:        (*domain.FeedError)(item.Error),
		}
	}

	return domainResponse, nil
}

func (s *bankFeedService) GetConnections(ctx context.Context, page, pageSize int) (*domain.ConnectionsResponse, error) {
	// Fetch from repository (returns entity struct)
	entityResponse, err := s.BankFeedRepository.FetchConnections(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	// Map entity struct to domain struct
	domainConnections := make([]domain.Connection, len(entityResponse.Items))
	for i, conn := range entityResponse.Items {
		domainConnections[i] = domain.Connection{
			ID:            conn.ID,
			AccountToken:  conn.AccountToken,
			AccountType:   conn.AccountType,
			AccountNumber: conn.AccountNumber,
			AccountName:   conn.AccountName,
			AccountID:     conn.AccountID,
			Currency:      conn.Currency,
		}
	}

	return &domain.ConnectionsResponse{
		Pagination: domain.Pagination{
			Page:      entityResponse.Pagination.Page,
			PageSize:  entityResponse.Pagination.PageSize,
			PageCount: entityResponse.Pagination.PageCount,
			ItemCount: entityResponse.Pagination.ItemCount,
		},
		Items: domainConnections,
	}, nil
}

func (s *bankFeedService) GetConnectionByID(ctx context.Context, feedConnectionID string) (*domain.Connection, error) {
	entityConnection, err := s.BankFeedRepository.FetchConnectionByID(ctx, feedConnectionID)
	if err != nil {
		return nil, err
	}

	// Map entity struct to domain struct
	domainConnection := &domain.Connection{
		ID:            entityConnection.ID,
		AccountToken:  entityConnection.AccountToken,
		AccountType:   entityConnection.AccountType,
		AccountNumber: entityConnection.AccountNumber,
		AccountName:   entityConnection.AccountName,
		AccountID:     entityConnection.AccountID,
		Currency:      entityConnection.Currency,
	}

	return domainConnection, nil
}

func (s *bankFeedService) DeleteConnection(ctx context.Context, feedConnectionID string) (*domain.DeleteResult, error) {
	// Convert request to entity format
	entityRequest := entity.DeleteRequest{
		Items: []entity.DeleteItem{
			{ID: feedConnectionID},
		},
	}

	// Call repository
	entityResponse, err := s.BankFeedRepository.DeleteConnection(ctx, entityRequest)
	if err != nil {
		return nil, err
	}

	// Ensure we received a valid response
	if len(entityResponse.Items) == 0 {
		return nil, fmt.Errorf("no response received for feed connection ID: %s", feedConnectionID)
	}

	// Extract the first (and only) item
	item := entityResponse.Items[0]

	// Convert entity response to domain response
	domainResponse := &domain.DeleteResult{
		ID:           item.ID,
		AccountToken: item.AccountToken,
		Status:       item.Status,
		Error:        (*domain.FeedError)(item.Error),
	}

	return domainResponse, nil
}

func (s *bankFeedService) PostStatements(ctx context.Context, request domain.PostStatementRequest) (*domain.PostStatementResponse, error) {
	// Convert domain request to entity request
	entityRequest := entity.StatementRequest{
		Items: make([]entity.StatementItem, len(request.Items)),
	}

	for i, item := range request.Items {
		entityRequest.Items[i] = entity.StatementItem{
			FeedConnectionID: item.FeedConnectionID,
			StartDate:        item.StartDate,
			EndDate:          item.EndDate,
			StartBalance: entity.Balance{
				Amount:               item.StartBalance.Amount,
				CreditDebitIndicator: item.StartBalance.CreditDebitIndicator,
			},
			EndBalance: entity.Balance{
				Amount:               item.EndBalance.Amount,
				CreditDebitIndicator: item.EndBalance.CreditDebitIndicator,
			},
			StatementLines: mapStatementLines(item.StatementLines),
		}
	}

	// Call repository
	entityResponse, err := s.BankFeedRepository.PostStatements(ctx, entityRequest)
	if err != nil {
		return nil, err
	}

	// Convert entity response to domain response
	domainResponse := &domain.PostStatementResponse{
		Items: make([]domain.StatementResult, len(entityResponse.Items)),
	}

	for i, item := range entityResponse.Items {
		domainResponse.Items[i] = domain.StatementResult{
			ID:               item.ID,
			FeedConnectionID: item.FeedConnectionID,
			Status:           item.Status,
			Errors:           mapErrors(item.Errors),
		}
	}

	return domainResponse, nil
}

func (s *bankFeedService) GetStatements(ctx context.Context, page, pageSize int) (*domain.GetStatementsResponse, error) {
	entityResponse, err := s.BankFeedRepository.GetStatements(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	domainResponse := &domain.GetStatementsResponse{
		Pagination: mapPagination(entityResponse.Pagination),
		Items:      mapStatementResults(entityResponse.Items),
	}

	return domainResponse, nil
}

func (s *bankFeedService) GetStatementByID(ctx context.Context, statementID string) (*domain.StatementResult, error) {
	entityResponse, err := s.BankFeedRepository.GetStatementByID(ctx, statementID)
	if err != nil {
		return nil, err
	}

	return &domain.StatementResult{
		ID:                 entityResponse.ID,
		FeedConnectionID:   entityResponse.FeedConnectionID,
		Status:             entityResponse.Status,
		StartDate:          entityResponse.StartDate,
		EndDate:            entityResponse.EndDate,
		StartBalance:       mapBalance(entityResponse.StartBalance),
		EndBalance:         mapBalance(entityResponse.EndBalance),
		StatementLineCount: entityResponse.StatementLineCount,
	}, nil
}
