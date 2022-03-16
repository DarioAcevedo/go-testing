package utils

import "github.com/gin-gonic/gin"


func Respond(c *gin.Context, status int, data interface{}) {
	accept_string := c.GetHeader("Accept")
	switch accept_string {
		case "application/xml":
			c.XML(status, data)
		default:
			c.JSON(status, data)
	}

}