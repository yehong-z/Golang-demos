package router

import (
	"demo/handlers/user_login"
	"demo/middleware"
	"demo/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()
	baseGroup := r.Group("/demo")
	baseGroup.POST("/sessions/", middleware.SHAMiddleWare(), user_login.UserLoginHandler)
	baseGroup.POST("/users", middleware.SHAMiddleWare(), user_login.UserRegisterHandler)
	return r
}
