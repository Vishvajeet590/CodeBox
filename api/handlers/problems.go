package handlers

import (
	"CodeBox/api/services"
	"CodeBox/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProblem(c *gin.Context) {
	res, err := services.AddProblem(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetProblemById(c *gin.Context) {
	res, err := services.GetProblemById(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetProblems(c *gin.Context) {
	res, err := services.GetProblems(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
