package dao

import (
	"github.com/jmoiron/sqlx"

	"studentrecord/model"
)

// InsertPersEmail 新增一条人员邮箱
func InsertPersEmail(tx *sqlx.Tx, params *model.PersEmail) (*model.PersEmail, error) {

	_, err := tx.Exec(
		`insert into e_emails_tbl 
		(pers_id, email_type, email, pref_email_flag, u_id) 
		values 
		(?,?,?,?,?)`,
		params.PersID,
		params.EmailType,
		params.Email,
		params.PrefEmailFlag,
		params.UID,
	)
	if err != nil {
		return nil, err
	}

	return params, nil
}