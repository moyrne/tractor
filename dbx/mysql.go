package dbx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const DefaultMySQLDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true"

func ConnectMySQL(format string, data DSN) (db Database, err error) {
	dsn := fmt.Sprintf(format, data.User, data.Password, data.Host, data.Port, data.DBName)
	dbMySQL, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dbMySQL.SetMaxIdleConns(5)
	dbMySQL.SetMaxOpenConns(30)
	db = &DatabaseSQLX{DB: dbMySQL}
	return db, errors.WithStack(db.Ping())
}
