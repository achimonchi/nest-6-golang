package view

import "sesi4/server/model"

type MenuCreateResponse struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

func NewMenuCreateResponse(menu *model.Menu) *MenuCreateResponse {
	return &MenuCreateResponse{
		Name:     menu.Name,
		Category: menu.Category,
		Desc:     menu.Desc,
	}
}

type MenuGetAllResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

func NewMenuGetAllResponse(menus *[]model.Menu) *[]MenuGetAllResponse {
	var menusResponse []MenuGetAllResponse

	for _, menu := range *menus {
		menusResponse = append(menusResponse, MenuGetAllResponse{
			ID:       menu.Id,
			Name:     menu.Name,
			Category: menu.Category,
			Desc:     menu.Desc,
		})
	}

	return &menusResponse
}
