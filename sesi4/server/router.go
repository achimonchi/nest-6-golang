package server

import (
	"log"
	"net/http"
	"sesi4/server/controller"
)

func Start(port string) {
	http.HandleFunc("/users", controller.GetUsers)
	http.HandleFunc("/users/register", controller.Register)

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}
