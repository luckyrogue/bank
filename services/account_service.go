package services

import (
	"bank/models"
	"bank/utils"
	"errors"
	"sync"
)

type AccountService struct {
	Accounts map[string]*models.Account
	mu       sync.Mutex
}

func NewAccountService() *AccountService {
	return &AccountService{
		Accounts: make(map[string]*models.Account),
	}
}

func (s *AccountService) CreateAccount(accountType string) (*models.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if accountType != "Current" && accountType != "Savings" {
		return nil, errors.New("invalid account type")
	}

	id := utils.GenerateID()
	account := &models.Account{
		ID:          id,
		AccountType: accountType,
		Balance:     0,
		IsFrozen:    false,
	}
	s.Accounts[id] = account
	return account, nil
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.Accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func (s *AccountService) UpdateAccount(id string, updatedData models.Account) (*models.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.Accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	account.AccountType = updatedData.AccountType
	account.Balance = updatedData.Balance
	account.IsFrozen = updatedData.IsFrozen

	return account, nil
}

func (s *AccountService) DeleteAccount(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.Accounts[id]; !exists {
		return errors.New("account not found")
	}

	delete(s.Accounts, id)
	return nil
}
