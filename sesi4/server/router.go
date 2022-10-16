package server

import (
	"log"
	"net/http"
	"sesi4/server/controller"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
	user   *controller.UserHandler
}

func NewRouter(router *httprouter.Router, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/employees", r.user.GetUsers)
	r.router.POST("/employees/register", r.user.Register)

	log.Println("server running at port", port)
	http.ListenAndServe(port, r.router)
}
