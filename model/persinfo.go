package model

type PersInfo struct {
	PersID			string	`db:"pers_id" cj:"pers_id"`
	Names	[]*PersName
	Phones  []*PersPhone
	Emails  []*PersEmail
	Employs []*PersEmploy
	OprDefn *PersOprDefn
}