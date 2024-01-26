package rest

// If you want to use different or other services

//type ginGonicHandler struct {
//	cfg     *config.Config
//	client  *gin.Engine
//	useCase *usecase.UseCase
//}

// Serve serves & runs the gin gonic
//func (g *ginGonicHandler) Serve() {
//	// Your gin gonic handler logic here...
//
//	for rootPath, endpoints := range Routes(g.cfg, g.useCase) {
//		for _, endpoint := range endpoints {
//			g.client.Handle(endpoint["method"].(string), rootPath+endpoint["path"].(string), endpoint["ginGonicHandlers"].([]gin.HandlerFunc)...)
//		}
//	}
//
//	if err := g.client.Run(":" + g.cfg.RESTPort); err != nil {
//		panic(err)
//	}
//}

// NewGinGonic creates new gin gonic
//func NewGinGonic(cfg *config.Config, client *gin.Engine, useCase *usecase.UseCase) Contract {
//	return &ginGonicHandler{
//		cfg:     cfg,
//		client:  client,
//		useCase: useCase,
//	}
//}
