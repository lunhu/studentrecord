package dao

import (

	"github.com/jmoiron/sqlx"

	"studentrecord/model"

)

// SelectPersOprDefnByPersID 根据 pers_id 获取单个 PersOprDefn 对象
func SelectPersOprDefnByPersID(sqlxDB *sqlx.DB, pers_id string) (*model.PersOprDefn, error) {
	var persOprDefn model.PersOprDefn
	err := sqlxDB.Get(&persOprDefn, "SELECT * FROM e_oprdefn_tbl a WHERE a.pers_id = ? ", pers_id)
	if err != nil {
		return nil, err
	}

	return &persOprDefn, nil
}

// SelectPersOprDefnByOprID 根据 opr_id 获取单个 PersOprDefn 对象
func SelectPersOprDefnByOprID(sqlxDB *sqlx.DB, opr_id string) (*model.PersOprDefn, error) {
	var persOprDefn model.PersOprDefn
	err := sqlxDB.Get(&persOprDefn, "SELECT * FROM e_oprdefn_tbl a WHERE a.opr_id = ? ", opr_id)
	if err != nil {
		return nil, err
	}

	return &persOprDefn, nil
}