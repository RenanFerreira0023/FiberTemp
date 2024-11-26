package repository_receptor

import (
	"database/sql"
	"fmt"

	"github.com/RenanFerreira0023/FiberTemp/models"
)

type ReceptorRepository struct {
	db *sql.DB
}

func NewReceptorRepository(db *sql.DB) *ReceptorRepository {
	return &ReceptorRepository{db: db}
}

func (r *ReceptorRepository) InsertReqCopy(bodyReqCopy models.QueryRequestReqCopy) (int, error) {

	var channelID = bodyReqCopy.ChannelID
	var receptorID = bodyReqCopy.ReceptorID
	var allCopyID = bodyReqCopy.AllCopyID
	var dateSendOrder = bodyReqCopy.DateSendOrder

	var idUser int
	errCheck := r.db.QueryRow("SELECT id FROM req_copy WHERE channel_id = ? AND users_receptor_id = ? all_copy_id = ? ;", channelID, receptorID, allCopyID).Scan(&idUser)
	if errCheck == nil {
		return 0, fmt.Errorf("Essa ação ja foi executada")
	}
	request, err := r.db.Exec("INSERT INTO req_copy (dt_send_copy, channel_id, users_receptor_id, all_copy_id)  VALUES (?,?,?,?)",
		dateSendOrder, channelID, receptorID, allCopyID)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir uma requisição de copy no banco de dados  ", err.Error())
	}

	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}

/*

func (r *ReceptorRepository) CheckReqCopy(idChannel int, idReceptor int, idAllCopy int, dateSendOrder string) string {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM req_copy WHERE channel_id = ? AND users_receptor_id = ? AND all_copy_id = ? ;", idChannel, idReceptor, idAllCopy).Scan(&idUser)
	if err == nil {
		return "NAO_ACIONADO" //fmt.Errorf("Ja respondeu essa requisicao")
	}
	return "JA_ACIONADO"
}
*/

func (r *ReceptorRepository) CheckReqCopy(idChannel int, idReceptor int, idAllCopy int) string {
	var idUser int
	r.db.QueryRow("SELECT id FROM req_copy WHERE channel_id = ? AND users_receptor_id = ? AND all_copy_id = ? ;", idChannel, idReceptor, idAllCopy).Scan(&idUser)
	if idUser == 0 {
		return "NAO_ACIONADO" //fmt.Errorf("Ja respondeu essa requisicao")
	}
	return "JA_ACIONADO"
}

func (r *ReceptorRepository) GetCopyTrader(structURL models.StrutcURLCopyTrader) ([]models.BodyCopyTrader, error) {
	var allBodyCopyTraders []models.BodyCopyTrader
	structURL.Offset = 0

	for {
		//		fmt.Printf("\n\nSELECT id, symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id FROM all_copy WHERE dt_send_order BETWEEN '%s' AND '%s' AND user_agent_id = %d AND channel_id = %d LIMIT %d, %d;",
		//			structURL.DateStart, structURL.DateEnd, structURL.AgentID, structURL.ChannelID, structURL.Offset, structURL.PageLimit)

		rows, err := r.db.Query("SELECT id, symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id FROM all_copy WHERE dt_send_order BETWEEN ? AND ? AND user_agent_id = ? AND channel_id = ? LIMIT ?,?;",
			structURL.DateStart, structURL.DateEnd, structURL.AgentID, structURL.ChannelID, structURL.Offset, structURL.PageLimit)

		if err != nil {
			fmt.Println("\n\n ERRO : ", err.Error())
			return nil, err
		}

		defer rows.Close()

		var bodyCopyTraders []models.BodyCopyTrader

		for rows.Next() {
			var bodyCopyTrader models.BodyCopyTrader
			err = rows.Scan(
				&bodyCopyTrader.ID,
				&bodyCopyTrader.Symbol,
				&bodyCopyTrader.ActionType,
				&bodyCopyTrader.Ticket,
				&bodyCopyTrader.Lot,
				&bodyCopyTrader.TargetPedding,
				&bodyCopyTrader.TakeProfit,
				&bodyCopyTrader.StopLoss,
				&bodyCopyTrader.DateEntry,
				&bodyCopyTrader.AgentID,
				&bodyCopyTrader.ChannelID,
			)
			if err != nil {
				return nil, err
			}
			bodyCopyTraders = append(bodyCopyTraders, bodyCopyTrader)
		}

		if len(bodyCopyTraders) == 0 {
			break
		}

		allBodyCopyTraders = append(allBodyCopyTraders, bodyCopyTraders...)
		structURL.Offset += structURL.PageLimit
	}

	return allBodyCopyTraders, nil
}

func (r *ReceptorRepository) GetCopyTrader2(structURL models.StrutcURLCopyTrader) ([]models.BodyCopyTrader, error) {

	fmt.Printf("\n\nSELECT id, symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id FROM all_copy WHERE dt_send_order BETWEEN '%s' AND '%s' AND user_agent_id = %d AND channel_id = %d LIMIT %d, %d;",
		structURL.DateStart, structURL.DateEnd, structURL.AgentID, structURL.ChannelID, structURL.Offset, structURL.PageLimit)

	rows, err := r.db.Query("SELECT id, symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id FROM all_copy WHERE dt_send_order BETWEEN ? AND ? AND user_agent_id = ? AND channel_id = ? LIMIT ?,?;",
		structURL.DateStart, structURL.DateEnd, structURL.AgentID, structURL.ChannelID, structURL.Offset, structURL.PageLimit)

	defer rows.Close()

	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}

	//	if !rows.Next() {
	//		return nil, nil
	//	}
	var bodyCopyTraders []models.BodyCopyTrader

	for rows.Next() {
		var bodyCopyTrader models.BodyCopyTrader
		err = rows.Scan(
			&bodyCopyTrader.ID,
			&bodyCopyTrader.Symbol,
			&bodyCopyTrader.ActionType,
			&bodyCopyTrader.Ticket,
			&bodyCopyTrader.Lot,
			&bodyCopyTrader.TargetPedding,
			&bodyCopyTrader.TakeProfit,
			&bodyCopyTrader.StopLoss,
			&bodyCopyTrader.DateEntry,
			&bodyCopyTrader.AgentID,
			&bodyCopyTrader.ChannelID,
		)
		if err != nil {
			return nil, err
		}
		bodyCopyTraders = append(bodyCopyTraders, bodyCopyTrader)
	}

	return bodyCopyTraders, nil
}

