package service

import (
	"usdw/internal/domain"
	"usdw/internal/domain/entity"
)

func mapStatementLines(lines []domain.StatementLine) []entity.StatementLine {
	entityLines := make([]entity.StatementLine, len(lines))
	for i, line := range lines {
		entityLines[i] = entity.StatementLine{
			PostedDate:           line.PostedDate,
			Description:          line.Description,
			Amount:               line.Amount,
			CreditDebitIndicator: line.CreditDebitIndicator,
			TransactionID:        line.TransactionID,
			PayeeName:            line.PayeeName,
			Reference:            line.Reference,
			ChequeNumber:         line.ChequeNumber,
			TransactionType:      line.TransactionType,
		}
	}
	return entityLines
}

func mapErrors(entityErrors *[]entity.FeedError) *[]domain.FeedError {
	if entityErrors == nil {
		return nil
	}

	domainErrors := make([]domain.FeedError, len(*entityErrors))
	for i, err := range *entityErrors {
		domainErrors[i] = domain.FeedError{
			Type:   err.Type,
			Title:  err.Title,
			Status: err.Status,
			Detail: err.Detail,
		}
	}

	return &domainErrors
}

func mapPagination(entityPagination *entity.Pagination) *domain.Pagination {
	if entityPagination == nil {
		return nil
	}
	return &domain.Pagination{
		Page:      entityPagination.Page,
		PageCount: entityPagination.PageCount,
		PageSize:  entityPagination.PageSize,
		ItemCount: entityPagination.ItemCount,
	}
}

func mapStatementResults(entityItems []entity.StatementResult) []domain.StatementResult {
	domainItems := make([]domain.StatementResult, len(entityItems))
	for i, item := range entityItems {
		domainItems[i] = domain.StatementResult{
			ID:                 item.ID,
			FeedConnectionID:   item.FeedConnectionID,
			Status:             item.Status,
			StartDate:          item.StartDate,
			EndDate:            item.EndDate,
			StartBalance:       mapBalance(item.StartBalance),
			EndBalance:         mapBalance(item.EndBalance),
			StatementLineCount: item.StatementLineCount,
			Errors:             mapErrors(item.Errors),
		}
	}
	return domainItems
}

func mapBalance(entityBalance *entity.Balance) *domain.Balance {
	if entityBalance == nil {
		return nil
	}
	return &domain.Balance{
		Amount:               entityBalance.Amount,
		CreditDebitIndicator: entityBalance.CreditDebitIndicator,
	}
}
