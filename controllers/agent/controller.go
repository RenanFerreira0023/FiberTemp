package controller_agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	middlewareController "github.com/RenanFerreira0023/FiberTemp/controllers/middleware"

	"github.com/RenanFerreira0023/FiberTemp/models"

	"github.com/RenanFerreira0023/FiberTemp/middleware"
	repositories "github.com/RenanFerreira0023/FiberTemp/repositories/agent"
)

type AgentController struct {
	repository *repositories.AgentRepository
}

func NewAgentController(repository *repositories.AgentRepository) *AgentController {
	return &AgentController{repository: repository}
}

func (a *AgentController) InsertPermissionChannel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyPermission models.QueryBodyInsertPermission
		if err := json.NewDecoder(r.Body).Decode(&bodyPermission); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(bodyPermission.UserReceptorID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(bodyPermission.UserReceptorID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(bodyPermission.ChannelID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(bodyPermission.ChannelID))), http.StatusBadRequest)
			return
		}

		idChannel, err := a.repository.InsertPermissionChannel(bodyPermission)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		fmt.Println("Canal inserido com sucesso ! ", idChannel)

	})
}

func (a *AgentController) SendCopy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recupera tudo do body
		var sendCopyBody models.QueryBodySendCopy
		if err := json.NewDecoder(r.Body).Decode(&sendCopyBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", (sendCopyBody.Symbol)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", fmt.Sprint(sendCopyBody.Symbol))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("tag", (sendCopyBody.ActionType)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", fmt.Sprint(sendCopyBody.ActionType))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.Ticket)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.Ticket))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.Lot)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.Lot))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.TargetPedding)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.TargetPedding))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.TakeProfit)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.TakeProfit))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.StopLoss)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.StopLoss))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", fmt.Sprint(sendCopyBody.DateEntry)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", fmt.Sprint(sendCopyBody.DateEntry))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.UserAgentID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.UserAgentID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(sendCopyBody.ChannelID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(sendCopyBody.ChannelID))), http.StatusBadRequest)
			return
		}

		idCopy, err := a.repository.SendCopy(sendCopyBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		fmt.Println("copy inserida com sucesso ", idCopy)

	})
}

func (a *AgentController) CreateChannel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var channelBody models.QueryBodyCreateChannel
		if err := json.NewDecoder(r.Body).Decode(&channelBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", (channelBody.NameChannel)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", fmt.Sprint(channelBody.NameChannel))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(channelBody.AgentID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(channelBody.AgentID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", (channelBody.CreateChannel)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", fmt.Sprint(channelBody.CreateChannel))), http.StatusBadRequest)
			return
		}

		idChannel, err := a.repository.InsertChannel(channelBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Println("Canal criado com sucesso ID : ", idChannel)

	})
}

func (a *AgentController) InsertAgent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var copyBody models.QueryBodyUsersAgent
		if err := json.NewDecoder(r.Body).Decode(&copyBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", (copyBody.FirstName)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", fmt.Sprint(copyBody.FirstName))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", (copyBody.SecondName)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", fmt.Sprint(copyBody.SecondName))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("email", (copyBody.Email)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "email", fmt.Sprint(copyBody.Email))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", (copyBody.CreateAccount)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", fmt.Sprint(copyBody.CreateAccount))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", (copyBody.ExpiredAccount)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", fmt.Sprint(copyBody.ExpiredAccount))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("bool", fmt.Sprint(copyBody.AccountValid)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "bool", fmt.Sprint(copyBody.AccountValid))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(copyBody.QuantityAlert)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "bool", fmt.Sprint(copyBody.QuantityAlert))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(copyBody.AccountCopy)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(copyBody.AccountCopy))), http.StatusBadRequest)
			return
		}

		inte, err := a.repository.InsertClient(copyBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Println(inte)
		next.ServeHTTP(w, r)
	})
}

func (a *AgentController) CheckURLDatas(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()

		// Percorra cada parâmetro e valide
		for key, values := range params {
			for _, value := range values {
				if !middlewareController.IsValidInput(key, value) {
					http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", key, value)), http.StatusBadRequest)
					return
				}
			}
		}
		next.ServeHTTP(w, r)

	})
}

func (a *AgentController) GetLoginAgent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loginValue := r.URL.Query().Get("login")
		if loginValue == "" {
			http.Error(w, middleware.ConvertStructError("Login não recebido"), http.StatusForbidden)
			return
		}
		channelValue := r.URL.Query().Get("channel")
		if channelValue == "" {
			http.Error(w, middleware.ConvertStructError("Canal não recebido"), http.StatusForbidden)
			return
		}

		agent, err := a.repository.GetisValidLogin(loginValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}
		if time.Now().After(agent[0].ExpiredAccount) {
			http.Error(w, middleware.ConvertStructError("Conta Expirada, Envie um e-mail imediato para appsskilldeveloper@gmail.com para regularizar sua situação"), http.StatusForbidden)
			return
		}

		if !agent[0].AccountValid {
			http.Error(w, middleware.ConvertStructError("Conta Desativada, Envie um e-mail imediato para appsskilldeveloper@gmail.com para regularizar sua situação"), http.StatusForbidden)
			return
		}

		req, err := a.repository.GetAgentFromEmailAndChannel(loginValue, channelValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusNotFound)
			return
		}

		jsonResponse, err := json.Marshal(req)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Trasnformação de json invalido"), http.StatusInternalServerError)
		}

		w.Write([]byte(jsonResponse))
		next.ServeHTTP(w, r)
	})
}

func (a *AgentController) CheckUserExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loginValue := r.URL.Query().Get("login")

		agent, err := a.repository.GetisValidLogin(loginValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusNotFound)
			return
		}

		if time.Now().After(agent[0].ExpiredAccount) {
			http.Error(w, middleware.ConvertStructError("Conta Expirada, Envie um e-mail imediato para appsskilldeveloper@gmail.com para regularizar sua situação"), http.StatusForbidden)
			return
		}

		if !agent[0].AccountValid {
			http.Error(w, middleware.ConvertStructError("Conta Desativada, Envie um e-mail imediato para appsskilldeveloper@gmail.com para regularizar sua situação"), http.StatusForbidden)
			return
		}

		middlewareController.CreateAuthMiddleware(agent[0].ID, next).ServeHTTP(w, r)

		//		next.ServeHTTP(w, r)
	})
}
