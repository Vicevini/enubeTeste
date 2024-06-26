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

	// Endpoints protegidos
	s := r.PathPrefix("/api").Subrouter()
	s.Use(handlers.JwtVerify)
	s.HandleFunc("/data", handlers.GetData).Methods("GET")

	// Endpoints para colunas do Excel
	r.HandleFunc("/names", handlers.GetNames).Methods("GET")
	r.HandleFunc("/ages", handlers.GetAges).Methods("GET")
	r.HandleFunc("/schools", handlers.GetSchools).Methods("GET")
	r.HandleFunc("/partnerId", handlers.GetPartnerId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
