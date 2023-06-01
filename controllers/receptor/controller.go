package controller_receptor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"
	"time"

	middlewareController "github.com/RenanFerreira0023/FiberTemp/controllers/middleware"
	"github.com/RenanFerreira0023/FiberTemp/middleware"
	"github.com/RenanFerreira0023/FiberTemp/models"
	repositories "github.com/RenanFerreira0023/FiberTemp/repositories/receptor"
)

type ReceptorController struct {
	repository *repositories.ReceptorRepository
}

func NewReceptorController(repository *repositories.ReceptorRepository) *ReceptorController {
	return &ReceptorController{repository: repository}
}

func (c *ReceptorController) SendReqCopy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var reqCopyBody models.QueryRequestReqCopy
		if err := json.NewDecoder(r.Body).Decode(&reqCopyBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(reqCopyBody.AllCopyID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(reqCopyBody.AllCopyID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(reqCopyBody.ReceptorID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(reqCopyBody.ReceptorID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("number", fmt.Sprint(reqCopyBody.ChannelID)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "number", fmt.Sprint(reqCopyBody.ChannelID))), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", fmt.Sprint(reqCopyBody.DateSendOrder)) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", reqCopyBody.DateSendOrder)), http.StatusBadRequest)
			return
		}

		idReqCopy, err := c.repository.InsertReqCopy(reqCopyBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}

		var strReq200 models.JsonRequest200
		strReq200.DataBaseID = idReqCopy
		strReq200.MsgInsert = "Requisição inserida com sucesso"
		jsonResponse, err := json.Marshal(strReq200)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Trasnformação de json invalido"), http.StatusInternalServerError)
		}
		w.Write([]byte(jsonResponse))

		next.ServeHTTP(w, r)

	})
}

func (c *ReceptorController) GetCopy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agentID := r.URL.Query().Get("id_agent")
		agentIDstr, err := strconv.Atoi(agentID)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Erro ao converter a string para int:"+err.Error()), http.StatusInternalServerError)
			return
		}

		channelID := r.URL.Query().Get("id_channel")
		channelIDstr, err := strconv.Atoi(channelID)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Erro ao converter a string para int:"+err.Error()), http.StatusInternalServerError)
			return
		}

		receptorID := r.URL.Query().Get("id_receptor")
		receptorIDstr, err := strconv.Atoi(receptorID)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Erro ao converter a string para int:"+err.Error()), http.StatusInternalServerError)
			return
		}

		//	receptorID := r.URL.Query().Get("id_receptor")
		startDateStr := r.URL.Query().Get("start_date")
		endDateStr := r.URL.Query().Get("end_date")

		offset := r.URL.Query().Get("page")
		offsetStr, err := strconv.Atoi(offset)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Erro ao converter a string para int:"+err.Error()), http.StatusInternalServerError)
			return
		}

		limitPage := r.URL.Query().Get("limit")
		limitPageStr, err := strconv.Atoi(limitPage)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Erro ao converter a string para int:"+err.Error()), http.StatusInternalServerError)
			return
		}

		var structURL models.StrutcURLCopyTrader
		structURL.AgentID = agentIDstr
		structURL.ChannelID = channelIDstr
		structURL.DateStart = startDateStr
		structURL.DateEnd = endDateStr
		structURL.Offset = offsetStr
		structURL.PageLimit = limitPageStr
		//////////////////////
		//// GET a lista de copys
		//////////////////////
		requestCopyTrader, err := c.repository.GetCopyTrader(structURL)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}

		//////////////////////
		//// Loop para conferir se  a lista ja foi executada
		//////////////////////
		var countLoop = len(requestCopyTrader)
		var arrayBodyResponde []models.BodyCopyTrader
		for i := 0; i < (countLoop); i++ {
			errReq := c.repository.CheckReqCopy(requestCopyTrader[i].ChannelID, receptorIDstr, requestCopyTrader[i].ID, requestCopyTrader[i].DateEntry.Format(models.LayoutDate))
			if errReq == nil {
				arrayBodyResponde = append(arrayBodyResponde, requestCopyTrader[i])
			}
		}

		if len(arrayBodyResponde) > 0 {
			jsonResponse, err := json.Marshal(arrayBodyResponde)
			if err != nil {
				http.Error(w, middleware.ConvertStructError("Trasnformação de json invalido"), http.StatusInternalServerError)
			}
			w.Write([]byte(jsonResponse))
		} else {
			http.Error(w, middleware.ConvertStructError("Sem dados para retornar"), http.StatusNotFound)
		}
		next.ServeHTTP(w, r)

	})
}

