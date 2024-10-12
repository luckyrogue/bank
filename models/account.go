package models

type Account struct {
	ID          string  `json:"id"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
	IsFrozen    bool    `json:"is_frozen"`
}
