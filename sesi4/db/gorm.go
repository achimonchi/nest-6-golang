package db

import (
	"fmt"
	"sesi4/server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbs, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(model.User{}, model.Menu{})

	return db, nil
}
