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
	errCheck := r.db.QueryRow("SELECT id FROM req_copy WHERE channel_id = ? AND all_copy_id = ? ;", channelID, allCopyID).Scan(&idUser)
	if errCheck == nil {
		return 0, fmt.Errorf("Essa ação ja foi executada")
	}
	_, err := r.db.Exec("INSERT INTO req_copy (dt_send_copy, channel_id, users_receptor_id, all_copy_id)  VALUES (?,?,?,?)",
		dateSendOrder, channelID, receptorID, allCopyID)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir uma requisição de copy no banco de dados  ", err.Error())
	}

	return 0, nil
}

func (r *ReceptorRepository) CheckReqCopy(idChannel int, idReceptor int, idAllCopy int, dateSendOrder string) error {
	fmt.Println("\n\n idChannel    ", idChannel)
	fmt.Println(" idReceptor    ", idReceptor)
	fmt.Println(" idAllCopy    ", idAllCopy)
	var idUser int
	err := r.db.QueryRow("SELECT id FROM req_copy WHERE channel_id = ? AND users_receptor_id = ? AND all_copy_id = ? ;", idChannel, idReceptor, idAllCopy).Scan(&idUser)
	if err == nil {
		return fmt.Errorf("Ja respondeu essa requisicao")
	}
	return nil
}

func (r *ReceptorRepository) GetCopyTrader(structURL models.StrutcURLCopyTrader) ([]models.BodyCopyTrader, error) {

	rows, err := r.db.Query("SELECT id, symbol, action_type, ticket, lot, target_pedding, takeprofit, stoploss, dt_send_order, user_agent_id, channel_id   FROM all_copy WHERE dt_send_order BETWEEN ? AND ? AND user_agent_id = ? AND channel_id = ? LIMIT ?,?;",
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
		return fmt.Errorf("Você não tem permissão para acessar esse canal")
	}
	return nil
}

func (r *ReceptorRepository) GetAgentID(emailAgent string) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM users_agent WHERE email = ? ", emailAgent).Scan(&idUser)
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
	rows, err := r.db.Query("SELECT id,  dt_expired_account FROM users_receptor WHERE email = ?", email)
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

	_, err = r.db.Exec("INSERT INTO users_receptor (first_name, second_name, email, dt_create_account, dt_expired_account)  VALUES (?,?,?,?,?)",
		firtNameBody, secondNameBody, emailBody, dtCreateBody, dtExpiredBody)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir uma copy no banco de dados  ", err.Error())
	}

	return 0, nil
}

func (r *ReceptorRepository) checkExistLogin(emailCheck string) (int, error) {
	var idUser int
	err := r.db.QueryRow("SELECT id FROM users_receptor WHERE email = ? ", emailCheck).Scan(&idUser)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return idUser, nil
}
