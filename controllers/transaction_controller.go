package controllers

import (
	"bank/services"
	"bank/utils"
	"net/http"
)

type TransactionController struct {
	BalanceService services.BalanceOperations
}

func NewTransactionController(balanceService services.BalanceOperations) *TransactionController {
	return &TransactionController{
		BalanceService: balanceService,
	}
}

func (tc *TransactionController) TransferFunds(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FromAccountID string  `json:"from_account_id"`
		ToAccountID   string  `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}

	if err := utils.ParseJSONBody(w, r, &req); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Amount <= 0 {
		utils.JSONErrorResponse(w, http.StatusBadRequest, "Invalid amount")
		return
	}

	if err := tc.BalanceService.DebitAccount(req.FromAccountID, req.Amount); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := tc.BalanceService.CreditAccount(req.ToAccountID, req.Amount); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Transfer successful")
}
