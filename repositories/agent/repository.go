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

	var receptorID = bodyChannel.UserReceptorID
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

	fmt.Print("actionType   " + actionType)
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

func (r *AgentRepository) SetNewPasswordAgent(body models.BodyPasswordAgent) (bool, error) {
	// update
	//
	_, err := r.db.Exec("UPDATE users_agent SET password_agent = ? WHERE id = ?;", body.PasswordAgent, body.ID)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return false, err
	}
	return true, nil
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

func (r *AgentRepository) checkExistChannel(channelCheck string, agentID int) (int, error) {
	var idChannel int
	err := r.db.QueryRow("SELECT id FROM channels WHERE channel_name = ? AND users_agent_id = ?", channelCheck, agentID).Scan(&idChannel)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idChannel, nil
}
func (r *AgentRepository) InsertChannel(bodyChannelReq models.QueryBodyCreateChannel) (int, error) {
	idChannel, err := r.checkExistChannel(bodyChannelReq.NameChannel, bodyChannelReq.AgentID)

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

//

func (r *AgentRepository) RemovePasswordAgent(emailAgent string) (models.RequestPasswordExist, error) {

	_, err := r.db.Exec("UPDATE users_agent SET password_agent = NULL WHERE email = ? AND password_agent IS NOT NULL", emailAgent)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return models.RequestPasswordExist{}, err
	}

	row := r.db.QueryRow("SELECT id, first_name, email FROM users_agent WHERE email = ?", emailAgent)
	var user models.RequestPasswordExist
	err = row.Scan(&user.ID, &user.FirstName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.RequestPasswordExist{}, nil // Não encontrou nenhum registro correspondente
		}
		fmt.Println("\n\n ERRO : ", err.Error())
		return models.RequestPasswordExist{}, err
	}

	return user, nil
}

func (r *AgentRepository) GetisValidLoginAdm(bodyLogin models.BodyPostLoginAdm) ([]models.QueryGetUsersAgent, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, password_agent, dt_create_account, dt_expired_account, account_valid, quantity_alerts, quantity_account_copy FROM users_agent WHERE email = ? ", bodyLogin.Login)
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
			&user.PasswordAgent,
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
	/*
		fmt.Println(&bodyLogin.Password)
		if users[0].PasswordAgent != &bodyLogin.Password {
			return nil, fmt.Errorf("Senha inválida")
		}
	*/
	return users, nil
}

//

