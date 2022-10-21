package params

import (
	"sesi4/server/model"
	"time"
)

type CreateMenuRequest struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

func (c *CreateMenuRequest) ParseToModel() *model.Menu {
	return &model.Menu{
		UserId:   c.UserId,
		Name:     c.Name,
		Category: c.Category,
		Desc:     c.Desc,
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
