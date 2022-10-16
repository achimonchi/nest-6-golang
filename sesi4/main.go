package main

import (
	"log"
	"sesi4/db"
	"sesi4/server"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := httprouter.New()

	app := server.NewRouter(router)

	app.Start(":4000")
}