func (r *AgentRepository) GetCredentialsReceptorChannels(email string) ([]models.RequestEmailsReceptor, error) {

	msgQuery := ""
	msgQuery += "SELECT UR.email AS login , CH.channel_name  AS channel_name "
	msgQuery += "FROM Permission AS PER "
	msgQuery += "JOIN channels AS CH ON PER.channel_id = CH.id "
	msgQuery += "JOIN users_agent AS UA ON CH.users_agent_id = UA.id "
	msgQuery += "JOIN users_receptor AS UR ON UR.agent_id = UA.id "
	msgQuery += "WHERE PER.user_receptor_id = (SELECT id FROM users_receptor WHERE email =? ) AND UR.email = ?;	"

	rows, err := r.db.Query(msgQuery, email, email)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.RequestEmailsReceptor
	for rows.Next() {
		var user models.RequestEmailsReceptor
		err := rows.Scan(
			&user.Login,
			&user.ChannelName,
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

func (r *AgentRepository) GetDataAgent(email string) ([]models.QueryGetUsersAgent, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, dt_create_account, dt_expired_account, account_valid, quantity_alerts, quantity_account_copy , password_agent FROM users_agent WHERE email = ?", email)
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
			&user.PasswordAgent,
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

func (r *AgentRepository) GetisValidLoginMt5(email string) ([]models.QueryGetUsersAgent, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, dt_create_account, dt_expired_account, account_valid, quantity_alerts, quantity_account_copy , password_agent FROM users_agent WHERE email = ?", email)
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
			&user.PasswordAgent,
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
	var passwordBody = bodyClientReq.Password_Agent
	var dtCreateBody = bodyClientReq.CreateAccount
	var dtExpiredBody = bodyClientReq.ExpiredAccount
	var accountValidBody = bodyClientReq.AccountValid
	var quantityAlertBody = bodyClientReq.QuantityAlert
	var quantityAccountCopyBody = bodyClientReq.AccountCopy
	request, err := r.db.Exec("INSERT INTO users_agent (first_name, 	second_name, 		email, 		password_agent,	dt_create_account, 	dt_expired_account, 	account_valid, 		quantity_alerts, 	quantity_account_copy) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		firtNameBody, secondNameBody, emailBody, passwordBody, dtCreateBody, dtExpiredBody, accountValidBody, quantityAlertBody, quantityAccountCopyBody)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir um agente no banco de dados  ", err.Error())
	}

	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}

func (r *AgentRepository) GetPermissionListOutReceptor(channelID string, agentID string) ([]models.RequestPermissionRequest, error) {
	//	rows, err := r.db.Query("SELECT id, first_name, second_name, email FROM users_receptor WHERE id IN (SELECT id FROM permission WHERE channel_id = ?);", channelID)
	rows, err := r.db.Query("SELECT  ur.id, ur.first_name, ur.second_name, ur.email FROM users_receptor AS ur WHERE ur.id NOT IN (SELECT user_receptor_id FROM Permission WHERE channel_id = ? )  AND agent_id = ?;", channelID, agentID)

	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}

	//	if !rows.Next() {
	//		fmt.Println("null")
	//		return nil, nil
	//	}
	defer rows.Close()

	var bodyChannelsList []models.RequestPermissionRequest

	for rows.Next() {
		var bodyCopyTrader models.RequestPermissionRequest
		err = rows.Scan(
			&bodyCopyTrader.ID,
			&bodyCopyTrader.FirstName,
			&bodyCopyTrader.SecondName,
			&bodyCopyTrader.Email)
		if err != nil {
			return nil, err
		}
		bodyChannelsList = append(bodyChannelsList, bodyCopyTrader)
	}

	return bodyChannelsList, nil
}

func (r *AgentRepository) GetPermissionListReceptor(channelID string) ([]models.RequestPermissionRequest, error) {
	msgQuery := ""
	msgQuery += " SELECT  u.id,  u.first_name,  u.second_name,  u.email,   p.channel_id, COALESCE((SELECT MAX(r.dt_send_copy) FROM req_copy r WHERE r.users_receptor_id = u.id), '0') AS dt_last_update"
	msgQuery += " FROM    users_receptor u"
	msgQuery += " LEFT JOIN    permission p ON u.id = p.user_receptor_id AND p.channel_id = ?"
	msgQuery += " WHERE   p.channel_id = ?;"
	rows, err := r.db.Query(msgQuery, channelID, channelID)

	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	var bodyChannelsList []models.RequestPermissionRequest

	for rows.Next() {
		var bodyCopyTrader models.RequestPermissionRequest
		err = rows.Scan(
			&bodyCopyTrader.ID,
			&bodyCopyTrader.FirstName,
			&bodyCopyTrader.SecondName,
			&bodyCopyTrader.Email,
			&bodyCopyTrader.ChannelID,
			&bodyCopyTrader.DateLastUpdate)
		if err != nil {
			return nil, err
		}
		bodyChannelsList = append(bodyChannelsList, bodyCopyTrader)
	}

	return bodyChannelsList, nil
}

func (r *AgentRepository) GetChannelList(structURL models.StrutcURLGetChannelList) ([]models.RequestChannelList, error) {

	rows, err := r.db.Query("SELECT id , users_agent_id , channel_name , dt_create_channel    FROM channels WHERE users_agent_id = ? AND dt_create_channel BETWEEN ? AND ?  LIMIT ?,?;",
		structURL.AgentID, structURL.DateStart, structURL.DateEnd, structURL.Offset, structURL.PageLimit)

	defer rows.Close()

	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	/*
		if !rows.Next() {
			fmt.Println("null")
			return nil, nil
		}
	*/
	var bodyChannelsList []models.RequestChannelList

	for rows.Next() {

		var bodyCopyTrader models.RequestChannelList
		err = rows.Scan(
			&bodyCopyTrader.ID,
			&bodyCopyTrader.AgentID,
			&bodyCopyTrader.ChannelName,
			&bodyCopyTrader.DateCreate,
		)
		if err != nil {
			return nil, err
		}
		bodyChannelsList = append(bodyChannelsList, bodyCopyTrader)
	}
	return bodyChannelsList, nil
}
func (r *AgentRepository) GetPermissionChannelList(structURL models.StrutcURLGetChannelList) ([]models.RequestChannelList, error) {

	msgQuery := ""
	msgQuery += " SELECT  c.id,  c.users_agent_id,  c.channel_name,  c.dt_create_channel,"
	msgQuery += " (SELECT COUNT(*) FROM Permission p WHERE p.channel_id = c.id) AS total_receptor_copy"
	msgQuery += " FROM  channels c"
	msgQuery += " WHERE  c.users_agent_id = ?"
	msgQuery += " AND c.dt_create_channel BETWEEN ? AND ?"
	msgQuery += " LIMIT ?,?"
	//	msgQuery := "SELECT id , users_agent_id , channel_name , dt_create_channel    FROM channels WHERE users_agent_id = ? AND dt_create_channel BETWEEN ? AND ?  LIMIT ?,?;"
	rows, err := r.db.Query(msgQuery,
		structURL.AgentID, structURL.DateStart, structURL.DateEnd, structURL.Offset, structURL.PageLimit)

	defer rows.Close()

	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}

	var bodyChannelsList []models.RequestChannelList

	for rows.Next() {

		var bodyCopyTrader models.RequestChannelList
		err = rows.Scan(
			&bodyCopyTrader.ID,
			&bodyCopyTrader.AgentID,
			&bodyCopyTrader.ChannelName,
			&bodyCopyTrader.DateCreate,
			&bodyCopyTrader.TotalReceptorCopy,
		)
		if err != nil {
			return nil, err
		}
		bodyChannelsList = append(bodyChannelsList, bodyCopyTrader)
	}
	return bodyChannelsList, nil
}

func (r *AgentRepository) GetInformationChannel(channelID int) (models.RequestInformationChannel, error) {
	// Preparar a declaração SQL
	query := `
        SELECT
            c.channel_name,
            c.dt_create_channel,
            (SELECT COUNT(*) FROM all_copy ac WHERE ac.channel_id = c.id) AS count_channel
        FROM
            channels c
        WHERE
            c.id = ?;
    `

	// Executar a consulta SQL e recuperar os resultados
	var result models.RequestInformationChannel
	err := r.db.QueryRow(query, channelID).Scan(
		&result.NameChannel,
		&result.DateCreateChannel,
		&result.CountCopy,
	)
	if err != nil {
		return models.RequestInformationChannel{}, err
	}

	return result, nil
}

func (r *AgentRepository) DeleteChannel(structURL models.BodyDelete) bool {

	_, err := r.db.Exec("DELETE FROM req_copy WHERE channel_id = ?;", structURL.ID)
	if err != nil {
		// Lida com o erro
	}

	_, err = r.db.Exec("DELETE FROM all_copy WHERE channel_id = ?;", structURL.ID)
	if err != nil {
		// Lida com o erro
	}

	_, err = r.db.Exec("DELETE FROM permission WHERE channel_id = ?;", structURL.ID)
	if err != nil {
		// Lida com o erro
	}

	result, err := r.db.Exec("DELETE FROM channels WHERE id = ? AND users_agent_id = ?;", structURL.ID, structURL.AgentID)
	if err != nil {
		// Lida com o erro
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		// Nenhum registro foi deletado
		return false
	}

	return true
}

func (r *AgentRepository) EditChannel(structUpdate models.BodyUpdate) (int, error) {
	updateSQL := "UPDATE channels SET channel_name=? WHERE id=? AND users_agent_id=?"

	// Executar a instrução SQL
	_, err := r.db.Exec(updateSQL, structUpdate.NewNameChannel, structUpdate.ID, structUpdate.AgentID)
	if err != nil {
		return 0, err
	}
	return structUpdate.ID, nil
}
