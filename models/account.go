package models

import "time"

type Account struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Currency  string    `json:"currency"`
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	Id        int       `json:"id" gorm:"primary_key"`
	AccountID int       `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	Id            int       `json:"id" gorm:"primary_key"`
	FromAccountID int       `json:"from_account_id" mapper:"fromAccountId"`
	ToAccountID   int       `json:"to_account_id" mapper:"toAccountId"`
	Amount        float64   `json:"amount"  mapper:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
