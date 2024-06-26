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

	// Simulação de autenticação
	if user.Username == "user" && user.Password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
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

func GetData(w http.ResponseWriter, r *http.Request) {
	// Retorno de dados de exemplo
	data := map[string]string{
		"data": "Este é um dado protegido",
	}
	json.NewEncoder(w).Encode(data)
}

// Funções para ler colunas do Excel

func GetNames(w http.ResponseWriter, r *http.Request) {
	names, err := readColumn("Nome")
	if err != nil {
		log.Println("Erro ao ler a coluna Nome:", err)
		http.Error(w, "Erro ao ler a coluna Nome", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(names)
}

func GetAges(w http.ResponseWriter, r *http.Request) {
	ages, err := readColumn("Idade")
	if err != nil {
		log.Println("Erro ao ler a coluna Idade:", err)
		http.Error(w, "Erro ao ler a coluna Idade", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ages)
}

func GetSchools(w http.ResponseWriter, r *http.Request) {
	schools, err := readColumn("Escola")
	if err != nil {
		log.Println("Erro ao ler a coluna Escola:", err)
		http.Error(w, "Erro ao ler a coluna Escola", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schools)
}

func GetPartnerId(w http.ResponseWriter, r *http.Request) {
	partnerId, err := readTrueColumn("PartnerId")
	if err != nil {
		log.Println("Erro ao ler a coluna PartnerId:", err)
		http.Error(w, "Erro ao ler a coluna PartnerId", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(partnerId)
}
