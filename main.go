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

		//abstract factory design pattern
		designPatternRoute.POST("/useFurniture", controller.UseFurniture)

		//builder pattern
		designPatternRoute.POST("/orderBurger", controller.OrderBurger)
	}

	router.GET("/ws", controller.HandleWebSocket)

	// Run the application
	go router.Run(":8080")
	router.Run(":8081") // Port for WebSocket
}
