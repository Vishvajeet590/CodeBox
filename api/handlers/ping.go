package handlers

import (
	"CodeBox/api/services"
	"CodeBox/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	res, err := services.Ping(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
