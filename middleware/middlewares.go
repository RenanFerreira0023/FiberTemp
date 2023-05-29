package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/RenanFerreira0023/FiberTemp/models"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateTokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cria um token JWT com a chave secreta
		token := jwt.New(jwt.SigningMethodHS256)
		// Define os claims do token
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = "RenanDutra"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		// Assina o token com a chave secreta
		tokenString, err := token.SignedString([]byte("my-secret-key"))
		if err != nil {
			http.Error(w, "Erro ao criar o token", http.StatusInternalServerError)
			return
		}
		// Adiciona o token como um cabeçalho HTTP
		//	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
		w.Write([]byte(tokenString))
		// Chama o próximo handler na cadeia
		next.ServeHTTP(w, r)
	})
}

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é HMAC com SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
		}
		// Retorna a chave secreta usada para assinar o token
		return []byte("my-secret-key"), nil
	})
	if err != nil {
		return "", fmt.Errorf("Erro ao analisar o token: %v", err)
	}
	// Verifica se o token é válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extrai o nome de usuário do token
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
	}
	return "", fmt.Errorf("Token inválido")
}

func CheckTokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrai o token JWT do cabeçalho de autorização
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {

			//			http.Error(w, "Token de autenticação ausente", http.StatusUnauthorized)
			http.Error(w, ConvertStructError(fmt.Sprintf("Token de autenticação ausente", http.StatusUnauthorized)), http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		// Valida o token JWT e extrai o nome de usuário
		username, err := validateToken(tokenString)
		if err != nil {

			if err != nil {
				http.Error(w, ConvertStructError(err.Error()), http.StatusInternalServerError)
				return
			}
			http.Error(w, ConvertStructError(fmt.Sprintf("Token de autenticação inválido: %v", err)), http.StatusUnauthorized)
			return
		}
		// Adiciona o nome de usuário ao contexto da solicitação
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ConvertStructError(message string) string {

	response := models.MessageError{
		MsgError: fmt.Sprintf(message),
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "{ \"message_error\" : \"Houve um erro super inesperado\"}"
	}
	return string(jsonResponse)
}

var clients = make(map[string]*models.MiddlewareStruct)
var mutex = &sync.Mutex{}

func AntiDDoS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		mutex.Lock()

		if clients[ip] == nil {
			clients[ip] = &models.MiddlewareStruct{
				NumRequests: 1,
				LastRequest: time.Now(),
			}
		} else {
			clients[ip].NumRequests++
			duration := time.Since(clients[ip].LastRequest)
			if duration < 1*time.Second && clients[ip].NumRequests > 10 {
				http.Error(w, ConvertStructError("Too many requests"), http.StatusTooManyRequests)
				mutex.Unlock()
				return
			}
			clients[ip].LastRequest = time.Now()
		}

		mutex.Unlock()

		next.ServeHTTP(w, r)
	})
}
