package services

import (
	"bank/models"
	"errors"
	"sync"
)

type FreezeService struct {
	accounts map[string]*models.Account
	mu       sync.Mutex
}

func NewFreezeService(accounts map[string]*models.Account) *FreezeService {
	return &FreezeService{
		accounts: accounts,
	}
}

func (s *FreezeService) FreezeAccount(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[id]
	if !exists {
		return errors.New("account not found")
	}
	account.IsFrozen = true
	return nil
}
