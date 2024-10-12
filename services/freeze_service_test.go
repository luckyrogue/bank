package services

import (
	"testing"
)

func TestFreezeAccount(t *testing.T) {
	accountService := NewAccountService()
	freezeService := NewFreezeService(accountService.Accounts)

	account, _ := accountService.CreateAccount("Current")

	err := freezeService.FreezeAccount(account.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !account.IsFrozen {
		t.Errorf("expected account to be frozen, got %v", account.IsFrozen)
	}
}

func TestFreezeNonExistingAccount(t *testing.T) {
	accountService := NewAccountService()
	freezeService := NewFreezeService(accountService.Accounts)

	err := freezeService.FreezeAccount("non-existing-id")
	if err == nil {
		t.Fatal("expected error for non-existing account, got none")
	}
}
