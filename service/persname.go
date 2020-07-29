package service

import (
	"strconv"
	"studentrecord/common"
	"studentrecord/dao"
	"studentrecord/model"

	"github.com/tidwall/gjson"
)
// PersNameAdd 新增一个人员姓名
func PersNameAdd(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {

	if reqJSON.Get(`template.data.#(name=="name")`).Exists() {
		
		names := reqJSON.Get(`template.data`).Array()

		persName := &model.PersName{}

		FillPersName(names, persName)

		if persName.PersID == "" {

			pers_id, err := dao.SelectMaxPersId(sqlxDB)
	
			if err != nil {
				return common.HandError("500", "")
			} 
		
			persName.PersID = strconv.Itoa(int(pers_id + 1))
		}


		if persName.StrictValid() {
			tx, err := sqlxDB.Beginx()
			if err != nil {
				return common.HandError("500", "")
			}
			persName, err := dao.InsertPersName(tx, persName)
			if err != nil {
				tx.Rollback()
				return common.HandError("500", "")
			}
			tx.Commit()
			return common.HandItem(*persName, 201)
		} else {
			return common.HandError("400", "")
		}
		
	}

	return common.HandError("400", "")

}

func FillPersName(result []gjson.Result, persName *model.PersName) {
	for _, v :=	range result {

		switch v.Get("name").String() {
		case "pers_id":
			persName.PersID = v.Get("value").String()
		case "name_type":
			persName.NameType = v.Get("value").String()
		case "eff_dt":
			persName.EffDt = v.Get("value").String()
		case "eff_status":
			persName.EffStatus = v.Get("value").String()
		case "name":
			persName.Name = v.Get("value").String()
		case "last_name":
			persName.LastName = v.Get("value").String()
		case "middle_name":
			persName.MiddleName = v.Get("value").String()
		case "first_name":
			persName.FirstName = v.Get("value").String()
		case "pref_name_flag":
			persName.PrefNameFlag = v.Get("value").String()
		case "u_id":
			persName.UID = v.Get("value").String()
		}
	}
}