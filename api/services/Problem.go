package services

import (
	"CodeBox/models"
	"CodeBox/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddProblem(c *gin.Context) (interface{}, error) {
	var item models.Problem
	err := c.BindJSON(&item)
	if err != nil {
		return nil, err
	}
	item.ProblemId = uuid.New().String()
	if err := validator.New().Struct(item); err != nil {
		return nil, err
	}
	res, err := db.CreateOne(&item)
	return res, err
}

func GetProblems(c *gin.Context) (interface{}, error) {
	item := models.Problem{}
	res, err := db.GetAll(&item, 1, 100, "", "")
	for _, x := range res.Results {
		x.ExpectedOutput = ""
	}
	return res, err
}

func GetProblemById(c *gin.Context) (interface{}, error) {
	Id := c.Param("id")
	if Id == "" {
		return nil, fmt.Errorf("invalid problem Id")
	}
	item := models.Problem{
		ProblemId: Id,
	}
	res, err := db.GetOne(item)
	return res, err
}
