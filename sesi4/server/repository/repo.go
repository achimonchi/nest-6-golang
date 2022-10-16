package repository

import "sesi4/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
}
