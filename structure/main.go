package main

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	adp, err := adapter.InitAdapter(cfg)
	if err != nil {
		panic(err)
	}

	repo, err := repository.InitRepository(cfg, adp)
	if err != nil {
		panic(err)
	}

	uc, err := usecase.InitUseCase(cfg, repo)
	if err != nil {
		panic(err)
	}

	svc, err := service.InitService(cfg, adp, uc)
	if err != nil {
		panic(err)
	}

	svc.Serve()
}
