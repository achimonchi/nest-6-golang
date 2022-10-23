package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi4/server/params"
	"sesi4/server/service"
	"sesi4/server/view"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	svc *service.UserServices
}

func NewUserHandler(svc *service.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	resp := u.svc.GetUsers()

	WriteJsonResponse(w, resp)

}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.UserCreate
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(w, resp)
		return
	}

	if len(req.Fullname) < 5 {

		resp := view.ErrBadRequest("user name length must be greater than 4")
		WriteJsonResponse(w, resp)
		return
	}

	resp := u.svc.CreateUser(&req)
	WriteJsonResponse(w, resp)
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.UserLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(w, resp)
		return
	}

	resp := u.svc.Login(&req)
	WriteJsonResponse(w, resp)
}

func (u *UserHandler) GinGetUsers(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := u.svc.GetUsers()

	WriteJsonResponseGin(c, resp)

}

func (u *UserHandler) GinRegister(c *gin.Context) {
	var req params.UserCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	// if len(req.Fullname) < 5 {

	// 	resp := view.ErrBadRequest("user name length must be greater than 4")
	// 	WriteJsonResponseGin(c, resp)
	// 	return
	// }

	resp := u.svc.CreateUser(&req)
	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) GinLogin(c *gin.Context) {
	var req params.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := u.svc.Login(&req)
	WriteJsonResponseGin(c, resp)
}
