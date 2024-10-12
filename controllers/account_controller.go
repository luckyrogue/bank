package controllers

import (
	"bank/models"
	"bank/services"
	"bank/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountController struct {
	AccountService services.AccountOperations
	BalanceService services.BalanceOperations
	FreezeService  services.FreezeOperations
}

func NewAccountController(accountService services.AccountOperations, balanceService services.BalanceOperations, freezeService services.FreezeOperations) *AccountController {
	return &AccountController{
		AccountService: accountService,
		BalanceService: balanceService,
		FreezeService:  freezeService,
	}
}

func (ac *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountType string `json:"account_type"`
	}

	if err := utils.ParseJSONBody(w, r, &req); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.AccountType != "Current" && req.AccountType != "Savings" {
		utils.JSONErrorResponse(w, http.StatusBadRequest, "Invalid account type")
		return
	}

	account, err := ac.AccountService.CreateAccount(req.AccountType)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, account)
}

func (ac *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	account, err := ac.AccountService.GetAccount(vars["id"])
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, "Account not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, account)
}

func (ac *AccountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var req struct {
		AccountType string  `json:"account_type"`
		Balance     float64 `json:"balance"`
		IsFrozen    bool    `json:"is_frozen"`
	}

	if err := utils.ParseJSONBody(w, r, &req); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	account := &models.Account{
		AccountType: req.AccountType,
		Balance:     req.Balance,
		IsFrozen:    req.IsFrozen,
	}

	updatedAccount, err := ac.AccountService.UpdateAccount(vars["id"], *account)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, updatedAccount)
}

func (ac *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := ac.AccountService.DeleteAccount(vars["id"]); err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ac *AccountController) DebitAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := utils.ParseJSONBody(writer, request, &req); err != nil {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Amount <= 0 {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, "Invalid amount")
		return
	}

	if err := ac.BalanceService.DebitAccount(vars["id"], req.Amount); err != nil {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(writer, http.StatusOK, "Debit successful")
}

func (ac *AccountController) CreditAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := utils.ParseJSONBody(writer, request, &req); err != nil {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Amount <= 0 {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, "Invalid amount")
		return
	}

	if err := ac.BalanceService.CreditAccount(vars["id"], req.Amount); err != nil {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(writer, http.StatusOK, "Credit successful")
}

func (ac *AccountController) FreezeAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	if err := ac.FreezeService.FreezeAccount(vars["id"]); err != nil {
		utils.JSONErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(writer, http.StatusOK, "Account frozen")
}
