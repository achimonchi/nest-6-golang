package controller

import (
	"sesi4/server/params"
	"sesi4/server/service"
	"sesi4/server/view"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	svc *service.MenuService
}

func NewMenuHandler(svc *service.MenuService) *MenuHandler {
	return &MenuHandler{
		svc: svc,
	}
}

func (m *MenuHandler) CreateMenu(c *gin.Context) {
	var req params.CreateMenuRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, view.ErrBadRequest(err.Error()))
		return
	}

	resp := m.svc.CreateMenu(&req)

	WriteJsonResponseGin(c, resp)
}

func (m *MenuHandler) GetMenuById(c *gin.Context) {
	menuId, isExist := c.Params.Get("menuId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("menuId not found in params"))
		return
	}

	resp := m.svc.GetMenuById(menuId)
	WriteJsonResponseGin(c, resp)
}
