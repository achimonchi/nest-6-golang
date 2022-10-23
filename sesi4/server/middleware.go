package server

import (
	"log"
	"sesi4/helper"
	"sesi4/server/controller"
	"sesi4/server/model"
	"sesi4/server/service"
	"sesi4/server/view"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userSvc *service.UserServices
}

func NewMiddleware(userSvc *service.UserServices) *Middleware {
	return &Middleware{
		userSvc: userSvc,
	}
}

func (m *Middleware) Auth(c *gin.Context) {
	// get bearer token
	bearerToken := c.GetHeader("Authorization") // Bearer <token>

	// get token
	tokenArr := strings.Split(bearerToken, "Bearer ")

	// validate
	if len(tokenArr) != 2 {
		controller.WriteErrorJsonResponseGin(c, view.ErrUnauthorized())
		return
	}

	// verify token
	myTok, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		controller.WriteErrorJsonResponseGin(c, view.ErrUnauthorized())
		return
	}

	// send to next handler
	c.Set("USER_ID", myTok.UserId)
	c.Set("USER_EMAIL", myTok.Email)

	// process to another handler
	c.Next()

}

func (m *Middleware) CheckRole(next gin.HandlerFunc, roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("USER_EMAIL")
		user := m.userSvc.FindUserByEmail(email)
		userDetail := user.Payload.(*model.User)

		isExist := false

		for _, role := range roles {
			if role == userDetail.Role {
				isExist = true
				break
			}
		}

		if !isExist {
			controller.WriteErrorJsonResponseGin(ctx, view.ErrUnauthorized())
			return
		}

		next(ctx)
	}
}

func (m *Middleware) Trace(c *gin.Context) {
	now := time.Now()

	c.Next()

	end := time.Since(now).Milliseconds()
	log.Println("response time:", end)
}
