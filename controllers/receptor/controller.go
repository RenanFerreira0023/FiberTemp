package controller_receptor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"../../middleware"
	"../../models"
	repositories "../../repositories/receptor"
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

		requestDateSendOrder := c.checkDatas("datetime", reqCopyBody.DateSendOrder)
		if requestDateSendOrder != "" {
			http.Error(w, middleware.ConvertStructError("Data resposta da copy \n"+requestDateSendOrder), http.StatusBadRequest)
			return
		}

		requestInsert, err := c.repository.InsertReqCopy(reqCopyBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Print("requestInsert     ", requestInsert)
		next.ServeHTTP(w, r)

	})
}

func (c *ReceptorController) GetCopy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agentID := r.URL.Query().Get("id_agent")
		agentIDstr, err := strconv.Atoi(agentID)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
			return
		}

		channelID := r.URL.Query().Get("id_channel")
		channelIDstr, err := strconv.Atoi(channelID)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
			return
		}

		receptorID := r.URL.Query().Get("id_receptor")
		receptorIDstr, err := strconv.Atoi(receptorID)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
			return
		}

		//	receptorID := r.URL.Query().Get("id_receptor")
		startDateStr := r.URL.Query().Get("start_date")
		endDateStr := r.URL.Query().Get("end_date")

		offset := r.URL.Query().Get("page")
		offsetStr, err := strconv.Atoi(offset)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
			return
		}

		limitPage := r.URL.Query().Get("limit")
		limitPageStr, err := strconv.Atoi(limitPage)
		if err != nil {
			fmt.Println("Erro ao converter a string para int:", err)
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
		loginValue := r.URL.Query().Get("login")
		datasLogin := c.checkDatas("email", loginValue)
		if datasLogin != "" && loginValue != "" {
			http.Error(w, "Dados da URL inválidos [ login ] ", http.StatusInternalServerError)
			return
		}

		emailAgentValue := r.URL.Query().Get("emailAgent")
		emailAgent := c.checkDatas("email", emailAgentValue)
		if emailAgent != "" && emailAgentValue != "" {
			http.Error(w, "Dados da URL inválidos [ emailAgent ] ", http.StatusInternalServerError)
			return
		}

		startDateStr := r.URL.Query().Get("start_date")
		dateStart := c.checkDatas("datetime", startDateStr)
		if dateStart != "" && startDateStr != "" {
			http.Error(w, "Dados da URL inválidos [ start_date ]", http.StatusBadRequest)
			return
		}

		endDateStr := r.URL.Query().Get("end_date")
		dateEnd := c.checkDatas("datetime", endDateStr)
		if dateEnd != "" && endDateStr != "" {
			http.Error(w, "Dados da URL inválidos [ end_date ]", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *ReceptorController) checkDatas(typeValid string, data string) string {
	if len(data) > 200 {
		return "Muitos caracteres na requisição"
	}
	if typeValid == "email" {
		regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !regex.MatchString(data) {
			return "Formato do e-mail inválido"
		}
	}
	if typeValid == "datetime" {
		_, err := time.Parse(models.LayoutDate, data)
		if err != nil {
			return ("A string não está no formato de data correto.")
		}
	}
	return ""
}

func (c *ReceptorController) CheckUserExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loginValue := r.URL.Query().Get("login")
		datasLogin := c.checkDatas("email", loginValue)
		if datasLogin != "" {
			http.Error(w, "Dados da URL inválidos", http.StatusBadRequest)
			return
		}

		receptor, err := c.repository.GetisValidLogin(loginValue)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusInternalServerError)
			return
		}

		if time.Now().After(receptor[0].ExpiredAccount) {
			http.Error(w, middleware.ConvertStructError("Conta Expirada, Envie um e-mail imediato com o agente provedor do sinal para regularizar sua situação"), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *ReceptorController) InsertReceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var receptorBody models.QueryGetUserReceptor
		if err := json.NewDecoder(r.Body).Decode(&receptorBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		//		fmt.Printf("Dados do Body: %+v\n", copyBody)
		requestFirstName := c.checkDatas("", receptorBody.FirstName)
		if requestFirstName != "" {
			http.Error(w, middleware.ConvertStructError("Primeiro nome \n"+requestFirstName), http.StatusBadRequest)
			return
		}

		requestSecondName := c.checkDatas("", receptorBody.SecondName)
		if requestSecondName != "" {
			http.Error(w, middleware.ConvertStructError("Segundo nome \n"+requestSecondName), http.StatusBadRequest)
			return
		}

		requestEmail := c.checkDatas("email", receptorBody.Email)
		if requestEmail != "" {
			http.Error(w, middleware.ConvertStructError("Email \n"+requestEmail), http.StatusBadRequest)
			return
		}

		requestDtCreateAccount := c.checkDatas("datetime", receptorBody.CreateAccount)
		if requestDtCreateAccount != "" {
			http.Error(w, middleware.ConvertStructError("Data Criação \n"+requestDtCreateAccount), http.StatusBadRequest)
			return
		}

		requestDtExpiredAccount := c.checkDatas("datetime", receptorBody.ExpiredAccount)
		if requestDtExpiredAccount != "" {
			http.Error(w, middleware.ConvertStructError("Data Expiração \n"+requestDtExpiredAccount), http.StatusBadRequest)
			return
		}

		idInsert, err := c.repository.InsertClient(receptorBody)
		if err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Println(idInsert)
		next.ServeHTTP(w, r)
	})
}
