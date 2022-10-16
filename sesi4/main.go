package main

import (
	"log"
	"sesi4/db"
	"sesi4/server"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}
	server.Start(":4000")
}
