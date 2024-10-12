package services

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	accountService := NewAccountService()

	account, err := accountService.CreateAccount("Current")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if account.AccountType != "Current" {
		t.Errorf("expected account type 'Current', got %s", account.AccountType)
	}

	if account.Balance != 0 {
		t.Errorf("expected initial balance to be 0, got %f", account.Balance)
	}
}

func TestCreateInvalidAccount(t *testing.T) {
	accountService := NewAccountService()

	_, err := accountService.CreateAccount("Invalid")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestGetAccount(t *testing.T) {
	accountService := NewAccountService()

	createdAccount, _ := accountService.CreateAccount("Current")
	account, err := accountService.GetAccount(createdAccount.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if account.ID != createdAccount.ID {
		t.Errorf("expected account ID %s, got %s", createdAccount.ID, account.ID)
	}
}

func TestGetNonExistingAccount(t *testing.T) {
	accountService := NewAccountService()

	_, err := accountService.GetAccount("non-existing-id")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
