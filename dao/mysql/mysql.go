package mysql

import (
	"fmt"
	"web_app/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var driverName string = "mysql"

func Init(cfg settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.PassWord,
		cfg.Host,
		cfg.Port,
		cfg.DbName)
	DB, err = sqlx.Connect(driverName, dsn)
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(cfg.MaxOpenConn)
	DB.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}
