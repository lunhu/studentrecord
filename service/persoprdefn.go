package service 

import (
	"studentrecord/common"
	"studentrecord/dao"

	"github.com/tidwall/gjson"
)

// SelectPersOprDefnByPersID 根据 pers_id 访问dao.
func SelectPersOprDefnByPersID(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {

	if reqJSON.Get(`pers_id`).Exists() {

		pers_id := reqJSON.Get(`pers_id`).String()

		if pers_id == "" {
			return common.HandError("400", "")
		}
	
		persOprDefn, err := dao.SelectPersOprDefnByPersID(sqlxDB, pers_id)
		if err != nil {
			return common.HandError("500", "")
		}
	
		return common.HandItem(*persOprDefn, 200)
	}
	return common.HandError("400", "")
}

// SelectPersOprDefnByOprID 根据 opr_id 访问dao.
func SelectPersOprDefnByOprID(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {
	if reqJSON.Get(`opr_id`).Exists() {

		opr_id := reqJSON.Get(`opr_id`).String()

		if opr_id == "" {
			return common.HandError("400", "")
		}

		persOprDefn, err := dao.SelectPersOprDefnByOprID(sqlxDB, opr_id)
		if err != nil {
			return common.HandError("500", "")
		}

		return common.HandItem(*persOprDefn, 200)
	}
	return common.HandError("400", "")
}