package repository_agent

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/RenanFerreira0023/FiberTemp/models"
)

type AgentRepository struct {
	db *sql.DB
}

func NewAgentRepository(db *sql.DB) *AgentRepository {
	return &AgentRepository{db: db}
}

func (r *AgentRepository) InsertPermissionChannel(bodyChannel models.QueryBodyInsertPermission) (int, error) {

	idPermission, errPermission := r.checkExistPermission(bodyChannel)
	if errPermission == nil {
		return idPermission, fmt.Errorf("Canal ja existe ", idPermission)
	}

	var receptorID = bodyChannel.ChannelID
	var channelID = bodyChannel.ChannelID
	request, err := r.db.Exec("INSERT INTO permission (user_receptor_id, channel_id) VALUES (?, ?)",
		receptorID, channelID)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir um canal no banco de dados  ", err.Error())
	}
	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}
func (r *AgentRepository) checkExistPermission(bodyChannel models.QueryBodyInsertPermission) (int, error) {

	var idUser int
	err := r.db.QueryRow("SELECT id FROM permission WHERE user_receptor_id = ? AND channel_id = ? ", bodyChannel.UserReceptorID, bodyChannel.ChannelID).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}

func (r *AgentRepository) checkExistCopy(ticketCheck int, entryCheck string) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM all_copy WHERE ticket = ? AND dt_send_order = ?", ticketCheck, entryCheck).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}

func (r *AgentRepository) SendCopy(bodyCopy models.QueryBodySendCopy) (int, error) {
	idUserAgent, err := r.checkExistCopy(int(bodyCopy.Ticket), bodyCopy.DateEntry)

	if err == nil || idUserAgent != 0 {
		return idUserAgent, fmt.Errorf("Copy ja existe")
	}
	var symbol = bodyCopy.Symbol
	var actionType = bodyCopy.ActionType
	var ticket = bodyCopy.Ticket
	var lot = bodyCopy.Lot
	var targetPedding = bodyCopy.TargetPedding
	var takeprofit = bodyCopy.TakeProfit
	var stoploss = bodyCopy.StopLoss
	var dateEntry = bodyCopy.DateEntry
	var agentId = bodyCopy.UserAgentID
	var channelId = bodyCopy.ChannelID

	request, err := r.db.Exec("INSERT INTO all_copy (symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		symbol, actionType, ticket, lot, targetPedding, takeprofit, stoploss, dateEntry, agentId, channelId)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir um canal no banco de dados  ", err.Error())
	}
	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}

func (r *AgentRepository) GetAgentFromEmailAndChannel(email string, channel string) (models.QueryGetLogin, error) {
	var idAgent int
	var idChannel int
	var qttAlert int
	var qttCopyAccounts int
	var structLogin models.QueryGetLogin
	err := r.db.QueryRow("SELECT a.id  , c.id  , a.quantity_alerts, a.quantity_account_copy FROM users_agent AS a, channels AS c WHERE c.users_agent_id = a.id and a.email = ? and c.channel_name = ?", email, channel).Scan(&idAgent, &idChannel, &qttAlert, &qttCopyAccounts)
	if err != nil {
		structLogin.AgentID = -1
		structLogin.ChannelID = -1
		structLogin.QuantityAlert = -1
		structLogin.AccountCopy = -1

		return structLogin, fmt.Errorf("Agente ou canal não encontrado")
	}

	structLogin.AgentID = idAgent
	structLogin.ChannelID = idChannel
	structLogin.QuantityAlert = qttAlert
	structLogin.AccountCopy = qttCopyAccounts

	return structLogin, nil
}

func (r *AgentRepository) checkExistChannel(channelCheck string) (int, error) {
	var idChannel int
	err := r.db.QueryRow("SELECT id FROM channels WHERE channel_name = ? ", channelCheck).Scan(&idChannel)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idChannel, nil
}
func (r *AgentRepository) InsertChannel(bodyChannelReq models.QueryBodyCreateChannel) (int, error) {
	idChannel, err := r.checkExistChannel(bodyChannelReq.NameChannel)

	if err == nil || idChannel != 0 {
		return idChannel, fmt.Errorf("Canal ja existe	", strconv.Itoa(idChannel))
	}

	var idAgent = bodyChannelReq.AgentID
	var nameChannel = bodyChannelReq.NameChannel
	var dateCreateChannel = bodyChannelReq.CreateChannel

	request, err := r.db.Exec("INSERT INTO channels (users_agent_id, 	channel_name, 		dt_create_channel) VALUES (?, ?, ?)",
		idAgent, nameChannel, dateCreateChannel)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir um canal no banco de dados  ", err.Error())
	}

	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}

func (r *AgentRepository) GetisValidLogin(email string) ([]models.QueryGetUsersAgent, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, dt_create_account, dt_expired_account, account_valid, quantity_alerts, quantity_account_copy FROM users_agent WHERE email = ?", email)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.QueryGetUsersAgent
	for rows.Next() {
		var user models.QueryGetUsersAgent
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
			&user.CreateAccount,
			&user.ExpiredAccount,
			&user.AccountValid,
			&user.QuantityAlert,
			&user.AccountCopy,
		)
		if err != nil {
			fmt.Println("\n\n ERRO : ", err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("Usuário não encontrado")
	}
	return users, nil
}

func (r *AgentRepository) checkExistLogin(emailCheck string) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM users_agent WHERE email = ? ", emailCheck).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}

func (r *AgentRepository) InsertClient(bodyClientReq models.QueryBodyUsersAgent) (int, error) {
	idUserAgent, err := r.checkExistLogin(bodyClientReq.Email)

	if err == nil || idUserAgent != 0 {
		return 0, fmt.Errorf("Cadastro ja existe")
	}

	var firtNameBody = bodyClientReq.FirstName
	var secondNameBody = bodyClientReq.SecondName
	var emailBody = bodyClientReq.Email
	var dtCreateBody = bodyClientReq.CreateAccount
	var dtExpiredBody = bodyClientReq.ExpiredAccount
	var accountValidBody = bodyClientReq.AccountValid
	var quantityAlertBody = bodyClientReq.QuantityAlert
	var quantityAccountCopyBody = bodyClientReq.AccountCopy
	request, err := r.db.Exec("INSERT INTO users_agent (first_name, 	second_name, 		email, 		dt_create_account, 	dt_expired_account, 	account_valid, 		quantity_alerts, 	quantity_account_copy) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		firtNameBody, secondNameBody, emailBody, dtCreateBody, dtExpiredBody, accountValidBody, quantityAlertBody, quantityAccountCopyBody)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir uma copy no banco de dados  ", err.Error())
	}

	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}
