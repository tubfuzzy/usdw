package entity

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	ItemCount int `json:"itemCount"`
}

type FeedConnection struct {
	ID            string `json:"id"`
	AccountToken  string `json:"accountToken"`
	AccountType   string `json:"accountType"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	AccountID     string `json:"accountId"`
	Currency      string `json:"currency"`
}

type FetchConnectionsResponse struct {
	Pagination Pagination       `json:"pagination"`
	Items      []FeedConnection `json:"items"`
}

type FeedConnectionRequest struct {
	Items []FeedConnectionItem `json:"items"`
}

type FeedConnectionItem struct {
	AccountToken  string `json:"accountToken"`
	AccountNumber string `json:"accountNumber,omitempty"`
	AccountName   string `json:"accountName,omitempty"`
	AccountType   string `json:"accountType"`
	Currency      string `json:"currency"`
	AccountID     string `json:"accountId,omitempty"`
	Country       string `json:"country,omitempty"`
}

type FeedConnectionResponse struct {
	Items []FeedConnectionResult `json:"items"`
}

type FeedConnectionResult struct {
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

type DeleteRequest struct {
	Items []DeleteItem `json:"items"`
}

type DeleteItem struct {
	ID           string `json:"id,omitempty"`
	AccountToken string `json:"accountToken,omitempty"`
}

type DeleteResponse struct {
	Items []DeleteResult `json:"items"`
}

type DeleteResult struct {
	ID           string     `json:"id,omitempty"`
	AccountToken string     `json:"accountToken,omitempty"`
	Status       string     `json:"status"`
	Error        *FeedError `json:"error,omitempty"`
}
