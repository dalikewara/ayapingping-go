package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/apps/api/apis"
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// For testing purpose, we just mock the database.
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
	apis.NewAPIExample(apis.NewAPIExampleParam{
		Router: router,
		ExampleUseCase: exampleUseCase,
	})

	if err = router.Run(); err != nil {
		panic(err)
	}
}