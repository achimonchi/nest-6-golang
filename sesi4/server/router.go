package server

import (
	"log"
	"net/http"
	"sesi4/server/controller"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
	user   *controller.UserHandler
}

type GinRouter struct {
	router *gin.Engine
	user   *controller.UserHandler
}

func NewRouter(router *httprouter.Router, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}
func NewRouterGin(router *gin.Engine, user *controller.UserHandler) *GinRouter {
	return &GinRouter{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/employees", r.user.GetUsers)
	r.router.POST("/employees/register", r.user.Register)
	r.router.POST("/employees/login", r.user.Login)

	log.Println("server running at port", port)
	http.ListenAndServe(port, r.router)
}

func (r *GinRouter) Start(port string) {
	emp := r.router.Group("/employees")
	emp.GET("/", r.user.GinGetUsers)
	emp.POST("/register", r.user.GinRegister)
	emp.POST("/login", r.user.GinLogin)

	r.router.Run(port)
}
