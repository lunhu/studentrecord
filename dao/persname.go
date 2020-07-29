package dao

import (

	"github.com/jmoiron/sqlx"

	"studentrecord/model"
)

// 查找当前最大的pers_id

func SelectMaxPersId(sqlxDB *sqlx.DB) (int64, error) {
	var maxID int64
	err := sqlxDB.Get(&maxID, "SELECT MAX(pers_id) FROM e_names_tbl")
	if err != nil {
		return 0, err
	}

	return maxID, nil
}

//InsertPersName 新增一条人员姓名
func InsertPersName(tx *sqlx.Tx, params *model.PersName) (*model.PersName, error) {

	_, err := tx.Exec(
		`insert into e_names_tbl 
		(pers_id, name_type, eff_dt, eff_status, name, last_name, middle_name, first_name, pref_name_flag, u_id) 
		values 
		(?,?,?,?,?,?,?,?,?,?)`,
		params.PersID,
		params.NameType,
		params.EffDt,
		params.EffStatus,
		params.Name,
		params.LastName,
		params.MiddleName,
		params.FirstName,
		params.PrefNameFlag,
		params.UID,
	)

	if err != nil {
		return nil, err
	}

	return params, nil
}