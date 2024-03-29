package controller

import (
	"go_practice/helper"
	singletondesignpattern "go_practice/service/singletonDesignPattern"

	"github.com/gin-gonic/gin"
)

func GetSingletonInstance(c *gin.Context) {

	singletonInstance := singletondesignpattern.GetSingletonInstance()

	resp := gin.H{
		"singleton_data": singletonInstance.Data,
	}
	helper.SendResponse(c, "singleton_instance_information", resp)
	return
}

func UpdateSingletonInstance(c *gin.Context) {
	type Input struct {
		Input string `json:"input"`
	}
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		return
	}
	singletonInstance := singletondesignpattern.GetSingletonInstance()
	singletonInstance.Data = input.Input
	resp := gin.H{
		"singleton_data": singletonInstance.Data,
	}
	helper.SendResponse(c, "update_singleton_instance_information", resp)
	return

}
