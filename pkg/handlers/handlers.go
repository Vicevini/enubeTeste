package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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

	// Autenticação
	if user.Username == "user" && user.Password == "password" {
		// Derruba o token anterior
		if tokenString := getTokenFromCache(user.Username); tokenString != "" {
			RevokeToken(tokenString)
		}

		// Define as claims do novo token
		expirationTime := time.Now().Add(1 * time.Hour) // Expira em 1 hora
		claims := jwt.MapClaims{
			"username": user.Username,
			"exp":      expirationTime.Unix(),
		}

		// Cria o novo token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println("Erro ao assinar o token:", err)
			http.Error(w, "Erro ao processar a requisição", http.StatusInternalServerError)
			return
		}

		log.Println("Token gerado com sucesso:", tokenString)

		// Guarda o novo token em cache
		saveTokenInCache(user.Username, tokenString)

		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	} else {
		log.Println("Usuário ou senha inválidos")
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
	}
}

// Funções auxiliares para gerenciar o cache de tokens
func saveTokenInCache(username, tokenString string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	tokenCache[username] = tokenString
}

func getTokenFromCache(username string) string {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	return tokenCache[username]
}

// Funções para ler colunas do Excel

func GetPartnerId(w http.ResponseWriter, r *http.Request) {
	partnerId, err := readTrueColumn("PartnerId")
	if err != nil {
		log.Println("Erro ao ler a coluna PartnerId:", err)
		http.Error(w, "Erro ao ler a coluna PartnerId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(partnerId)
}

func GetPartnerName(w http.ResponseWriter, r *http.Request) {
	partnerName, err := readTrueColumn("PartnerName")
	if err != nil {
		log.Println("Erro ao ler a coluna PartnerName:", err)
		http.Error(w, "Erro ao ler a coluna PartnerName", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(partnerName)
}

func GetCustomerId(w http.ResponseWriter, r *http.Request) {
	CustomerId, err := readTrueColumn("CustomerId")
	if err != nil {
		log.Println("Erro ao ler a coluna CustomerId:", err)
		http.Error(w, "Erro ao ler a coluna CustomerId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(CustomerId)
}

func GetCustomerName(w http.ResponseWriter, r *http.Request) {
	customerId, err := readTrueColumn("CustomerName")
	if err != nil {
		log.Println("Erro ao ler a coluna CustomerName:", err)
		http.Error(w, "Erro ao ler a coluna CustomerName", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customerId)
}
