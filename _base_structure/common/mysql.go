package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectMySQL(host string, port string, user string, pass string, dbName string) (*sqlx.DB, error) {
	if pass != "" && pass != " " {
		pass = ":" + pass
	}

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
