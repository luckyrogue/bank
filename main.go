package main

import (
	"bank/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
