package main

import (
	"github.com/gin-gonic/gin"
	routes "jwt-auth/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success: Access granted for api-1",
		})
	})

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success: Access granted for api-2",
		})
	})

	router.Run(":" + port)
}
