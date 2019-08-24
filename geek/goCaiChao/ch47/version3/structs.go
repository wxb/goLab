package version3

// Request Request
type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad       []int  `json:"payload"`
}

// Response Response
type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
