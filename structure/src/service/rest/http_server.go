package rest

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase"
	"net/http"
)

type httpServerHandler struct {
	cfg     *config.Config
	client  *http.Server
	useCase *usecase.UseCase
}

// Serve serves & runs the http server
func (g *httpServerHandler) Serve() {
	handler := http.NewServeMux()

	for rootPath, endpoints := range Routes(g.cfg, g.useCase) {
		for _, endpoint := range endpoints {
			handler.HandleFunc(rootPath+endpoint["path"].(string), endpoint["httpServerHandlers"].([]http.HandlerFunc)[0])
		}
	}

	g.client.Handler = handler
	g.client.Addr = ":" + g.cfg.RESTPort

	fmt.Println("App running on port: ", g.cfg.RESTPort)

	if err := g.client.ListenAndServe(); err != nil {
		panic(err)
	}
}

// NewHttpServer creates new http server
func NewHttpServer(cfg *config.Config, client *http.Server, useCase *usecase.UseCase) Contract {
	return &httpServerHandler{
		cfg:     cfg,
		client:  client,
		useCase: useCase,
	}
}
