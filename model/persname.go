package model 

type PersName struct {
	PersID			string	`db:"pers_id" cj:"pers_id"`
	NameType		string  `db:"name_type" cj:"name_type"`
	EffDt			string	`db:"eff_dt" cj:"eff_dt"`
	EffStatus		string  `db:"eff_status" cj:"eff_status"`
	Name			string  `db:"name" cj:"name"`
	LastName		string  `db:"last_name" cj:"last_name"`
	MiddleName		string  `db:"middle_name" cj:"middle_name"`
	FirstName		string  `db:"first_name" cj:"first_name"`
	PrefNameFlag	string  `db:"pref_name_flag" cj:"pref_name_flag"`
	UID				string  `db:"u_id" cj:"u_id"`
}

func (name *PersName) StrictValid() bool {
	if name.PersID == "" || name.Name == "" || name.LastName =="" || name.FirstName ==""{
		return false
	}
	return true
}

func (name *PersName) StrictFaildMessage() string {
	var ErrMessage string
	if name.PersID == "" {
		ErrMessage += "pers_id is null; "
	}
	if name.Name =="" {
		ErrMessage += "name is null; "
	}

	if name.LastName =="" {
		ErrMessage += "last_name is null; "
	}

	if name.FirstName =="" {
		ErrMessage += "first_name is null; "
	}

	return ErrMessage
}

func (name *PersName) Valid() bool {
	if name.Name == "" || name.LastName =="" || name.FirstName ==""{
		return false
	}
	return true
}

func (name *PersName) FaildMessage() string {
	var ErrMessage string
	if name.Name =="" {
		ErrMessage += "name is null; "
	}

	if name.LastName =="" {
		ErrMessage += "last_name is null; "
	}

	if name.FirstName =="" {
		ErrMessage += "first_name is null; "
	}

	return ErrMessage
}
