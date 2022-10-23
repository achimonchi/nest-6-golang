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
	menu   *controller.MenuHandler
}

type GinRouter struct {
	router     *gin.Engine
	user       *controller.UserHandler
	menu       *controller.MenuHandler
	middleware *Middleware
}

func NewRouter(router *httprouter.Router, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}
func NewRouterGin(router *gin.Engine, user *controller.UserHandler, menu *controller.MenuHandler, middleware *Middleware) *GinRouter {
	return &GinRouter{
		router:     router,
		user:       user,
		menu:       menu,
		middleware: middleware,
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
	r.router.Use(r.middleware.Trace)

	emp := r.router.Group("/employees")
	emp.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUsers, []string{"admin"}))
	emp.POST("/register", r.user.GinRegister)
	emp.POST("/login", r.user.GinLogin)

	menu := r.router.Group("/menus")
	menu.POST("/", r.menu.CreateMenu)
	menu.GET("/id/:menuId", r.menu.GetMenuById)

	r.router.Run(port)
}
