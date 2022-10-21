package repository

import "sesi4/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type MenuRepo interface {
	GetMenus() (*[]model.Menu, error)
	CreateMenu(m *model.Menu) error
	GetMenuById(menuId string) (*model.Menu, error)
}
