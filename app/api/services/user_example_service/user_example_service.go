package user_example_service

import (
	"context"
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Service object.
type Service struct {
	useCase user_example.UseCase
}

// NewService generates new `gin-gonic` service handler.
func NewService(router *gin.Engine, useCase user_example.UseCase) {
	service := &Service{
		useCase: useCase,
	}
	router.POST("/user/login", service.Login)
	router.POST("/user/login-context", service.LoginContext)
}

// Login serves example api for login (`/user/login`).
func (s *Service) Login(g *gin.Context) {
	request := &user_example.APIServiceLoginRequest{}
	if err := g.ShouldBindJSON(&request); err != nil {
		g.JSON(http.StatusBadRequest, user_example.APIServiceLoginResponse{
			Code: "InvalidRequest",
			Message: err.Error(),
		})
		return
	}
	if request.Username != "guest" && request.Password != "guest" {
		g.JSON(http.StatusUnauthorized, user_example.APIServiceLoginResponse{
			Code: "NotLoggedIn",
			Message: "username must be 'guest' and password must be 'guest'",
		})
		return
	}
	useCaseResponse := s.useCase.Login(user_example.UseCaseLoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	g.JSON(http.StatusOK, user_example.APIServiceLoginResponse{
		Code: "LoggedIn",
		Message: "login success",
		Username: useCaseResponse.User.Username,
	})
}

// LoginContext serves example api for login with context (`/user/login-context`)
func (s *Service) LoginContext(g *gin.Context) {
	ctx := g.Request.Context()
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	request := &user_example.APIServiceLoginContextRequest{}
	if err := g.ShouldBindJSON(&request); err != nil {
		g.JSON(http.StatusBadRequest, user_example.APIServiceLoginContextResponse{
			Code: "InvalidRequest",
			Message: err.Error(),
		})
		return
	}
	if request.Username != "guest" && request.Password != "guest" {
		g.JSON(http.StatusUnauthorized, user_example.APIServiceLoginContextResponse{
			Code: "NotLoggedIn",
			Message: "username must be 'guest' and password must be 'guest'",
		})
		return
	}
	useCaseResponse := s.useCase.LoginContext(user_example.UseCaseLoginContextRequest{
		Ctx: ctxWithTimeout,
		Username: request.Username,
		Password: request.Password,
	})
	if useCaseResponse.Error != nil {
		g.JSON(http.StatusBadGateway, user_example.APIServiceLoginContextResponse{
			Code: "NotLoggedIn",
			Message: useCaseResponse.Error.Error() + ": this means our login with context example scenario is working",
			Username: request.Username,
		})
		return
	}
	g.JSON(http.StatusOK, user_example.APIServiceLoginContextResponse{
		Code: "LoggedIn",
		Message: "login success, but you should not see this",
		Username: useCaseResponse.User.Username,
	})
}
