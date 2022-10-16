package main

import (
	"sesi4/db"
	"sesi4/server"
	"sesi4/server/controller"
	"sesi4/server/repository/postgres"
	"sesi4/server/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepo := postgres.NewUserRepo(db)
	userSvc := service.NewServices(userRepo)
	userHandler := controller.NewUserHandler(userSvc)

	router := httprouter.New()

	app := server.NewRouter(router, userHandler)

	app.Start(":4000")
}
