package mysql_test

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/src/configs/env"
	"github.com/dalikewara/ayapingping-go/src/databases/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validConnection = mysql.ConnectParam{
	Host:   env.MySQLHost,
	Port:   env.MySQLPort,
	User:   env.MySQLUser,
	Pass:   env.MySQLPass,
	DBName: env.MySQLDBName,
}

// TestConnect tests mysql.Connect function and all possible scenarios.
func TestConnect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, err := mysql.Connect(validConnection)
		defer func(db *sql.DB, t *testing.T) {
			err := db.Close()
			assert.Nil(t, err)
		}(db, t)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(&sql.DB{}, db)
	})
	t.Run("`Open` error", func(t *testing.T) {
		db, err := mysql.Connect(mysql.ConnectParam{
			User: "error",
		})
		assert.Nil(t, db)
		assert.Error(t, err)
	})
	t.Run("`Ping` error", func(t *testing.T) {
		db, err := mysql.Connect(mysql.ConnectParam{
			Host:   validConnection.Host,
			Port:   "0000",
			User:   validConnection.User,
			DBName: validConnection.DBName,
		})
		assert.ObjectsAreEqual(&sql.DB{}, db)
		assert.Error(t, err)
	})
}

// TestConnectMock tests mysql.ConnectMock function.
func TestConnectMock(t *testing.T) {
	db, sqlMock, err := mysql.ConnectMock()
	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.NotNil(t, sqlMock)
	assert.ObjectsAreEqual(sql.DB{}, db)
}
