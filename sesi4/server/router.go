package server

import (
	"log"
	"net/http"
	"sesi4/server/controller"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
}

func NewRouter(router *httprouter.Router) *Router {
	return &Router{
		router: router,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/users", controller.GetUsers)
	r.router.POST("/users", controller.Register)

	log.Println("server running at port", port)
	http.ListenAndServe(port, r.router)
}
