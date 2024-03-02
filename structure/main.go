package main

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/v4/structure/commons/env"
	"github.com/dalikewara/ayapingping-go/v4/structure/commons/mysql"
	"github.com/dalikewara/ayapingping-go/v4/structure/commons/netHttp"
	"github.com/dalikewara/ayapingping-go/v4/structure/features/example/delivery/handlers/exampleGet"
	"github.com/dalikewara/ayapingping-go/v4/structure/features/example/repositories/findExampleByID"
	"github.com/dalikewara/ayapingping-go/v4/structure/features/example/usecases/getExample"
)

func main() {
	// Parse env

	envCfg, err := env.Parse()
	if err != nil {
		panic(err)
	}

	// Database connection

	mysqlDB, err := mysql.Connect()
	if err != nil {
		panic(err)
	}

	// Http server initialization

	httpServer := netHttp.Server()
	httpServerMux := netHttp.ServerMux()

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