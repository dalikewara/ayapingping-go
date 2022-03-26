// user API application executable using gin-gonic.

package main

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/src/configs/env"
	"github.com/dalikewara/ayapingping-go/src/databases/mysql"
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	userHttpHandler "github.com/dalikewara/ayapingping-go/src/domains/user/delivery/http/ginGonic"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	appPort := env.AppPort
	appEnv := env.AppEnv
	router := gin.Default()

	// We mock the database for this case.
	db, _, err := mysql.ConnectMock()
	if err != nil {
		panic(err)
	}

	// Business repositories.
	userRepository := user.NewMySQLRepository(user.NewMySQLRepositoryParam{
		Db: db,
	})

	// Business services.
	userService := user.NewService(user.NewServiceParam{
		Repository: userRepository,
	})

	// Business use cases.
	userUseCase := user.NewUseCase(user.NewUseCaseParam{
		Service: userService,
	})

	// API delivery handler.
	userHttpHandler.NewHandler(userHttpHandler.NewHandlerParam{
		Router:  router,
		UseCase: userUseCase,
	})

	// Run the server.
	log.Println(fmt.Sprintf("APP start on: env=%s, port=%s", appEnv, appPort))
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	if err = router.Run(fmt.Sprintf(":%v", appPort)); err != nil {
		panic(err)
	}
}
