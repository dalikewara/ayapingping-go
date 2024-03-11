package common

import (
	"database/sql"
	"fmt"
)

func ConnectMySQL(host string, port string, user string, pass string, dbName string) (*sql.DB, error) {
	if pass != "" && pass != " " {
		pass = ":" + pass
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
