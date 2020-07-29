package service

import (
	"strconv"
	"studentrecord/common"
	"studentrecord/dao"
	"studentrecord/model"

	"github.com/tidwall/gjson"
)

// PersEmployAdd 新增一条人员工作经历
func PersEmployAdd(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {

	if reqJSON.Get(`template.data.#(name=="employer")`).Exists() {
		employs := reqJSON.Get(`template.data`).Array()
		persEmploy := &model.PersEmploy{}
		FillPersEmploy(employs, persEmploy)

		if persEmploy.EffSeq == "" {

			eff_seq, err := dao.SelectMaxSeq(sqlxDB, persEmploy.PersID)
	
			if err != nil {
				return common.HandError("500", "")
			} 
		
			persEmploy.EffSeq = strconv.Itoa(int(eff_seq + 1))
		}

		if persEmploy.StrictValid() {
			tx, err := sqlxDB.Beginx()
			if err != nil {
				return common.HandError("500", "")
			}
			persEmploy, err := dao.InsertPersEmploy(tx, persEmploy)
			if err != nil {
				tx.Rollback()
				return common.HandError("500", "")
			}
			tx.Commit()
			return common.HandItem(*persEmploy, 201)
		} else {
			return common.HandError("400", "")
		}
	}
	
	return common.HandError("400", "")
}


func FillPersEmploy(result []gjson.Result, persEmploy *model.PersEmploy) {
	for _, v :=	range result {

		switch v.Get("name").String() {
		case "pers_id":
			persEmploy.PersID = v.Get("value").String()
	    case "eff_dt":
			persEmploy.EffDt = v.Get("value").String()
		case "eff_seq":
			persEmploy.EffSeq = v.Get("value").String()
	    case "employer":
			persEmploy.Employer = v.Get("value").String()
		case "empl_status":
			persEmploy.EmplStatus = v.Get("value").String()
	    case "hr_status":
			persEmploy.HrStatus = v.Get("value").String()
		case "start_dt":
			persEmploy.StartDt = v.Get("value").String()
		case "end_dt":
			persEmploy.EndDt = v.Get("value").String()
		case "depart_id":
			persEmploy.DepartID = v.Get("value").String()
		case "position":
			persEmploy.Position = v.Get("value").String()
		case "reg_temp":
			persEmploy.RegTemp = v.Get("value").String()
		case "full_part_time":
			persEmploy.FullPartTime = v.Get("value").String()
		case "job_indicator":
			persEmploy.JobIndicator = v.Get("value").String()
		case "job_number":
			persEmploy.JobNumber = v.Get("value").String()
		case "work_location":
			persEmploy.WorkLocation = v.Get("value").String()
		case "depart_out":
			persEmploy.DepartOut = v.Get("value").String()
		case "staff_indicator":
			persEmploy.StaffIndicator = v.Get("value").String()
	   case "u_id":
			persEmploy.UID = v.Get("value").String()
		}
	}
}