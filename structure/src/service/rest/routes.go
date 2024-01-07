package rest

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest/handler/ping"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest/handler/userGetDetail"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase"
	"net/http"
)

// Routes registers REST routes
func Routes(cfg *config.Config, useCase *usecase.UseCase) map[string][]map[string]interface{} {
	return map[string][]map[string]interface{}{
		"": {
			{
				"method": "GET",
				"path":   "/ping",
				"httpServerHandlers": []http.HandlerFunc{
					ping.NewV1HttpServer().String,
				},
				// If you want to use different or other route handlers
				//"ginGonicHandlers": []gin.HandlerFunc{
				//	ping.NewV1GinGonic().String,
				//},
			},
		},
		"/api/v1/user": {
			{
				"method": "GET",
				"path":   "/detail",
				"httpServerHandlers": []http.HandlerFunc{
					userGetDetail.NewV1HttpServer(useCase.GetUserDetail).JSON,
				},
			},
		},
	}
}
