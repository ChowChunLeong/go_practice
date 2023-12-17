package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmptyObj struct{}

//Response is used for static shape json return
type Response struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

/* Use; Build response with Object provided and Invoke Response function for response c (gin.Context) */ //
/* On Success; Response{ status: "200", title, object } */ //
/* On Error; {{none}} */ //
func SendResponse(c *gin.Context, title string, obj interface{}) {
	response := BuildResponse(http.StatusOK, title, obj)
	c.JSON(http.StatusOK, response)
}
