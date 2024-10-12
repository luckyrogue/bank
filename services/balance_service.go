package services

import (
	"bank/models"
	"errors"
	"sync"
)

type BalanceService struct {
	accounts map[string]*models.Account
	mu       sync.Mutex
}

func NewBalanceService(accounts map[string]*models.Account) *BalanceService {
	return &BalanceService{
		accounts: accounts,
	}
}

func (s *BalanceService) DebitAccount(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[id]
	if !exists {
		return errors.New("account not found")
	}
	if account.IsFrozen {
		return errors.New("account is frozen")
	}
	if account.Balance < amount {
		return errors.New("insufficient funds")
	}
	account.Balance -= amount
	return nil
}

func (s *BalanceService) CreditAccount(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[id]
	if !exists {
		return errors.New("account not found")
	}
	if account.IsFrozen {
		return errors.New("account is frozen")
	}
	account.Balance += amount
	return nil
}