func (r *ReceptorRepository) GetPermissionChannel(idChannel int, idReceptor int) error {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM permission WHERE user_receptor_id = ? AND channel_id = ? ", idReceptor, idChannel).Scan(&idUser)
	if err != nil {
		return fmt.Errorf("Você não tem permissão para acessar esse canal  ", err.Error())
	}
	return nil
}

func (r *ReceptorRepository) GetAgentID(agentID int) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM users_agent WHERE id = ? ", agentID).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}

func (r *ReceptorRepository) GetChannel(channelName string, idAgent int) (models.QueryGetChannel, error) {
	var chn models.QueryGetChannel
	//	var idChannel int
	err := r.db.QueryRow("SELECT id , users_agent_id FROM channels WHERE channel_name = ? AND users_agent_id = ? ", channelName, idAgent).Scan(&chn.ID, &chn.AgentID)
	if err != nil {
		return chn, fmt.Errorf("Canal ou Agente não encontrado.")
	}
	return chn, nil
}

func (r *ReceptorRepository) GetValidReceptor(email string) ([]models.QueryGetValidReceptor, error) {
	rows, err := r.db.Query("SELECT id,  dt_expired_account , agent_id FROM users_receptor WHERE email = ?", email)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []models.QueryGetValidReceptor
	for rows.Next() {
		var user models.QueryGetValidReceptor
		err := rows.Scan(
			&user.ID,
			&user.ExpiredAccount,
			&user.AgentID,
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
		return nil, fmt.Errorf("Receptor não encontrado")
	}
	return users, nil
}

func (r *ReceptorRepository) GetisValidLogin(email string) ([]models.QueryBodyUserReceptor, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, dt_create_account, dt_expired_account FROM users_receptor WHERE email = ?", email)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.QueryBodyUserReceptor
	for rows.Next() {
		var user models.QueryBodyUserReceptor
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
			&user.CreateAccount,
			&user.ExpiredAccount,
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

func (r *ReceptorRepository) GetListReceptor(agentID int) ([]models.QueryBodyUserReceptor, error) {
	rows, err := r.db.Query("SELECT id, first_name, second_name, email, dt_create_account, dt_expired_account,agent_id FROM users_receptor WHERE agent_id = ?;", agentID)
	if err != nil {
		fmt.Println("\n\n ERRO : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.QueryBodyUserReceptor
	for rows.Next() {
		var user models.QueryBodyUserReceptor
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
			&user.CreateAccount,
			&user.ExpiredAccount,
			&user.AgentID,
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

func (r *ReceptorRepository) DeleteReceptor(id_receptor int, id_agent int) bool {

	result, err := r.db.Exec("DELETE FROM users_receptor WHERE id = ? AND agent_id = ?;", id_receptor, id_agent)
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

func (r *ReceptorRepository) DeleteChannelPermissionReceptor(id_receptor int, channelID int) bool {

	result, err := r.db.Exec("DELETE FROM permission  WHERE user_receptor_id = ? AND channel_id = ?;", id_receptor, channelID)
	if err != nil {
		// Lida com o erro
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		// Nenhum registro foi deletado
		return true
	}

	return true
}

func (r *ReceptorRepository) EditReceptor(structReceptor models.BodyEditReceptor) (int, bool) {
	result, err := r.db.Exec("UPDATE users_receptor SET first_name = ?, second_name = ?, email = ? WHERE id = ? AND agent_id = ?;",
		structReceptor.FirstName, structReceptor.SecondName, structReceptor.Email, structReceptor.ReceptorID, structReceptor.AgentID)
	if err != nil {
		// Lida com o erro
		return 0, false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		// Nenhum registro foi alterado
		return 0, false
	}

	return structReceptor.ReceptorID, true
}

func (r *ReceptorRepository) InsertClient(bodyClientReq models.QueryGetUserReceptor) (int, error) {
	idUserAgent, err := r.checkExistLogin(bodyClientReq.Email)

	if err == nil || idUserAgent != 0 {
		return idUserAgent, fmt.Errorf("Cadastro ja existe")
	}

	var firtNameBody = bodyClientReq.FirstName
	var secondNameBody = bodyClientReq.SecondName
	var emailBody = bodyClientReq.Email
	var dtCreateBody = bodyClientReq.CreateAccount
	var dtExpiredBody = bodyClientReq.ExpiredAccount
	var agentID = bodyClientReq.AgentID
	request, err := r.db.Exec("INSERT INTO users_receptor (first_name, second_name, email, dt_create_account, dt_expired_account , agent_id)  VALUES (?,?,?,?,?,?)",
		firtNameBody, secondNameBody, emailBody, dtCreateBody, dtExpiredBody, agentID)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir uma copy no banco de dados  ", err.Error())
	}

	insertID, err := request.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(insertID), nil
}

func (r *ReceptorRepository) checkExistLogin(emailCheck string) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM users_receptor WHERE email = ? ", emailCheck).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}
