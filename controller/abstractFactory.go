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
	furnitureChair, furnitureTable := abstractfactorydesignpattern.CreateFurniture(input.Variant, input.Product)
	resp := gin.H{}
	if furnitureChair == nil && furnitureTable == nil {
		resp = gin.H{
			"furniture": "no_matched",
		}
	}

	if furnitureChair != nil {
		resp = gin.H{
			"furniture_chair": furnitureChair.SitOn(),
		}
	}

	if furnitureTable != nil {
		resp = gin.H{
			"furniture_table": furnitureTable.PutOn(),
		}
	}

	helper.SendResponse(c, "furniture", resp)

}
