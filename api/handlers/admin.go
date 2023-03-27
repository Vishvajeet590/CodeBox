package handlers

import (
	"CodeBox/api/services"
	"CodeBox/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminSignIn(c *gin.Context) {
	res, err := services.AdminSignIn(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AdminSignUp(c *gin.Context) {
	res, err := services.CreateAdmin(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
