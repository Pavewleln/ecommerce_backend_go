package routes

import (
	"github.com/gin-gonic/gin"
	"main/src/controllers"
	"main/src/middlewares"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	auth := baseRouter.Group("/auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	auth.PATCH("/edit", middlewares.AuthMiddleware(), controllers.Edit)
	auth.GET("/logout", middlewares.AuthMiddleware(), controllers.Logout)
	auth.POST("/refresh", controllers.Refresh)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api")
	startupsGroupRouter(versionRouter)

	return r
}
