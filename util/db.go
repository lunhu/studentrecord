package util

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"studentrecord/common"
)

var (
	// SQLXDB 声明一个 sqlx DB 实例对象
	SQLXDB *sqlx.DB
)

func init() {
	if SQLXDB != nil {
		return
	}

	SQLXDB = NewMySQLXDB()
}

// NewSQLXDB 实例化一个新的 sqlx DB 实例对象
func NewMySQLXDB() *sqlx.DB {
	c := common.GetConfig()
	if c == nil {
		return nil
	}

	dbDrivers, err := c.String("db-mysql", "db_drivers")
	if err != nil {
		panic("无法获取配置文件的 db_drivers :"+err.Error())
	}
	dbConnection, err := c.String("db-mysql", "db_connection")
	if err != nil {
		panic("无法获取配置文件的 db_connection :"+ err.Error())
	}

	maxIdleConn, _ := c.Int("db-mysql", "db_max_idle_conn")
	maxOpenConn, _ := c.Int("db-mysql", "db_max_open_conn")
	connMaxLifetime, _ := c.Int("db-mysql", "db_conn_max_lifetime")

	db, err := sqlx.Open(dbDrivers, dbConnection)
	if err != nil {
		panic("sqlx 初始化数据库出错："+err.Error())
		panic(err.Error())
	}

	db.SetMaxIdleConns(maxIdleConn)                                    //数据库最大闲置数
	db.SetMaxOpenConns(maxOpenConn)                                    //数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second) //数据库最大生命周期

	return db
}
