package service

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase"
)

type Service struct {
	cfg     *config.Config
	adapter *adapter.Adapter
	useCase *usecase.UseCase
}

// Serve serves & runs the service
func (s *Service) Serve() {

	// Run REST http server
	go rest.NewHttpServer(s.cfg, s.adapter.HttpServer, s.useCase).Serve()

	// If you want to use different or other services
	// go rest.NewGinGonic(s.cfg, s.adapter.GinGonicClient, s.useCase).Serve()
	// go cron.NewCron(s.cfg, s.adapter.CronClient, s.useCase).Serve()

	preventMainThreadToExit := make(chan bool)
	<-preventMainThreadToExit
}

// InitService initializes services
func InitService(cfg *config.Config, adapterClient *adapter.Adapter, useCase *usecase.UseCase) (*Service, error) {
	return &Service{
		cfg:     cfg,
		adapter: adapterClient,
		useCase: useCase,
	}, nil
}
