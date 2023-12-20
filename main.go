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
		//Singleton
		designPatternRoute.POST("/getSingleton", controller.GetSingletonInstance)
		designPatternRoute.POST("/updateSingleton", controller.UpdateSingletonInstance)

		//Factory pattern
		designPatternRoute.POST("/drawShape", controller.DrawShape)
	}

	// Run the application
	router.Run(":8080")
}
