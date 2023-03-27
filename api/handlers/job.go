package handlers

import (
	"CodeBox/api/services"
	"CodeBox/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJob(c *gin.Context) {
	res, err := services.CreateJob(c)
	//fmt.Printf("Validator Err = %v", err.Error())
	if err != nil {
		fmt.Printf("%f", err)
	}
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetJobResult(c *gin.Context) {
	res, err := services.GetJobResult(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
