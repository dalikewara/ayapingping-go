package rest

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest/handler/exampleGet"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase"
	"net/http"
)

// Routes registers REST routes
func Routes(cfg *config.Config, useCase *usecase.UseCase) map[string][]map[string]interface{} {
	return map[string][]map[string]interface{}{
		"": {
			{
				"method": "GET",
				"path":   "/example",
				"httpServerHandlers": []http.HandlerFunc{
					exampleGet.NewV1HttpServer(useCase.GetExample).JSON,
				},
				// If you want to use different or other route handlers
				//"ginGonicHandlers": []gin.HandlerFunc{
				//	ping.NewV1GinGonic().String,
				//},
			},
		},
		// If you want to add more route path groups. For example: `GET /api/v1/user/get-detail`
		//"/api/v1/user": {
		//	{
		//		"method": "GET",
		//		"path":   "/get-detail",
		//		"httpServerHandlers": []http.HandlerFunc{
		//			exampleGet.NewV1HttpServer(useCase.GetUserDetail).JSON,
		//		},
		//	},
		//},
	}
}
