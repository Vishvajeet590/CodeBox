package runner

import (
	"CodeBox/models"
	"CodeBox/repository/db"
	"CodeBox/utils"
	"bytes"
	"encoding/base64"
	"os"
	"time"
)

type Runner struct {
	Task models.TaskPackage
}

type channelObj struct {
	t         time.Duration
	output    bytes.Buffer
	errorType string
	err       error
	memory    int64
	//	errOutput bytes.Buffer
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(taskPack models.TaskPackage) error {

	//Create file of code
	r.Task = taskPack
	languageCode := utils.GetLanguageCode(r.Task.Language)
	fileName := utils.CodeDIR + "/" + r.Task.TaskId + utils.GetExtension(languageCode)

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	rawCode, err := base64.StdEncoding.DecodeString(r.Task.SourceCode)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(rawCode))
	if err != nil {
		return err
	}

	//TODO add running status for the task in DB later on

	//Running
	//var result models.JudgementResult
	switch languageCode {
	case utils.JAVA:
		result, err := RunJava(fileName, r.Task.Input, r.Task.Output, r.Task.TimeLimit, r.Task.TaskId)
		if err != nil {

			//TODO fix this shit code, DB transaction should happen at one point only
			_, err = db.CreateOne(result)
			if err != nil {
				return err
			}
			return err
		}
		_, err = db.UpdateOne(result)
		if err != nil {
			return err
		}

	case utils.CPP:
		result, err := RunCPP(fileName, r.Task.Input, r.Task.Output, r.Task.TimeLimit, r.Task.TaskId)
		if err != nil {
			_, err = db.CreateOne(result)
			if err != nil {
				return err
			}
			return err
		}
		_, err = db.UpdateOne(result)
		if err != nil {
			return err
		}

	case utils.PYTHON:
		result, err := RunPY(fileName, r.Task.Input, r.Task.Output, r.Task.TimeLimit, r.Task.TaskId)
		if err != nil {
			_, err = db.CreateOne(result)
			if err != nil {
				return err
			}
			return err
		}
		_, err = db.UpdateOne(result)
		if err != nil {
			return err
		}
	}

	return nil

}
