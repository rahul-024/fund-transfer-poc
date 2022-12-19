package models

type Account struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Currency string `json:"currency"`
	Owner    string `json:"owner"`
}
