package main

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/common"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/delivery/handlers/exampleGet"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/repositories/findExampleByID"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/usecases/getExample"
)

func main() {
	// Parse env

	envCfg, err := common.ParseEnv()
	if err != nil {
		//panic(err)
	}

	// Database connection

	mysqlDB, err := common.ConnectMySQL(envCfg.MySQLHost, envCfg.MySQLPort, envCfg.MySQLUser, envCfg.MySQLPass, envCfg.MySQLDBName)
	if err != nil {
		//panic(err)
	}

	// Http server initialization

	httpServer := common.NewNetHttpServer()
	httpServerMux := common.NewNetHttpServerMux()

	// Repositories

	findExampleByIDRepositoryMySQL := findExampleByID.NewMySQL(mysqlDB)

	// Use cases

	getExampleUseCaseV1 := getExample.NewV1(findExampleByIDRepositoryMySQL)

	// Register handlers

	exampleGet.NewV1NetHttp(httpServerMux, getExampleUseCaseV1).RegisterHandler("GET", "/api/v1/example/get")

	// Start & listen application

	httpServer.Handler = httpServerMux
	httpServer.Addr = ":" + envCfg.RESTPort

	fmt.Println("App running on port: ", envCfg.RESTPort)

	if err = httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
