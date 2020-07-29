package service

import (
	"studentrecord/util"
	"github.com/jmoiron/sqlx"
)


var (
	sqlxDB *sqlx.DB
)

func init() {
	sqlxDB = util.SQLXDB
	if sqlxDB == nil {
		sqlxDB = util.NewMySQLXDB()
	}
}