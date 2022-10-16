package params

import "sesi4/server/model"

type UserCreate struct {
	Fullname string
	Email    string
	Password string
	Address  string
	NIP      string
}

type UserUpdate struct {
	Email    string
	Password string
}
type UserLogin struct {
	Email    string
	Password string
}

func (u *UserCreate) ParseToModel() *model.User {
	return &model.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Nip:      u.NIP,
		Address:  u.Address,
	}
}
