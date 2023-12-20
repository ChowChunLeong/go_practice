package controller

import (
	"go_practice/helper"
	"go_practice/service"

	"github.com/gin-gonic/gin"
)

func DrawShape(c *gin.Context) {
	type Input struct {
		ShapeType string `json:"shape_type"`
	}
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		return
	}
	shape := service.GetShape(input.ShapeType)
	resp := gin.H{
		"draw_shape": shape.Draw(),
	}
	helper.SendResponse(c, "success_draw_shape", resp)
	return
}
