package controller_middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RenanFerreira0023/FiberTemp/middleware"
	"github.com/RenanFerreira0023/FiberTemp/models"
)

func CreateAuthMiddleware(next http.Handler) http.Handler {
	return (middleware.CreateTokenHandler(next))
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
