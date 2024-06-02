package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
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
		duration := time.Since(client.StartTime) // Calcula a duração desde a primeira solicitação
		seconds := int(duration.Seconds())       // Converte a duração para segundos como um valor inteiro
		msgSave := "\n\nIP do Cliente: " + ip
		msgSave += "\nHorario atual " + time.Now().Format(models.LayoutDate)
		msgSave += "\nMétodo da Requisição: " + r.Method    // Método da requisição (GET, POST, etc.)
		msgSave += "\nURL da Requisição: " + r.URL.String() // URL da requisição
		msgSave += "\nHost da Requisição: " + r.Host        // Host da requisição
		msgSave += "\nNúmero de solicitações do cliente: " + strconv.Itoa(client.NumRequests) + " / " + strconv.Itoa(maxRequests)
		msgSave += "\nDuração desde a primeira solicitação (segundos):" + strconv.Itoa(seconds)

		formattedDateTime := time.Now().Format("2006-01-02")
		logger, err := NewLogger("./Logs/REGISTRO_SIS_" + formattedDateTime + ".txt")
		if err != nil {
			log.Fatal("Erro ao criar logger:", err)
		}
		defer logger.Close()

		log.SetOutput(logger.file)

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
				msgSave += "\n~~~~ RESETOU"

				client.NumRequests = 1
				client.StartTime = time.Now()
				client.LastUptdate = time.Now()
				client.CycleTime = time.Now().Add(cycleTimeout)

			} else {
				client.ExpiredTime.Add(expirationTimeout)
				fmt.Print(msgSave)
				logger.Log(msgSave)
				http.Error(w, ConvertStructError("Por favor, tente novamente mais tarde."), http.StatusTooManyRequests)
				return
			}
		}

		if isMaxRequest == false && isNextCycle == true {
			// reseta
			msgSave += "\n~~~~ RESETOU"
			client.NumRequests = 1
			client.StartTime = time.Now()
			client.LastUptdate = time.Now()
			client.CycleTime = time.Now().Add(cycleTimeout)
		}

		if isMaxRequest == true && isNextCycle == false {
			// travar
			client.ExpiredTime.Add(expirationTimeout)
			fmt.Print(msgSave)
			logger.Log(msgSave)
			http.Error(w, ConvertStructError("Por favor, tente novamente mais tarde."), http.StatusTooManyRequests)
			return
		}
		fmt.Print(msgSave)
		logger.Log(msgSave)

		next.ServeHTTP(w, r)
	})
}

type Logger struct {
	logger *log.Logger
	file   *os.File
}

// NewLogger cria um novo Logger com o arquivo especificado
func NewLogger(filename string) (*Logger, error) {
	// Abre ou cria o arquivo de log
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Cria o logger
	logger := log.New(file, "", log.LstdFlags)

	return &Logger{
		logger: logger,
		file:   file,
	}, nil
}

// Log escreve uma mensagem de log
func (l *Logger) Log(message string) {
	l.logger.Println(message)
}

// Close fecha o arquivo de log e o logger
func (l *Logger) Close() error {
	err := l.file.Close()
	if err != nil {
		return err
	}
	return nil
}
