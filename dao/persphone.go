package dao

import (
	"github.com/jmoiron/sqlx"

	"studentrecord/model"
)

// InsertPersPhone 新增一条人员电话
func InsertPersPhone(tx *sqlx.Tx, params *model.PersPhone) (*model.PersPhone, error) {

	_, err := tx.Exec(
		`insert into e_phones_tbl 
		(pers_id, phone_type, country_code, phone, extension, pref_phone_flag, u_id) 
		values 
		(?,?,?,?,?,?,?)`,
		params.PersID,
		params.PhoneType,
		params.CountryCode,
		params.Phone,
		params.Extension,
		params.PrefPhoneFlag,
		params.UID,
	)
	if err != nil {
		return nil, err
	}

	return params, nil
}