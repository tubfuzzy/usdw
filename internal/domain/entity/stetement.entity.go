package entity

type StatementRequest struct {
	Items []StatementItem `json:"items"`
}

type StatementResponse struct {
	Pagination *Pagination       `json:"pagination,omitempty"` // Added for Get Statements
	Items      []StatementResult `json:"items"`
}

type StatementResult struct {
	ID                 string       `json:"id,omitempty"`
	FeedConnectionID   string       `json:"feedConnectionId"`
	Status             string       `json:"status"`
	StartDate          string       `json:"startDate,omitempty"` // Exists in Get Statements
	EndDate            string       `json:"endDate,omitempty"`   // Exists in Get Statements
	StartBalance       *Balance     `json:"startBalance,omitempty"`
	EndBalance         *Balance     `json:"endBalance,omitempty"`
	StatementLineCount string       `json:"statementLineCount,omitempty"` // Exists in Get Statements
	Errors             *[]FeedError `json:"errors,omitempty"`             // Exists in Post Statements (if rejected)
}

type StatementItem struct {
	FeedConnectionID string          `json:"feedConnectionId"`
	StartDate        string          `json:"startDate"`
	EndDate          string          `json:"endDate"`
	StartBalance     Balance         `json:"startBalance"`
	EndBalance       Balance         `json:"endBalance"`
	StatementLines   []StatementLine `json:"statementLines"`
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
