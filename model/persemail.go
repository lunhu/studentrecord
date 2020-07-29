package model

type PersEmail struct {
	PersID			string	`db:"pers_id" cj:"pers_id"`
	EmailType		string  `db:"email_type" cj:"email_type"`
	Email			string  `db:"email" cj:"email"`
	PrefEmailFlag	string  `db:"pref_emial_flag" cj:"pref_email_flag"`
	UID				string  `db:"u_id" cj:"u_id"`
}

func (email *PersEmail) StrictValid() bool {
	if email.PersID == "" || email.Email =="" {
		return false
	}
	return true
}

func (email *PersEmail) StrictFaildMessage() string {
	var ErrMessage string
	if email.PersID =="" {
		ErrMessage += "pers_id is null; "
	}

	if email.Email =="" {
		ErrMessage += "email is null; "
	}

	return ErrMessage
}

func (email *PersEmail) Valid() bool {
	if email.Email =="" {
		return false
	}
	return true
}

func (email *PersEmail) FaildMessage() string {
	var ErrMessage string

	if email.Email =="" {
		ErrMessage += "email is null; "
	}

	return ErrMessage
}