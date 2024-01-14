package controller

import (
	"fmt"
	"go_practice/helper"
	"go_practice/service/builderDesignPattern/burger"

	"github.com/gin-gonic/gin"
)

func OrderBurger(c *gin.Context) {
	type Input struct {
		BurgerType string `json:"burger_type"`
	}
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		return
	}
	// Create the appropriate builder based on the user's choice
	burgerBuilder := burger.BurgerFactory(input.BurgerType)
	if burgerBuilder == nil {
		fmt.Println("Invalid burger type")
		return
	}

	// Create a director with the selected builder
	burgerDirector := burger.NewBurgerDirector(burgerBuilder)

	// Construct and display the burger
	burger := burgerDirector.Construct()
	resp := gin.H{
		"burger_type": burger.GetPatty(),
	}
	helper.SendResponse(c, "success_order_burger", resp)
	return
}
