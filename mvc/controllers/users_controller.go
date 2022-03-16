package controllers

import (
	"net/http"
	"strconv"

	"github.com/DarioAcevedo/go-testing/mvc/services"
	"github.com/DarioAcevedo/go-testing/mvc/utils"
	"github.com/gin-gonic/gin"
)

// Returns a user by id. If the user do not exist, it returns an error depending on conditions
func GetUser(c *gin.Context) {
	userIdParam  := c.Param("user_id")
	userId, error := strconv.ParseInt(userIdParam, 10, 64)
	if error != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	User, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, User)
}