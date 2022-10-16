package service

import (
	"sesi4/server/params"
	"sesi4/server/repository"
	"sesi4/server/view"
	"time"

	"github.com/google/uuid"
)

type UserServices struct {
	repo repository.UserRepo
}

func NewServices(repo repository.UserRepo) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) CreateUser(req *params.UserCreate) *view.Response {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := u.repo.Register(user)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewUserCreateResponse(user))
}
