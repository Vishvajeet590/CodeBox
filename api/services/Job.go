package services

import (
	"CodeBox/models"
	"CodeBox/repository/db"
	"CodeBox/repository/rabbitMQ"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"os"
)

func CreateJob(c *gin.Context) (interface{}, error) {
	var task models.JudgementTask
	c.BindJSON(&task)
	task.TaskId = uuid.New().String()
	if err := validator.New().Struct(task); err != nil {
		return nil, err
	}

	//Get the problem

	//Sandbox req
	if task.ProblemId == "sandbox-101" {
		item := models.Problem{}
		problems, err := db.GetAll(&item, 1, 100, "", "")
		if err != nil {
			return nil, err
		}

		for _, x := range problems.Results {
			task.ProblemId = x.ProblemId
			break
		}
	}

	problem := models.Problem{
		ProblemId: task.ProblemId,
	}
	problem, err := db.GetOne(problem)
	if err != nil {
		return nil, err
	}

	res, err := db.CreateOne(&task)
	if err != nil {
		return nil, err
	}

	pack := models.TaskPackage{
		TaskId:      task.TaskId,
		Language:    task.Language,
		TimeLimit:   task.TimeLimit,
		MemoryLimit: task.MemoryLimit,
		SourceCode:  task.SourceCode,
		Input:       problem.Input,
		Output:      problem.ExpectedOutput,
	}

	fmt.Printf("%v", pack)
	jsonData, err := json.Marshal(pack)
	if err != nil {
		return nil, err
	}
	fmt.Printf("JSONDATA = \n %s", string(jsonData))
	err = rabbitMQ.RmqEmitter.Push(string(jsonData), os.Getenv("RMQ_QUEUE_NAME"))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetJobResult(c *gin.Context) (*models.JudgementResult, error) {
	taskId := c.Param("id")
	if taskId == "" {
		return nil, fmt.Errorf("invalid task Id")
	}

	taskResult := models.JudgementResult{
		TaskId: taskId,
	}

	task, err := db.GetOne(taskResult)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
