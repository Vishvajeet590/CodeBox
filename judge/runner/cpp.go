package runner

import (
	"CodeBox/models"
	"CodeBox/utils"
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func RunCPP(sourceFileName, input, output string, maxRunTime int, taskId string) (models.JudgementResult, error) {
	var resObj channelObj
	result := models.JudgementResult{
		ResultId: uuid.New().String(),
		TaskId:   taskId,
	}
	if !utils.IsFilePresent(sourceFileName) {
		result.ErrorInfo = "source file is not present or output name is empty"
		return result, fmt.Errorf("source file is not present or output name is empty")
	}

	//compiling
	var compileOutput bytes.Buffer
	var stderr bytes.Buffer
	compileId := uuid.New().String()
	compileCmd := exec.Command("g++", "-o", utils.CodeDIR+"/"+compileId, sourceFileName)
	compileCmd.Stdout = &compileOutput
	compileCmd.Stderr = &stderr
	err := compileCmd.Run()

	if err != nil {
		result.ErrorInfo = utils.RESULT_COMPILE_ERROR
		result.WrongLine = stderr.Bytes()
		return result, err
	}

	//Checking is Codebox.class name file is present at ./downloads/
	checkCount := 0
	for {
		if checkCount > 5 {
			result.ErrorInfo = "something happened at compilation %v file is absent"
			return result, fmt.Errorf("something happened at compilation %v file is absent", compileId)
		}
		if !utils.IsFilePresent(utils.CodeDIR + "/" + compileId) {
			time.Sleep(1000)
			//result.ErrorInfo = "something happened at compilation %v file is absent"
			//return result, fmt.Errorf("something happened at compilation %v file is absent", compileId)
		} else {
			break
		}
		checkCount++
	}

	objChan := make(chan channelObj)
	//errChan := make(chan error)

	go func() {
		var outputBytes bytes.Buffer
		var runStderr bytes.Buffer
		obj := channelObj{
			err: nil,
		}

		cmd := exec.Command(utils.CodeDIR + "/" + compileId)
		cmd.Stdin = bytes.NewBuffer([]byte(input))
		cmd.Stdout = &outputBytes
		cmd.Stderr = &runStderr

		startTime := time.Now()
		err = cmd.Run()
		endTime := time.Now()
		if err != nil {
			obj.err = err
			objChan <- obj
			//errChan <- err
		}

		//obj.cmd = cmd

		obj.t = endTime.Sub(startTime)
		obj.output = outputBytes
		obj.memory = cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss

		objChan <- obj

	}()
	select {
	case <-time.After(time.Duration(maxRunTime) * time.Second):
		result.ErrorInfo = "timed out"
		return result, fmt.Errorf("timed out")
	case resObj = <-objChan:
		if resObj.err != nil {
			result.ErrorInfo = resObj.err.Error()
			return result, err
		}
		close(objChan)
		//close(errChan)
	}

	cmpRes := bytes.Compare([]byte(output), resObj.output.Bytes())
	var resErr string
	if cmpRes == -1 || cmpRes == 1 {
		resErr = utils.RESULT_FAIL
	} else if cmpRes == 0 {
		resErr = utils.RESULT_PASS
	}

	result.Status = resErr
	result.RuntimeTime = resObj.t.Milliseconds()
	result.RuntimeMemory = resObj.memory
	result.ExpectedOutput = output
	result.LastOutput = resObj.output.String()

	//Deleting the files
	if err = os.Remove(sourceFileName); err != nil {
		return result, err
	}

	if err = os.Remove(utils.CodeDIR + "/" + compileId); err != nil {
		return result, err
	}

	return result, nil

}
