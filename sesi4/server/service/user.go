package service

import (
	"database/sql"
	"sesi4/helper"
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

func (u *UserServices) GetUsers() *view.Response {
	users, err := u.repo.GetUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewUserFindAllResponse(users))
}

func (u *UserServices) CreateUser(req *params.UserCreate) *view.Response {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	user.Password = hash

	err = u.repo.Register(user)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewUserCreateResponse(user))
}

func (u *UserServices) Login(req *params.UserLogin) *view.Response {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return view.ErrUnauthorized()
	}

	token := helper.Token{
		UserId: user.Id,
		Email:  user.Email,
	}

	tokString, err := helper.CreateToken(&token)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(tokString)
}
