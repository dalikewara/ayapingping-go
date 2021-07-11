package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/app/api/services/user_example_service"
	"github.com/dalikewara/ayapingping-go/domain/user_example/repository"
	"github.com/dalikewara/ayapingping-go/domain/user_example/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// For testing & example purpose, we just mock the database.
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// User example services.
	userExampleRepo := repository.NewMySQL(db)
	userExampleUseCase := usecase.NewUseCase(userExampleRepo)
	user_example_service.NewService(router, userExampleUseCase)

	if err = router.Run(); err != nil {
		panic(err)
	}
}
