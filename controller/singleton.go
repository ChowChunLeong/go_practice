package controller

import (
	"database/sql"
	"fmt"
	"go_practice/helper"
	"go_practice/service"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetSingletonInstance(c *gin.Context) {
	cnn, err := sql.Open("mysql", "root:root@tcp(172.17.0.3:3306)/live_sp_admin")
	if err != nil {
		log.Fatal(err)
	}
	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM users WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	singletonInstance := service.GetSingletonInstance()

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
	singletonInstance := service.GetSingletonInstance()
	singletonInstance.Data = input.Input
	resp := gin.H{
		"singleton_data": singletonInstance.Data,
	}
	helper.SendResponse(c, "update_singleton_instance_information", resp)
	return

}
