// user API application executable using gin-gonic.

package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/v2/src/config/constant"
	"github.com/dalikewara/ayapingping-go/v2/src/config/env"
	"github.com/dalikewara/ayapingping-go/v2/src/delivery/rest"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
	"github.com/dalikewara/ayapingping-go/v2/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// We mock the database for this case.
	mysqlDB, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// Business repositories.
	repositories := repository.New(repository.NewParam{
		MySQLDB: mysqlDB,
	})

	// Business services.
	serviceConfig := &service.Config{
		SystemUserRole: constant.SystemUserRole,
		SystemUserIds:  constant.SystemUserIds,
	}
	services := service.New(service.NewParam{
		Repo:   repositories,
		Config: serviceConfig,
	})

	// Server handler.
	serverConfig := rest.Config{
		Env:             env.AppEnv,
		Port:            env.AppPort,
		OkCode:          constant.RESTOkCode,
		OkMessage:       constant.RESTOkMessage,
		OkHttpStatus:    constant.RESTOkHttpStatus,
		NotOkHttpStatus: constant.RESTNotOkHttpStatus,
	}
	server := rest.NewGin(rest.NewGinParam{
		Service: services,
		Client:  gin.Default(),
		Config:  &serverConfig,
	})
	server.RegisterRoutes()
	server.Serve()
}
