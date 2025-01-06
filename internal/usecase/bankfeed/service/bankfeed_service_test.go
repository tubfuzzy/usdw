package service

import (
	"testing"
	"usdw/internal/domain"
	"usdw/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestMapStatementLines(t *testing.T) {
	input := []domain.StatementLine{
		{
			PostedDate:           "2023-01-10",
			Description:          "Test transaction",
			Amount:               100.50,
			CreditDebitIndicator: "CREDIT",
			TransactionID:        "txn-001",
			PayeeName:            "John Doe",
			Reference:            "Ref123",
			ChequeNumber:         "0001",
			TransactionType:      "Transfer",
		},
	}

	expected := []entity.StatementLine{
		{
			PostedDate:           "2023-01-10",
			Description:          "Test transaction",
			Amount:               100.50,
			CreditDebitIndicator: "CREDIT",
			TransactionID:        "txn-001",
			PayeeName:            "John Doe",
			Reference:            "Ref123",
			ChequeNumber:         "0001",
			TransactionType:      "Transfer",
		},
	}

	result := mapStatementLines(input)
	assert.Equal(t, expected, result)
}

func TestMapErrors(t *testing.T) {
	input := &[]entity.FeedError{
		{
			Type:   "duplicate-statement",
			Title:  "Duplicate Statement",
			Status: 409,
			Detail: "The received statement was marked as a duplicate.",
		},
	}

	expected := &[]domain.FeedError{
		{
			Type:   "duplicate-statement",
			Title:  "Duplicate Statement",
			Status: 409,
			Detail: "The received statement was marked as a duplicate.",
		},
	}

	result := mapErrors(input)
	assert.Equal(t, expected, result)

	// Test nil input
	assert.Nil(t, mapErrors(nil))
}

func TestMapPagination(t *testing.T) {
	input := &entity.Pagination{
		Page:      1,
		PageCount: 5,
		PageSize:  50,
		ItemCount: 250,
	}

	expected := &domain.Pagination{
		Page:      1,
		PageCount: 5,
		PageSize:  50,
		ItemCount: 250,
	}

	result := mapPagination(input)
	assert.Equal(t, expected, result)

	// Test nil input
	assert.Nil(t, mapPagination(nil))
}

func TestMapStatementResults(t *testing.T) {
	input := []entity.StatementResult{
		{
			ID:               "stmt-001",
			FeedConnectionID: "conn-123",
			Status:           "PENDING",
			StartDate:        "2023-01-01",
			EndDate:          "2023-01-31",
			StartBalance: &entity.Balance{
				Amount:               500.00,
				CreditDebitIndicator: "DEBIT",
			},
			EndBalance: &entity.Balance{
				Amount:               600.00,
				CreditDebitIndicator: "CREDIT",
			},
			StatementLineCount: "10",
		},
	}

	expected := []domain.StatementResult{
		{
			ID:               "stmt-001",
			FeedConnectionID: "conn-123",
			Status:           "PENDING",
			StartDate:        "2023-01-01",
			EndDate:          "2023-01-31",
			StartBalance: &domain.Balance{
				Amount:               500.00,
				CreditDebitIndicator: "DEBIT",
			},
			EndBalance: &domain.Balance{
				Amount:               600.00,
				CreditDebitIndicator: "CREDIT",
			},
			StatementLineCount: "10",
		},
	}

	result := mapStatementResults(input)
	assert.Equal(t, expected, result)
}

func TestMapBalance(t *testing.T) {
	input := &entity.Balance{
		Amount:               100.00,
		CreditDebitIndicator: "DEBIT",
	}

	expected := &domain.Balance{
		Amount:               100.00,
		CreditDebitIndicator: "DEBIT",
	}

	result := mapBalance(input)
	assert.Equal(t, expected, result)

	// Test nil input
	assert.Nil(t, mapBalance(nil))
}
