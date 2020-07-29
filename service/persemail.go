package service

import (
	"studentrecord/common"
	"studentrecord/dao"
	"studentrecord/model"

	"github.com/tidwall/gjson"
)
// PersEmailAdd 新增一条人员邮箱
func PersEmailAdd(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {

	if reqJSON.Get(`template.data.#(name=="email")`).Exists() {
		emails := reqJSON.Get(`template.data`).Array()
		persEmail := &model.PersEmail{}
		FillPersEmail(emails, persEmail)

		if persEmail.StrictValid() {
			tx, err := sqlxDB.Beginx()
			if err != nil {
				return common.HandError("500", "")
			}
			persEmail, err := dao.InsertPersEmail(tx, persEmail)
			if err != nil {
				tx.Rollback()
				return common.HandError("500", "")
			}
			tx.Commit()
			return common.HandItem(*persEmail, 201)
		} else {
			return common.HandError("400", "")
		}
	}

	return common.HandError("400", "")
}



func FillPersEmail(result []gjson.Result, persEmail *model.PersEmail) {
	for _, v :=	range result {

		switch v.Get("name").String() {
		case "pers_id":
			persEmail.PersID = v.Get("value").String()
	   case "email_type":
			persEmail.EmailType = v.Get("value").String()
	   case "email":
			persEmail.Email = v.Get("value").String()
	   case "pref_email_flag":
			persEmail.PrefEmailFlag = v.Get("value").String()
	   case "u_id":
			persEmail.UID = v.Get("value").String()
		}
	}
}