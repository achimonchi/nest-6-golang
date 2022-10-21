package service

import (
	"fmt"
	"sesi4/server/params"
	"sesi4/server/repository"
	"sesi4/server/view"
)

type MenuService struct {
	repo repository.MenuRepo
}

func NewMenuServices(repo repository.MenuRepo) *MenuService {
	return &MenuService{
		repo: repo,
	}
}

func (m *MenuService) CreateMenu(req *params.CreateMenuRequest) *view.Response {
	menu := req.ParseToModel()
	totalMenu, err := m.repo.GetMenus()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	id := fmt.Sprintf("M-%v", len(*totalMenu)+1)
	menu.BaseModel.Id = id
	err = m.repo.CreateMenu(menu)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewMenuCreateResponse(menu))
}

func (m *MenuService) GetMenuById(menuId string) *view.Response {
	menu, err := m.repo.GetMenuById(menuId)
	if err != nil {
		return view.ErrInternalServer(err.Error())

	}

	return view.SuccessFindAll(menu)
}
