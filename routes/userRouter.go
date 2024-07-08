package routes

import (
	"github.com/gin-gonic/gin"
	controller "jwt-auth/controllers"
	"jwt-auth/middleware"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	// incomingRoutes.POST("users/signup", controller.Signup())
	// incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.GET("users", middleware.AuthMiddleware(), controller.GetUsers())
	incomingRoutes.GET("users/:user_id", middleware.AuthMiddleware(), controller.GetUser())
	// incomingRoutes.PUT("users/:id", middleware.AuthMiddleware(), controller.UpdateUser())
	// incomingRoutes.DELETE("users/:id", middleware.AuthMiddleware(), controller.DeleteUser())

}
