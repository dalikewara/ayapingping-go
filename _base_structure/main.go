package main

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/common"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/features/example"
	"net/http"
)

func main() {
	// Http server initialization

	httpServer := common.NewNetHttpServer()
	httpServerMux := common.NewNetHttpServerMux()

	// Repositories

	exampleRepositoryMySQL := example.NewRepositoryMySQL(nil)

	// Use cases

	exampleUseCaseV1 := example.NewUseCaseV1(exampleRepositoryMySQL)

	// Services

	exampleHttpServiceNetHttp := example.NewHttpServiceNetHttp(httpServerMux, exampleUseCaseV1)

	// Service handlers

	exampleHttpServiceNetHttp.ExampleDetail(http.MethodGet, "/example")

	// Start & listen application

	httpServer.Handler = httpServerMux
	httpServer.Addr = ":8080"

	fmt.Println("App running on port: 8080")

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
