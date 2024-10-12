package services

import (
	"testing"
)

func TestDebitAccount(t *testing.T) {
	accountService := NewAccountService()
	balanceService := NewBalanceService(accountService.Accounts)

	account, _ := accountService.CreateAccount("Current")
	account.Balance = 1000

	err := balanceService.DebitAccount(account.ID, 500)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if account.Balance != 500 {
		t.Errorf("expected balance to be 500, got %f", account.Balance)
	}
}

func TestDebitAccountInsufficientFunds(t *testing.T) {
	accountService := NewAccountService()
	balanceService := NewBalanceService(accountService.Accounts)

	account, _ := accountService.CreateAccount("Current")
	account.Balance = 100

	err := balanceService.DebitAccount(account.ID, 500)
	if err == nil {
		t.Fatal("expected error for insufficient funds, got none")
	}
}

func TestCreditAccount(t *testing.T) {
	accountService := NewAccountService()
	balanceService := NewBalanceService(accountService.Accounts)

	account, _ := accountService.CreateAccount("Current")
	account.Balance = 500

	err := balanceService.CreditAccount(account.ID, 500)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if account.Balance != 1000 {
		t.Errorf("expected balance to be 1000, got %f", account.Balance)
	}
}
