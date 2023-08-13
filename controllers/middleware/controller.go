package controller_middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"encoding/hex"

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
		return IsValidNumber(value)
	case "id_channel":
		return IsValidNumber(value)
	case "id_receptor":
		return IsValidNumber(value)
	case "end_date":
		return IsValidDateTime(value)
	case "start_date":
		return IsValidDateTime(value)
	case "page":
		return IsValidNumber(value)
	case "limit":
		return IsValidNumber(value)
	case "login":
		return IsValidEmail(value)
	case "emailAgent":
		return IsValidEmail(value)
	case "channel":
		return IsValidTag(value)
	case "name":
		return IsValidString(value)
	case "email":
		return IsValidEmail(value)
	case "date":
		return IsValidDateTime(value)
	case "number":
		return IsValidNumber(value)
	case "bool":
		return IsValidBoolean(value)
	case "tag":
		return IsValidTag(value)
	case "password":
		return IsValidSHA256Key(value)

	default:
		return true // Parâmetro desconhecido, considerado válido
	}
}

func IsValidNumber(value string) bool {
	if len(value) > 20 {
		return false
	}
	_, err := strconv.Atoi(value)
	if err == nil {
		return true // É um número inteiro válido
	}

	regex := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return regex.MatchString(value)
}

func IsValidDateTime(value string) bool {
	if len(value) > 20 {
		return false
	}
	// Verifica se o valor é uma data e hora no formato esperado
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	return regex.MatchString(value)
}

func IsValidEmail(value string) bool {
	// Verifica se o valor é um email válido

	if len(value) > 50 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(value)
}

func IsValidString(value string) bool {
	// Verifica se o valor é uma string sem caracteres especiais
	if len(value) > 50 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return regex.MatchString(value)
}

func IsValidTag(value string) bool {
	// Verifica se o valor é uma string sem caracteres especiais
	if len(value) > 50 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_\s]+$`)
	return regex.MatchString(value)
}

func IsValidBoolean(value string) bool {
	_, err := strconv.ParseBool(value)
	return err == nil
}

func IsValidSHA256Key(value string) bool {
	// Verifica se o valor tem exatamente 64 caracteres hexadecimais
	if len(value) != 64 {
		return false
	}

	regex := regexp.MustCompile(`^[a-fA-F0-9]+$`)
	if !regex.MatchString(value) {
		return false
	}

	// Verifica se o valor pode ser decodificado como bytes
	_, err := hex.DecodeString(value)
	return err == nil
}
