package domain

import (
	"context"
	"usdw/internal/domain/entity"
)

type BankFeedService interface {
	CreateConnections(ctx context.Context, request CreateConnectionsRequest) (*CreateConnectionsResponse, error)
	GetConnections(ctx context.Context, page, pageSize int) (*ConnectionsResponse, error)
	GetConnectionByID(ctx context.Context, feedConnectionID string) (*Connection, error)
	DeleteConnection(ctx context.Context, feedConnectionID string) (*DeleteResult, error)
	PostStatements(ctx context.Context, request PostStatementRequest) (*PostStatementResponse, error)
	GetStatements(ctx context.Context, page, pageSize int) (*GetStatementsResponse, error)
	GetStatementByID(ctx context.Context, statementID string) (*StatementResult, error)
}

type BankFeedRepository interface {
	CreateConnections(ctx context.Context, request entity.FeedConnectionRequest) (*entity.FeedConnectionResponse, error)
	FetchConnections(ctx context.Context, page, pageSize int) (*entity.FetchConnectionsResponse, error)
	FetchConnectionByID(ctx context.Context, feedConnectionID string) (*entity.FeedConnection, error)
	DeleteConnection(ctx context.Context, request entity.DeleteRequest) (*entity.DeleteResponse, error)
	PostStatements(ctx context.Context, request entity.StatementRequest) (*entity.StatementResponse, error)
	GetStatements(ctx context.Context, page, pageSize int) (*entity.StatementResponse, error)
	GetStatementByID(ctx context.Context, statementID string) (*entity.StatementResult, error)
}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	ItemCount int `json:"itemCount"`
}

type Connection struct {
	ID            string `json:"id"`
	AccountToken  string `json:"accountToken"`
	AccountType   string `json:"accountType"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	AccountID     string `json:"accountId"`
	Currency      string `json:"currency"`
}

type ConnectionsResponse struct {
	Pagination Pagination   `json:"pagination"`
	Items      []Connection `json:"items"`
}

type CreateConnectionsRequest struct {
	Items []CreateConnectionItem `json:"items"`
}

type CreateConnectionItem struct {
	AccountToken  string `json:"accountToken"`
	AccountNumber string `json:"accountNumber,omitempty"`
	AccountName   string `json:"accountName,omitempty"`
	AccountType   string `json:"accountType"`
	Currency      string `json:"currency"`
	AccountID     string `json:"accountId,omitempty"`
	Country       string `json:"country,omitempty"`
}

type CreateConnectionsResponse struct {
	Items []CreateConnectionResult `json:"items"`
}

type CreateConnectionResult struct {
	ID           string     `json:"id,omitempty"`
	AccountToken string     `json:"accountToken"`
	Status       string     `json:"status"`
	Error        *FeedError `json:"error,omitempty"`
}

type FeedError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail"`
}

type DeleteResult struct {
	ID           string     `json:"id,omitempty"`
	AccountToken string     `json:"accountToken,omitempty"`
	Status       string     `json:"status"`
	Error        *FeedError `json:"error,omitempty"`
}

type PostStatementRequest struct {
	Items []PostStatementItem `json:"items"`
}

type PostStatementItem struct {
	FeedConnectionID string          `json:"feedConnectionId"`
	StartDate        string          `json:"startDate"`
	EndDate          string          `json:"endDate"`
	StartBalance     Balance         `json:"startBalance"`
	EndBalance       Balance         `json:"endBalance"`
	StatementLines   []StatementLine `json:"statementLines"`
}

type PostStatementResponse struct {
	Items []StatementResult `json:"items"`
}

type Balance struct {
	Amount               float64 `json:"amount"`
	CreditDebitIndicator string  `json:"creditDebitIndicator"`
}

type StatementLine struct {
	PostedDate           string  `json:"postedDate"`
	Description          string  `json:"description"`
	Amount               float64 `json:"amount"`
	CreditDebitIndicator string  `json:"creditDebitIndicator"`
	TransactionID        string  `json:"transactionId"`
	PayeeName            string  `json:"payeeName,omitempty"`
	Reference            string  `json:"reference,omitempty"`
	ChequeNumber         string  `json:"chequeNumber,omitempty"`
	TransactionType      string  `json:"transactionType,omitempty"`
}

type StatementResult struct {
	ID                 string       `json:"id,omitempty"`
	FeedConnectionID   string       `json:"feedConnectionId"`
	Status             string       `json:"status"`
	StartDate          string       `json:"startDate,omitempty"`
	EndDate            string       `json:"endDate,omitempty"`
	StartBalance       *Balance     `json:"startBalance,omitempty"`
	EndBalance         *Balance     `json:"endBalance,omitempty"`
	StatementLineCount string       `json:"statementLineCount,omitempty"`
	Errors             *[]FeedError `json:"errors,omitempty"`
}

type GetStatementsResponse struct {
	Pagination *Pagination       `json:"pagination,omitempty"`
	Items      []StatementResult `json:"items"`
}
