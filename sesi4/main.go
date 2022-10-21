package main

import (
	"sesi4/db"
	"sesi4/server"
	"sesi4/server/controller"
	"sesi4/server/repository/gorm_postgres"
	"sesi4/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectGormDB()
	if err != nil {
		panic(err)
	}

	userRepo := gorm_postgres.NewUserRepoGormPostgres(db)
	userSvc := service.NewServices(userRepo)
	userHandler := controller.NewUserHandler(userSvc)

	menuRepo := gorm_postgres.NewMenuRepoGormPostgres(db)
	menuSvc := service.NewMenuServices(menuRepo)
	menuHandler := controller.NewMenuHandler(menuSvc)

	// router := httprouter.New()
	router := gin.Default()
	router.Use(gin.Logger())

	// app := server.NewRouter(router, userHandler)
	app := server.NewRouterGin(router, userHandler, menuHandler)

	app.Start(":4444")
}
