package controller_middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/RenanFerreira0023/FiberTemp/middleware"
	"github.com/RenanFerreira0023/FiberTemp/models"
)

func CreateAuthMiddleware(ID int, next http.Handler) http.Handler {
	return (middleware.CreateTokenHandler(ID, next))
}

func CheckValidToken(next http.Handler) http.Handler {
	return (middleware.CheckTokenHandler(next))
}

func CheckAntiDDoS(next http.Handler) http.Handler {
	return middleware.AntiDDoS(next)
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

func IsValidInput(key, value string) bool {
	// Aqui você pode adicionar as validações adequadas para cada tipo de dado
	switch key {
	case "id_agent":
		return isValidNumber(value)
	case "id_channel":
		return isValidNumber(value)
	case "id_receptor":
		return isValidNumber(value)
	case "end_date":
		return isValidDateTime(value)
	case "start_date":
		return isValidDateTime(value)
	case "page":
		return isValidNumber(value)
	case "limit":
		return isValidNumber(value)
	case "login":
		return isValidEmail(value)
	case "emailAgent":
		return isValidEmail(value)
	case "channel":
		return isValidString(value)
	default:
		return false // Parâmetro desconhecido, considerado válido
	}
}

func isValidNumber(value string) bool {
	if len(value) > 20 {
		return false
	}
	// Verifica se o valor é um número
	regex := regexp.MustCompile(`^-?\d+$`)
	return regex.MatchString(value)
}

func isValidDateTime(value string) bool {
	if len(value) > 20 {
		return false
	}
	// Verifica se o valor é uma data e hora no formato esperado
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	return regex.MatchString(value)
}

func isValidEmail(value string) bool {
	// Verifica se o valor é um email válido

	if len(value) > 50 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(value)
}

func isValidString(value string) bool {
	// Verifica se o valor é uma string sem caracteres especiais
	if len(value) > 50 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return regex.MatchString(value)
}
