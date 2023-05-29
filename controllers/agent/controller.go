package controller_agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

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

func (a *AgentController) SendCopy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recupera tudo do body
		var sendCopyBody models.QueryBodySendCopy
		if err := json.NewDecoder(r.Body).Decode(&sendCopyBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		symbol := sendCopyBody.Symbol
		requestSymbol := a.checkDatas("", symbol)
		if requestSymbol != "" {
			http.Error(w, middleware.ConvertStructError("Symbol \n"+requestSymbol), http.StatusBadRequest)
			return
		}

		actionType := sendCopyBody.ActionType
		requestActionType := a.checkDatas("", actionType)
		if requestActionType != "" {
			http.Error(w, middleware.ConvertStructError("Action Type \n"+requestActionType), http.StatusBadRequest)
			return
		}

		ticket := sendCopyBody.Ticket
		requestTicket := a.checkDatas("", string(ticket))
		if requestTicket != "" {
			http.Error(w, middleware.ConvertStructError("Ticket \n"+requestTicket), http.StatusBadRequest)
			return
		}

		//		lot := sendCopyBody.Lot
		//		requestLot := a.checkDatas("", (lot))
		//		if requestLot != "" {
		//			http.Error(w, middleware.ConvertStructError("Lote \n"+requestLot), http.StatusBadRequest)
		//			return
		//		}

		//		targetPedding := sendCopyBody.TargetPedding
		//		takeprofit := sendCopyBody.TakeProfit
		//		stoploss := sendCopyBody.StopLoss
		dtEntry := sendCopyBody.DateEntry
		requesDateEntry := a.checkDatas("datetime", (dtEntry))
		if requesDateEntry != "" {
			http.Error(w, middleware.ConvertStructError("Data Envio ordem \n"+requesDateEntry), http.StatusBadRequest)
			return
		}
		//		agentID := sendCopyBody.UserAgentID
		//		channelID := sendCopyBody.ChannelID

		// verifica se essa copy ja foi inserida ( compare a data de envio + tipo da acao)
		// grava a copy no banco

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
		fmt.Println("AAAAAAAAAAAAAAAAAAAAA")
		var channelBody models.QueryBodyCreateChannel
		if err := json.NewDecoder(r.Body).Decode(&channelBody); err != nil {
			http.Error(w, middleware.ConvertStructError(err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Println("bbbbbbbbbbbbb")

		requestChannelName := a.checkDatas("", channelBody.NameChannel)
		if requestChannelName != "" {
			http.Error(w, middleware.ConvertStructError("Nome do canal \n"+requestChannelName), http.StatusBadRequest)
			return
		}
		fmt.Println("cccccccccccccc")

		requestAgentID := a.checkDatas("", strconv.Itoa(channelBody.AgentID))
		if requestAgentID != "" {
			http.Error(w, middleware.ConvertStructError("Quantidade de cópias \n"+requestAgentID), http.StatusBadRequest)
			return
		}

		fmt.Println("dddddddddddddd")

		requestDtCreateChannel := a.checkDatas("datetime", channelBody.CreateChannel)
		if requestDtCreateChannel != "" {
			http.Error(w, middleware.ConvertStructError("Data Criação do canal \n"+requestDtCreateChannel), http.StatusBadRequest)
			return
		}
		fmt.Println("eeeeeeeeeeeeeeeeeee")

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
		//		fmt.Printf("Dados do Body: %+v\n", copyBody)
		requestFirstName := a.checkDatas("", copyBody.FirstName)
		if requestFirstName != "" {
			http.Error(w, middleware.ConvertStructError("Primeiro nome \n"+requestFirstName), http.StatusBadRequest)
			return
		}

		requestSecondName := a.checkDatas("", copyBody.SecondName)
		if requestSecondName != "" {
			http.Error(w, middleware.ConvertStructError("Segundo nome \n"+requestSecondName), http.StatusBadRequest)
			return
		}

		requestEmail := a.checkDatas("email", copyBody.Email)
		if requestEmail != "" {
			http.Error(w, middleware.ConvertStructError("Email \n"+requestEmail), http.StatusBadRequest)
			return
		}

		requestDtCreateAccount := a.checkDatas("datetime", copyBody.CreateAccount)
		if requestDtCreateAccount != "" {
			http.Error(w, middleware.ConvertStructError("Data Criação \n"+requestDtCreateAccount), http.StatusBadRequest)
			return
		}

		requestDtExpiredAccount := a.checkDatas("datetime", copyBody.ExpiredAccount)
		if requestDtExpiredAccount != "" {
			http.Error(w, middleware.ConvertStructError("Data Expiração \n"+requestDtExpiredAccount), http.StatusBadRequest)
			return
		}

		requestQuantityAlert := a.checkDatas("", strconv.Itoa(copyBody.QuantityAlert))
		if requestQuantityAlert != "" {
			http.Error(w, middleware.ConvertStructError("Quantidade de alertas \n"+requestQuantityAlert), http.StatusBadRequest)
			return
		}

		requestAccountCopy := a.checkDatas("", strconv.Itoa(copyBody.AccountCopy))
		if requestAccountCopy != "" {
			http.Error(w, middleware.ConvertStructError("Quantidade de cópias \n"+requestAccountCopy), http.StatusBadRequest)
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
		loginValue := r.URL.Query().Get("login")
		datasLogin := a.checkDatas("email", loginValue)
		if datasLogin != "" {
			http.Error(w, middleware.ConvertStructError("E-mail inválido  \n"+datasLogin), http.StatusBadRequest)
			return
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
		datasLogin := a.checkDatas("email", loginValue)
		if datasLogin != "" {
			http.Error(w, "Dados da URL inválidos", http.StatusBadRequest)
			return
		}

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
		next.ServeHTTP(w, r)
	})
}

func (c *AgentController) checkDatas(typeValid string, data string) string {
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
