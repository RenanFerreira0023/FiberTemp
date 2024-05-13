package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"time"

	"github.com/RenanFerreira0023/FiberTemp/models"
	"github.com/joho/godotenv"

	"github.com/golang-jwt/jwt/v5"
)

func CreateTokenHandler(ID int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errEnv := godotenv.Load()
		if errEnv != nil {
			fmt.Println("Error loading .env file  ", errEnv.Error())
		}
		SECRET_KEY := os.Getenv("SECRET_KEY_TOKEN")
		// Cria um token JWT com a chave secreta
		token := jwt.New(jwt.SigningMethodHS256)
		// Define os claims do token
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = "RDSTDR"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		// Assina o token com a chave secreta
		tokenString, err := token.SignedString([]byte(SECRET_KEY))
		if err != nil {
			http.Error(w, "Erro ao criar o token", http.StatusInternalServerError)
			return
		}
		var queryRequestToken models.QueryRequestToken
		queryRequestToken.UserID = ID
		queryRequestToken.Token = tokenString
		jsonResponse, err := json.Marshal(queryRequestToken)
		if err != nil {
			http.Error(w, ("Trasnformação de json invalido"), http.StatusInternalServerError)
		}
		w.Write([]byte(jsonResponse))
		next.ServeHTTP(w, r)
	})
}

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		errEnv := godotenv.Load()
		if errEnv != nil {
			fmt.Println("Error loading .env file  ", errEnv.Error())
		}
		SECRET_KEY := os.Getenv("SECRET_KEY_TOKEN")

		// Verifica se o método de assinatura é HMAC com SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
		}
		// Retorna a chave secreta usada para assinar o token
		return []byte(SECRET_KEY), nil
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

var clients sync.Map

func AntiDDoS(next http.Handler) http.Handler {
	//////////////////////////////////
	//////// SETUP ///////////////////
	cycleTimeout := 1 * time.Minute      // Tempo de expiração para novas requisições
	maxRequests := 350                   // Numero de requisição permitidas por minuto
	expirationTimeout := 5 * time.Minute // Tempo para redefinir o contador de solicitações
	//////////////////////////////////
	//////////////////////////////////

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr) // Obtém apenas o endereço IP, ignorando a porta

		value, _ := clients.LoadOrStore(ip, &models.MiddlewareStruct{
			NumRequests: 1,
			StartTime:   time.Now(),                   // horario da primeira solicitacao
			LastUptdate: time.Now(),                   // horario da ultima atualização
			CycleTime:   time.Now().Add(cycleTimeout), // Horario da dutração de um ciclo
			ExpiredTime: time.Now().Add(expirationTimeout),
		})
		client := value.(*models.MiddlewareStruct)

		client.LastUptdate = time.Now()
		client.NumRequests++

		// Log das informações do cliente e da duração desde a primeira solicitação
		/*
			duration := time.Since(client.StartTime) // Calcula a duração desde a primeira solicitação
			fmt.Println("\nIP do Cliente:", ip)
			fmt.Println("Número de solicitações do cliente:", client.NumRequests, " / ", maxRequests)
			fmt.Println("Duração desde a primeira solicitação (segundos):", int(duration.Seconds()))
			fmt.Println("duration.Seconds() ~~~~ expirationTimeout.Seconds()", duration.Seconds(), "		~~~~~  ", expirationTimeout.Seconds())
		*/

		// controle numero requisição
		isMaxRequest := false
		if client.NumRequests > maxRequests {
			isMaxRequest = true
		}

		// controle tempo
		isNextCycle := false
		if time.Now().After(client.CycleTime) {
			isNextCycle = true
		}

		if isMaxRequest == true && isNextCycle == true {
			if time.Now().After(client.ExpiredTime) {
				//				fmt.Println("~~~~ RESETOU")
				client.NumRequests = 1
				client.StartTime = time.Now()
				client.LastUptdate = time.Now()
				client.CycleTime = time.Now().Add(cycleTimeout)

			} else {
				client.ExpiredTime.Add(expirationTimeout)
				http.Error(w, ConvertStructError("Por favor, tente novamente mais tarde."), http.StatusTooManyRequests)
				return
			}
		}

		if isMaxRequest == false && isNextCycle == true {
			// reseta
			//			fmt.Println("~~~~ RESETOU")
			client.NumRequests = 1
			client.StartTime = time.Now()
			client.LastUptdate = time.Now()
			client.CycleTime = time.Now().Add(cycleTimeout)
		}

		if isMaxRequest == true && isNextCycle == false {
			// travar
			http.Error(w, ConvertStructError("Por favor, tente novamente mais tarde."), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
