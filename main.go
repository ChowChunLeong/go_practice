package main

import (
	"go_practice/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin's default router
	router := gin.Default()

	designPatternRoute := router.Group("api/designPattern")
	{
		//singleton
		designPatternRoute.POST("/getSingleton", controller.GetSingletonInstance)
		designPatternRoute.POST("/updateSingleton", controller.UpdateSingletonInstance)

	}

	// Run the application
	router.Run(":8080")
}
