package routes

import (
	"github.com/gin-gonic/gin"
	controllers "jwt-auth/controllers"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
