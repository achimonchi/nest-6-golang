package model

type Menu struct {
	BaseModel
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	UserId   string `json:"user_id"`

	User User `json:"user"`
}
