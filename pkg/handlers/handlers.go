package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string `json:"token"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Erro ao decodificar o corpo da requisição:", err)
		http.Error(w, "Erro ao processar a requisição", http.StatusBadRequest)
		return
	}

	log.Println("Credenciais recebidas:", user.Username, user.Password)

	if user.Username == "user" && user.Password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour).Unix(), // Token válido por 1 hora
		})

		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println("Erro ao assinar o token:", err)
			http.Error(w, "Erro ao processar a requisição", http.StatusInternalServerError)
			return
		}

		log.Println("Token gerado com sucesso:", tokenString)
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	} else {
		log.Println("Usuário ou senha inválidos")
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
	}
}

func getPageParams(r *http.Request) (int, int) {
	pageSize := 50
	pageNumber := 1

	query := r.URL.Query()
	if size := query.Get("tamanhoPagina"); size != "" {
		if s, err := strconv.Atoi(size); err == nil {
			pageSize = s
		}
	}
	if number := query.Get("numeroPagina"); number != "" {
		if n, err := strconv.Atoi(number); err == nil {
			pageNumber = n
		}
	}

	return pageSize, pageNumber
}

func GetPartnerId(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	partnerId, err := readTrueColumn("PartnerId", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna PartnerId:", err)
		http.Error(w, "Erro ao ler a coluna PartnerId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(partnerId)
}

func GetPartnerName(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	partnerName, err := readTrueColumn("PartnerName", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna PartnerName:", err)
		http.Error(w, "Erro ao ler a coluna PartnerName", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(partnerName)
}

func GetCustomerId(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	customerId, err := readTrueColumn("CustomerId", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna CustomerId:", err)
		http.Error(w, "Erro ao ler a coluna CustomerId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customerId)
}

func GetCustomerName(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	customerName, err := readTrueColumn("CustomerName", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna CustomerName:", err)
		http.Error(w, "Erro ao ler a coluna CustomerName", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customerName)
}

func GetInvoiceNumber(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	invoiceNumber, err := readTrueColumn("InvoiceNumber", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna InvoiceNumber:", err)
		http.Error(w, "Erro ao ler a coluna InvoiceNumber", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(invoiceNumber)
}

func GetProductId(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	productId, err := readTrueColumn("productId", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna productId:", err)
		http.Error(w, "Erro ao ler a coluna productId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(productId)
}
func GetSkuId(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	skuId, err := readTrueColumn("skuId", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna skuId:", err)
		http.Error(w, "Erro ao ler a coluna skuId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(skuId)
}
func GetSkuName(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	skuName, err := readTrueColumn("skuName", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna skuName:", err)
		http.Error(w, "Erro ao ler a coluna skuName", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(skuName)
}
func GetSubscriptionId(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	subscriptionId, err := readTrueColumn("subscriptionId", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna subscriptionId:", err)
		http.Error(w, "Erro ao ler a coluna subscriptionId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(subscriptionId)
}
func GetChargeStartDate(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	chargeStartDate, err := readTrueColumn("chargeStartDate", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna chargeStartDate:", err)
		http.Error(w, "Erro ao ler a coluna chargeStartDate", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chargeStartDate)
}
func GetChargeEndDate(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	chargeEndDate, err := readTrueColumn("chargeEndDate", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna chargeEndDate:", err)
		http.Error(w, "Erro ao ler a coluna chargeEndDate", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chargeEndDate)
}
func GetBillingPreTaxTotal(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	billingPreTaxTotal, err := readTrueColumn("billingPreTaxTotal", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna billingPreTaxTotal:", err)
		http.Error(w, "Erro ao ler a coluna billingPreTaxTotal", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(billingPreTaxTotal)
}
func GetPricingPreTaxTotal(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	pricingPreTaxTotal, err := readTrueColumn("pricingPreTaxTotal", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna pricingPreTaxTotal:", err)
		http.Error(w, "Erro ao ler a coluna pricingPreTaxTotal", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pricingPreTaxTotal)
}
func GetUsageDate(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	usageDate, err := readTrueColumn("usageDate", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna usageDate:", err)
		http.Error(w, "Erro ao ler a coluna usageDate", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usageDate)
}
func GetConsumedService(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	consumedService, err := readTrueColumn("consumedService", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna consumedService:", err)
		http.Error(w, "Erro ao ler a coluna consumedService", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(consumedService)
}
func GetResourceGroup(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	resourceGroup, err := readTrueColumn("resourceGroup", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna resourceGroup:", err)
		http.Error(w, "Erro ao ler a coluna resourceGroup", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resourceGroup)
}
func GetResourceURI(w http.ResponseWriter, r *http.Request) {
	pageSize, pageNumber := getPageParams(r)
	resourceURI, err := readTrueColumn("resourceURI", pageSize, pageNumber)
	if err != nil {
		log.Println("Erro ao ler a coluna resourceURI:", err)
		http.Error(w, "Erro ao ler a coluna resourceURI", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resourceURI)
}
