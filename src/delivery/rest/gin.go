package rest

import (
	"fmt"
	"github.com/dalikewara/ayapingping-go/src/service"
	"github.com/dalikewara/rflgo"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ginHandler struct {
	service *service.Service
	client  *gin.Engine
	config  *Config
}

// NewGin generates new gin-gonic REST handler.
func NewGin(param NewGinParam) Server {
	r := &ginHandler{
		service: param.Service,
		client:  param.Client,
		config:  param.Config,
	}

	return r
}

// extractHandlerFuncParam extracts params from handler function.
func (h *ginHandler) extractHandlerFuncParam(params ...interface{}) *gin.Context {
	return params[0].([]interface{})[0].(*gin.Context)
}

// wrapHandlerFunc wraps handler function into gin-gonic handler function.
func (h *ginHandler) wrapHandlerFunc(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(g *gin.Context) {
		_ = handlerFunc(g)
	}
}

// jsonOk writes success json payload into the response body.
func (h *ginHandler) jsonOk(g *gin.Context, data interface{}) {
	status, res := Ok(h.config, data)
	g.JSON(status, res)
}

// jsonNotOk writes error json payload into the response body.
func (h *ginHandler) jsonNotOk(g *gin.Context, httpStatus int, code, message string) {
	status, res := NotOk(h.config, httpStatus, code, message)
	g.JSON(status, res)
}

// RegisterRoutes registers gin-gonic REST routes.
func (h *ginHandler) RegisterRoutes() {
	h.client.GET("/api/v1/user/get-all", h.wrapHandlerFunc(h.APIUserV1GetAll()))
	h.client.GET("/api/v1/role/:userId", h.wrapHandlerFunc(h.APIRoleV1GetByUserID()))
}

// Serve serves gin-gonic REST application.
func (h *ginHandler) Serve() {
	log.Println(fmt.Sprintf("REST start on: env=%s, port=%s, engine=%s", h.config.Env, h.config.Port, "gin-gonic"))
	if h.config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := h.client.Run(fmt.Sprintf(":%v", h.config.Port)); err != nil {
		panic(err)
	}
}

// APIUserV1GetAll handles gin-gonic REST API to get all user data.
func (h *ginHandler) APIUserV1GetAll() HandlerFunc {
	return func(params ...interface{}) interface{} {
		g := h.extractHandlerFuncParam(params)
		ctx := g.Request.Context()

		users := h.service.User.GetAll(service.UserGetAllParam{
			Ctx: ctx,
		})
		if users.Error != nil {
			h.jsonNotOk(g, users.Error.GetHttpStatus(), users.Error.GetCode(), users.Error.GetMessage())
			return nil
		}

		var data []*APIUserV1GetAllData
		if err := rflgo.Compose(&data, users.Users); err != nil {
			h.jsonNotOk(g, ErrorComposeData.GetHttpStatus(), ErrorComposeData.GetCode(), ErrorComposeData.GetMessage())
			return nil
		}

		h.jsonOk(g, data)
		return nil
	}
}

// APIRoleV1GetByUserID handles gin-gonic REST API to get role data by user id.
func (h *ginHandler) APIRoleV1GetByUserID() HandlerFunc {
	return func(params ...interface{}) interface{} {
		g := h.extractHandlerFuncParam(params)
		ctx := g.Request.Context()

		userId, _ := strconv.Atoi(g.Param("userId"))

		role := h.service.Role.GetByUserID(service.RoleGetByUserIDParam{
			UserId: userId,
			Ctx:    ctx,
		})
		if role.Error != nil {
			h.jsonNotOk(g, role.Error.GetHttpStatus(), role.Error.GetCode(), role.Error.GetMessage())
			return nil
		}

		var data *APIRoleV1GetByUserIDData
		if err := rflgo.Compose(&data, role.Role); err != nil {
			h.jsonNotOk(g, ErrorComposeData.GetHttpStatus(), ErrorComposeData.GetCode(), ErrorComposeData.GetMessage())
			return nil
		}

		h.jsonOk(g, data)
		return nil
	}
}