func (c *ReceptorController) GetLoginReceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginValue := r.URL.Query().Get("login")
		if loginValue == "" {
			http.Error(w, middleware.ConvertStructError("Login não recebido"), http.StatusForbidden)
			return
		}
		//
		emailAgentValue := r.URL.Query().Get("emailAgent")
		if emailAgentValue == "" {
			http.Error(w, middleware.ConvertStructError("Email do agente não recebido"), http.StatusForbidden)
			return
		}

		channelValue := r.URL.Query().Get("channel")
		if channelValue == "" {
			http.Error(w, middleware.ConvertStructError("Canal não recebido"), http.StatusForbidden)
			return
		}

		//////////////////////
		//// Get valid receptor
		//////////////////////
		request, err := c.repository.GetValidReceptor(loginValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusNotFound)
			return
		}
		var receptorID = request[0].ID
		var dateExpired = request[0].ExpiredAccount
		if time.Now().After(dateExpired) {
			http.Error(w, middleware.ConvertStructError("Conta Expirada, Envie um e-mail imediato com o agente provedor do sinal para regularizar sua situação"), http.StatusForbidden)
			return
		}

		//////////////////////
		//// Get id agent
		//////////////////////
		idAgent, err := c.repository.GetAgentID(emailAgentValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Agente não encontrado"), http.StatusNotFound)
			return
		}

		//////////////////////
		//// Get channel
		//////////////////////
		requestChannel, err := c.repository.GetChannel(channelValue, idAgent)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Canal não encontrado"), http.StatusNotFound)
			return
		}
		var channelID = requestChannel.ID
		var agentID = requestChannel.AgentID

		//////////////////////
		//// Get Permission
		//////////////////////
		err = c.repository.GetPermissionChannel(channelID, receptorID)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Você não tem permissão para acessar esse canal"), http.StatusNotFound)
			return
		}
		var modelLogin models.QueryRequestLoginReceptor
		modelLogin.AgentID = agentID
		modelLogin.ReceptorID = receptorID
		modelLogin.ChannelID = channelID

		jsonResponse, err := json.Marshal(modelLogin)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Trasnformação de json invalido"), http.StatusInternalServerError)
		}
		w.Write([]byte(jsonResponse))
		next.ServeHTTP(w, r)
	})
}

func (c *ReceptorController) CheckURLDatas(next http.Handler) http.Handler {
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
func (c *ReceptorController) CheckUserExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loginValue := r.URL.Query().Get("login")

		receptor, err := c.repository.GetisValidLogin(loginValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}

		if time.Now().After(receptor[0].ExpiredAccount) {
			http.Error(w, middleware.ConvertStructError("Conta Expirada, Envie um e-mail imediato com o agente provedor do sinal para regularizar sua situação"), http.StatusForbidden)
			return
		}
		middlewareController.CreateAuthMiddleware(receptor[0].ID, next).
			ServeHTTP(w, r)
	})
}

func (c *ReceptorController) InsertReceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var receptorBody models.QueryGetUserReceptor
		if err := json.NewDecoder(r.Body).Decode(&receptorBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", receptorBody.FirstName) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", receptorBody.FirstName)), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("name", receptorBody.SecondName) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "name", receptorBody.SecondName)), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("email", receptorBody.Email) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "email", receptorBody.Email)), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", receptorBody.CreateAccount) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", receptorBody.CreateAccount)), http.StatusBadRequest)
			return
		}

		if !middlewareController.IsValidInput("date", receptorBody.ExpiredAccount) {
			http.Error(w, middlewareController.ConvertStructError(fmt.Sprintf("Valor inválido para o parâmetro '%s': %s", "date", receptorBody.ExpiredAccount)), http.StatusBadRequest)
			return
		}

		idReceptor, err := c.repository.InsertClient(receptorBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}

		var strReq200 models.JsonRequest200
		strReq200.DataBaseID = idReceptor
		strReq200.MsgInsert = "Receptor [ " + receptorBody.Email + " ] criado com sucesso"
		jsonResponse, err := json.Marshal(strReq200)
		if err != nil {
			http.Error(w, middleware.ConvertStructError("Trasnformação de json invalido"), http.StatusInternalServerError)
		}
		w.Write([]byte(jsonResponse))

		next.ServeHTTP(w, r)
	})
}
