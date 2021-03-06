package dbx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const DefaultMySQLDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true"

func ConnectMySQL(format string, data DSN) (err error) {
	dsn := fmt.Sprintf(format, data.User, data.Password, data.Host, data.Port, data.DBName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.WithStack(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(30)
	DB = &DatabaseSQLX{DB: db}
	return errors.WithStack(DB.Ping())
}
