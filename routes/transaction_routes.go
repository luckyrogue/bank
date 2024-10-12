package routes

import (
	"bank/controllers"
	"github.com/gorilla/mux"
)

func RegisterTransactionRoutes(router *mux.Router, balanceService *controllers.TransactionController) {
	transactionController := balanceService

	router.HandleFunc("/transfer", transactionController.TransferFunds).Methods("POST")
}
