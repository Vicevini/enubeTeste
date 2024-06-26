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
	r.Handle("/partnerId", handlers.JwtVerify(http.HandlerFunc(handlers.GetPartnerId))).Methods("GET")
	r.Handle("/partnerName", handlers.JwtVerify(http.HandlerFunc(handlers.GetPartnerName))).Methods("GET")
	r.Handle("/customerId", handlers.JwtVerify(http.HandlerFunc(handlers.GetCustomerId))).Methods("GET")
	r.Handle("/customerName", handlers.JwtVerify(http.HandlerFunc(handlers.GetCustomerName))).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
