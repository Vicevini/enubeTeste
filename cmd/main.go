package main

import (
	"enubeTeste/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoints de autenticação
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Endpoints para colunas do Excel
	r.HandleFunc("/partnerId", handlers.GetPartnerId).Methods("GET")
	r.HandleFunc("/partnerName", handlers.GetPartnerName).Methods("GET")
	r.HandleFunc("/customerId", handlers.GetCustomerId).Methods("GET")
	r.HandleFunc("/customerId", handlers.GetCustomerName).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
