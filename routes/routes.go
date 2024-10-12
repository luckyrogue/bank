package routes

import (
	"bank/controllers"
	"bank/services"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	accountService := services.NewAccountService()
	balanceService := services.NewBalanceService(accountService.Accounts)
	freezeService := services.NewFreezeService(accountService.Accounts)

	accountController := controllers.NewAccountController(accountService, balanceService, freezeService)
	transactionController := controllers.NewTransactionController(balanceService)

	RegisterAccountRoutes(router, accountController)
	RegisterTransactionRoutes(router, transactionController)

	return router
}
