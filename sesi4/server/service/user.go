package service

import (
	"sesi4/server/model"
	"sesi4/server/params"
	"sesi4/server/view"
)

func CreateUser(req *params.UserCreate) *view.Response {
	user := req.ParseToModel()

	user.Id = len(model.Users) + 1
	model.Users = append(model.Users, *user)

	return view.SuccessCreated(view.NewUserCreateResponse(user))
}
