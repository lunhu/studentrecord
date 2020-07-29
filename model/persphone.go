package model

type PersPhone struct {
	PersID			string	`db:"pers_id" cj:"pers_id"`
	PhoneType		string  `db:"phone_type" cj:"phone_type"`
	CountryCode		string	`db:"country_code" cj:"country_code"`
	Phone			string  `db:"phone" cj:"phone"`
	Extension		string  `db:"extension" cj:"extension"`
	PrefPhoneFlag	string  `db:"pref_phone_flag" cj:"pref_phone_flag"`
	UID				string  `db:"u_id" cj:"u_id"`
}


func (phone *PersPhone) StrictValid() bool {
	if phone.PersID == "" || phone.Phone =="" {
		return false
	}
	return true
}

func (phone *PersPhone) StrictFaildMessage() string {
	var ErrMessage string
	if phone.PersID =="" {
		ErrMessage += "pers_id is null; "
	}

	if phone.Phone =="" {
		ErrMessage += "phone is null; "
	}

	return ErrMessage
}

func (phone *PersPhone) Valid() bool {
	if phone.Phone =="" {
		return false
	}
	return true
}

func (phone *PersPhone) FaildMessage() string {
	var ErrMessage string

	if phone.Phone =="" {
		ErrMessage += "phone is null; "
	}

	return ErrMessage
}