package request

type TransferRequest struct {
	FromAccountID int     `json:"from_account_id" mapper:"fromAccountId"`
	ToAccountID   int     `json:"to_account_id" mapper:"toAccountId"`
	Amount        float64 `json:"amount" mapper:"amount"`
	Currency      string  `json:"currency"`
}
