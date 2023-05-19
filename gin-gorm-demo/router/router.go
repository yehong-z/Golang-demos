package router

import (
	"gin-gorm-demo/handlers/user_login"
	"gin-gorm-demo/middleware"
	"gin-gorm-demo/models"
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
