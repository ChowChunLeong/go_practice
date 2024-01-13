package controller

import (
	"go_practice/helper"
	abstractfactorydesignpattern "go_practice/service/abstractFactoryDesignPattern"

	"github.com/gin-gonic/gin"
)

func UseFurniture(c *gin.Context) {
	type Input struct {
		Variant string `json:"variant"`
		Product string `json:"product"`
	}
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		return
	}
	furniture, _ := abstractfactorydesignpattern.CreateFurniture(input.Variant, input.Product)
	resp := gin.H{
		"furniture": furniture.SitOn(),
	}
	helper.SendResponse(c, "furniture", resp)

}
