package ginGonic

import (
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIGetAll serves http API to get all user data.
func (h *Handler) APIGetAll(g *gin.Context) {
	res := user.HttpResponse{}
	users := h.useCase.GetAll(user.UseCaseGetAllParam{})
	if users.Error != nil {
		res.Error = true
		res.Message = users.Error.Error()
		g.JSON(http.StatusOK, res)
		return
	}
	res.Data = users.Users
	res.Message = "ok"
	g.JSON(http.StatusOK, res)
	return
}

// APICreate serves http API to create user data.
func (h *Handler) APICreate(g *gin.Context) {
	res := user.HttpResponse{}
	req := user.HttpCreateRequest{}
	ctx := g.Request.Context()
	if err := g.Bind(&req); err != nil {
		res.Error = true
		res.Message = err.Error()
		g.JSON(http.StatusBadRequest, res)
		return
	}
	reply := h.useCase.Create(user.UseCaseCreateParam{
		Username: req.Username,
		Ctx:      ctx,
	})
	if reply.Error != nil {
		res.Error = true
		res.Message = reply.Error.Error()
		g.JSON(http.StatusOK, res)
		return
	}
	res.Message = "ok"
	g.JSON(http.StatusOK, res)
	return
}
