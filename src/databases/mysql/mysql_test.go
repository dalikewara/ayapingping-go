package mysql_test

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/src/databases/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConnect tests mysql.Connect function and all possible scenarios.
func TestConnect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, err := mysql.Connect(mysql.ConnectParam{
			Host:   "localhost",
			Port:   "3306",
			User:   "root",
			Pass:   "p455w0rd",
			DBName: "mysql",
		})
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
			Host:   "localhost",
			Port:   "0000",
			User:   "root",
			DBName: "mysql",
		})
		assert.ObjectsAreEqual(&sql.DB{}, db)
		assert.Error(t, err)
	})
}
