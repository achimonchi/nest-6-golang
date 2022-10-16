package params

import "sesi4/server/model"

type UserCreate struct {
	Fullname string
	Email    string
	Password string
}

type UserUpdate struct {
	Email    string
	Password string
}

func (u *UserCreate) ParseToModel() *model.User {
	return &model.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
	}
}
