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

	// router := httprouter.New()
	router := gin.New()
	router.Use(gin.Logger())

	// app := server.NewRouter(router, userHandler)
	app := server.NewRouterGin(router, userHandler)

	app.Start(":4444")
}
