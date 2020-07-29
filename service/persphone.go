package service

import (
	"studentrecord/common"
	"studentrecord/dao"
	"studentrecord/model"

	"github.com/tidwall/gjson"
)
// PersPhoneAdd 新增一条人员电话
func PersPhoneAdd(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {
	if reqJSON.Get(`template.data.#(name=="phone")`).Exists() {
		
		phones := reqJSON.Get(`template.data`).Array()
		persPhone := &model.PersPhone{}
		FillPersPhone(phones, persPhone)

		if persPhone.StrictValid() {
			tx, err := sqlxDB.Beginx()
			if err != nil {
				return common.HandError("500", "")
			}
			persPhone, err := dao.InsertPersPhone(tx, persPhone)
			if err != nil {
				tx.Rollback()
				return common.HandError("500", "")
			}
			tx.Commit()
			return common.HandItem(*persPhone, 201)
		} else {
			return common.HandError("400", "")
		}
	}

	return common.HandError("400", "")
}

func FillPersPhone(result []gjson.Result, persPhone *model.PersPhone) {
	for _, v :=	range result {
		switch v.Get("name").String() {
		case "pers_id":
			persPhone.PersID = v.Get("value").String()
	   case "phone_type":
			persPhone.PhoneType = v.Get("value").String()
		case "country_code":
			persPhone.CountryCode = v.Get("value").String()
	   case "phone":
			persPhone.Phone = v.Get("value").String()
		case "extension":
			persPhone.Extension = v.Get("value").String()
	   case "pref_phone_flag":
			persPhone.PrefPhoneFlag = v.Get("value").String()
	   case "u_id":
			persPhone.UID = v.Get("value").String()
		}
	}
}