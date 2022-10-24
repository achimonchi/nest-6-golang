package service

import (
	"database/sql"
	"log"
	"sesi4/adaptor"
	"sesi4/helper"
	"sesi4/server/params"
	"sesi4/server/repository"
	"sesi4/server/view"
	"time"

	"github.com/google/uuid"
)

type UserServices struct {
	repo            repository.UserRepo
	typicodeAdaptor *adaptor.TypicodeAdaptor
}

func NewServices(repo repository.UserRepo, typicodeAdaptor *adaptor.TypicodeAdaptor) *UserServices {
	return &UserServices{
		repo:            repo,
		typicodeAdaptor: typicodeAdaptor,
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
	user.Role = "member"

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("get error when try to generate password %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	user.Password = hash

	err = u.repo.Register(user)
	if err != nil {
		log.Printf("get error register user with error %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	data, err := u.typicodeAdaptor.GetAllTypicode()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(data)
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

func (u *UserServices) FindUserByEmail(email string) *view.Response {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}
	return view.SuccessFindAll(user)
}
