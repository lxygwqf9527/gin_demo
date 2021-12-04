package mysql

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sql.DB

func Init() (err error) {
	dsn := fmt.Printf("%s:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8umb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt()
		)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

func Close(){
	_ = db.Close()
}