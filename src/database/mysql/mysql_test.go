package mysql_test

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/src/database/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConnect tests mysql.Connect function and all possible scenarios.
func TestConnect(t *testing.T) {
	validConnection := mysql.ConnectParam{
		Host:   "localhost",
		Port:   "3306",
		User:   "root",
		Pass:   "",
		DBName: "test",
	}

	t.Run("ok", func(t *testing.T) {
		db, err := mysql.Connect(validConnection)
		defer func(db *sql.DB, t *testing.T) {
			err := db.Close()
			assert.Nil(t, err)
		}(db, t)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(&sql.DB{}, db)
	})
	t.Run("`Open` errs", func(t *testing.T) {
		db, err := mysql.Connect(mysql.ConnectParam{
			User: "errs",
		})
		assert.Nil(t, db)
		assert.Error(t, err)
	})
	t.Run("`Ping` errs", func(t *testing.T) {
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
