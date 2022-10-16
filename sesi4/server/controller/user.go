package controller

import (
	"encoding/json"
	"net/http"
	"sesi4/server/model"
	"sesi4/server/params"
	"sesi4/server/service"
	"sesi4/server/view"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	users := model.Users

	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "NOT_FOUND",
		})
		return
	}

	usersFind := view.NewUserFindAllResponse(users)

	resp := view.SuccessFindAll(usersFind)

	WriteJsonResponse(w, resp)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"payload": users,
	// })
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	resp := service.CreateUser(&req)
	WriteJsonResponse(w, resp)
}
