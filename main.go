package main

import (
	"go_practice/controller"
	"go_practice/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin's default router
	router := gin.Default()
	database.SetUpRedisDatabaseConnection()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "content-type", "OS", "Browser_Name", "remote-site"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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

	wsRouter := router.Group("ws")
	{
		wsRouter.GET("", controller.HandleWebSocket)
		wsRouter.POST("/changeRedisClientId", controller.ChangeRedisClientId)

	}

	// Run the application
	go router.Run(":8080")
	router.Run(":8081") // Port for WebSocket
}
