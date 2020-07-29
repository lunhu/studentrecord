package model

type PersEmploy struct {
	PersID			string	`db:"pers_id" cj:"pers_id"`
	EffDt			string  `db:"eff_dt" cj:"eff_dt"`
	EffSeq			string	`db:"eff_seq" cj:"eff_seq"`
	Employer		string  `db:"employer" cj:"employer"`
	EmplStatus		string  `db:"empl_status" cj:"empl_status"`
	HrStatus		string  `db:"hr_status" cj:"hr_status"`
	StartDt			string  `db:"start_dt" cj:"start_dt"`
	EndDt           string  `db:"end_dt" cj:"end_dt"`
	DepartID		string  `db:"depart_id" cj:"depart_id"`
	Position		string  `db:"position" cj:"position"`
	RegTemp			string  `db:"reg_temp" cj:"reg_temp"`
	FullPartTime	string  `db:"full_part_time" cj:"full_part_time"`
	JobIndicator    string  `db:"job_indicator" cj:"job_indicator"`
	JobNumber		string  `db:"job_number" cj:"job_number"`
	WorkLocation	string  `db:"work_location" cj:"work_location"`
	DepartOut		string  `db:"depart_out" cj:"depart_out"`
	StaffIndicator	string  `db:"staff_indicator" cj:"staff_indicator"`
	UID				string  `db:"u_id" cj:"u_id"`
}



func (employ *PersEmploy) StrictValid() bool {
	if employ.PersID == "" || employ.Employer =="" {
		return false
	}
	return true
}

func (employ *PersEmploy) StrictFaildMessage() string {
	var ErrMessage string
	if employ.PersID =="" {
		ErrMessage += "pers_id is null; "
	}

	if employ.Employer =="" {
		ErrMessage += "employer is null; "
	}

	return ErrMessage
}

func (employ *PersEmploy) Valid() bool {
	if employ.Employer =="" {
		return false
	}
	return true
}

func (employ *PersEmploy) FaildMessage() string {
	var ErrMessage string

	if employ.Employer =="" {
		ErrMessage += "employer is null; "
	}

	return ErrMessage
}