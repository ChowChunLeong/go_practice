package controller

import (
	"fmt"
	"go_practice/database"
	"go_practice/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// Read client ID from query parameter
	clientID := c.Query("clientId")
	if clientID == "" {
		fmt.Println("Client ID is required.")
		return
	}
	fmt.Println("clientId", clientID)
	database.RedisDb.SetKeyValue("clientId", "redis_value")

	for {
		testingRedisValue, _ := database.RedisDb.GetKeyValue("clientId")
		msg := "redis_value"
		if testingRedisValue != "redis_value" {
			msg = testingRedisValue
		}
		msgBytes := []byte(msg)
		conn.WriteMessage(websocket.TextMessage, msgBytes)
		fmt.Println("ws")
	}
}

func ChangeRedisClientId(c *gin.Context) {
	type Input struct {
		ClientId string `json:"client_id"`
	}
	ori, err := database.RedisDb.GetKeyValue("clientId")
	fmt.Println("err", err)
	fmt.Println("ori", ori)
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		return
	}
	database.RedisDb.SetKeyValue("clientId", input.ClientId)
	helper.SendResponse(c, "send_ws_msg", helper.EmptyObj{})
	return
}
