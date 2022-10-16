package mysql

import (
	"fmt"
	"go-template/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	// 也可以使用MustConnect连接不成功就panic
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	Db.SetMaxOpenConns(cfg.MaxOpenConns)
	Db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = Db.Close()
}
