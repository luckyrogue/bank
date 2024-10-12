package services

import "bank/models"

type AccountOperations interface {
	CreateAccount(accountType string) (*models.Account, error)
	GetAccount(id string) (*models.Account, error)
	UpdateAccount(id string, updatedData models.Account) (*models.Account, error)
	DeleteAccount(id string) error
}

type BalanceOperations interface {
	DebitAccount(id string, amount float64) error
	CreditAccount(id string, amount float64) error
}

type FreezeOperations interface {
	FreezeAccount(id string) error
}
