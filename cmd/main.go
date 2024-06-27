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

	// Endpoints para colunas do Excel com JWT Middleware
	r.Handle("/partnerId", handlers.JwtVerify(http.HandlerFunc(handlers.GetPartnerId))).Methods("GET")
	r.Handle("/partnerName", handlers.JwtVerify(http.HandlerFunc(handlers.GetPartnerName))).Methods("GET")
	r.Handle("/customerId", handlers.JwtVerify(http.HandlerFunc(handlers.GetCustomerId))).Methods("GET")
	r.Handle("/customerName", handlers.JwtVerify(http.HandlerFunc(handlers.GetCustomerName))).Methods("GET")
	r.Handle("/invoiceNumber", handlers.JwtVerify(http.HandlerFunc(handlers.GetInvoiceNumber))).Methods("GET")
	r.Handle("/productId", handlers.JwtVerify(http.HandlerFunc(handlers.GetProductId))).Methods("GET")
	r.Handle("/skuId", handlers.JwtVerify(http.HandlerFunc(handlers.GetSkuId))).Methods("GET")
	r.Handle("/skuName", handlers.JwtVerify(http.HandlerFunc(handlers.GetSkuName))).Methods("GET")
	r.Handle("/subscriptionId", handlers.JwtVerify(http.HandlerFunc(handlers.GetSubscriptionId))).Methods("GET")
	r.Handle("/chargeStartDate", handlers.JwtVerify(http.HandlerFunc(handlers.GetChargeStartDate))).Methods("GET")
	r.Handle("/chargeEndDate", handlers.JwtVerify(http.HandlerFunc(handlers.GetChargeEndDate))).Methods("GET")
	r.Handle("/billingPreTaxTotal", handlers.JwtVerify(http.HandlerFunc(handlers.GetBillingPreTaxTotal))).Methods("GET")
	r.Handle("/pricingPreTaxTotal", handlers.JwtVerify(http.HandlerFunc(handlers.GetPricingPreTaxTotal))).Methods("GET")
	r.Handle("/usageDate", handlers.JwtVerify(http.HandlerFunc(handlers.GetUsageDate))).Methods("GET")
	r.Handle("/consumedService", handlers.JwtVerify(http.HandlerFunc(handlers.GetConsumedService))).Methods("GET")
	r.Handle("/resourceGroup", handlers.JwtVerify(http.HandlerFunc(handlers.GetResourceGroup))).Methods("GET")
	r.Handle("/resourceURI", handlers.JwtVerify(http.HandlerFunc(handlers.GetResourceURI))).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
