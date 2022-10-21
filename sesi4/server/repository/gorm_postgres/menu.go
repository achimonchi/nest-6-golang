package gorm_postgres

import (
	"sesi4/server/model"
	"sesi4/server/repository"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepoGormPostgres(db *gorm.DB) repository.MenuRepo {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenus() (*[]model.Menu, error) {
	var menus []model.Menu

	err := m.db.Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return &menus, nil
}

func (m *menuRepo) CreateMenu(menu *model.Menu) error {
	return m.db.Create(menu).Error
}

func (m *menuRepo) GetMenuById(menuId string) (*model.Menu, error) {
	var menu model.Menu
	err := m.db.Joins("User").Where("menus.id=?", menuId).First(&menu).Error
	if err != nil {
		return nil, err
	}

	return &menu, nil
}
