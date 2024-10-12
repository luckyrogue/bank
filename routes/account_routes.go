package routes

import (
	"bank/controllers"
	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(router *mux.Router, accountController *controllers.AccountController) {
	router.HandleFunc("/accounts", accountController.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", accountController.GetAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}", accountController.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", accountController.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/accounts/{id}/debit", accountController.DebitAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}/credit", accountController.CreditAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}/freeze", accountController.FreezeAccount).Methods("POST")
}
