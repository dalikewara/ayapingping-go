package mysql_test

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/src/databases/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// dialect should generates valid string of data source name.
// dialect is required to connect to the real database.
// You may change the value to suit your environments and to pass
// the unit testing.
func dialect() string {
	return "root:p455w0rd@tcp(localhost:3306)/mysql"
}

// TestConnect tests `mysql.Connect` method
// and its possible scenarios.
func TestConnect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, err := mysql.Connect(dialect())
		defer func(db *sql.DB, t *testing.T) {
			err := db.Close()
			assert.Nil(t, err)
		}(db, t)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(&sql.DB{}, db)
	})
	t.Run("`Open` error", func(t *testing.T) {
		errDialect := "error"
		db, err := mysql.Connect(errDialect)
		assert.Nil(t, db)
		assert.Error(t, err)
	})
	t.Run("`Ping` error", func(t *testing.T) {
		errDialect := "root@tcp(localhost:0000)/mysql"
		db, err := mysql.Connect(errDialect)
		assert.ObjectsAreEqual(&sql.DB{}, db)
		assert.Error(t, err)
	})
}
