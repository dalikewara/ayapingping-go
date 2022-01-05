package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/configs/env"
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	exampleHttpHandler "github.com/dalikewara/ayapingping-go/src/domains/example/delivery/http/ginGonic"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// For testing purpose, we mock the database.
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// Repositories.
	exampleRepository := example.NewMySQLRepository(example.NewMySQLRepositoryParam{
		Db: db,
	})

	// Services.
	exampleService := example.NewService(example.NewServiceParam{
		Repository: exampleRepository,
	})

	// Use cases.
	exampleUseCase := example.NewUseCase(example.NewUseCaseParam{
		ExampleService: exampleService,
	})

	// Delivery API.
	exampleHttpHandler.NewHandler(exampleHttpHandler.NewHandlerParam{
		Router:  router,
		UseCase: exampleUseCase,
	})

	log.Println("App start on env " + env.AppEnv)
	if err = router.Run(); err != nil {
		panic(err)
	}
}
