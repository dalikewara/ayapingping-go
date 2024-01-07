package adapter

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter/httpServer"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter/mysql"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"net/http"
)

type Adapter struct {
	MySQLDB    *sql.DB
	HttpServer *http.Server
}

// InitAdapter initializes frameworks or resources adapter
func InitAdapter(cfg *config.Config) (*Adapter, error) {

	// Init MySQL database connection
	mysqlDB, err := mysql.NewMySQL(cfg.MySQLConnString)
	if err != nil {
		return nil, err
	}

	// Init Http Server connection
	httpServerConn := httpServer.NewHttpServer()

	return &Adapter{
		MySQLDB:    mysqlDB,
		HttpServer: httpServerConn,
	}, nil
}
