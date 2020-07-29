package model

type PersOprDefn struct{
	PersID			string	`db:"pers_id" cj:"pers_id"`
	OprID			string  `db:"opr_id" cj:"opr_id"`
	UID				string	`db:"u_id" cj:"u_id"`
}

