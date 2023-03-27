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

func RunPY(sourceFileName, input, output string, maxRunTime int, taskId string) (models.JudgementResult, error) {
	var resObj channelObj
	result := models.JudgementResult{
		ResultId: uuid.New().String(),
		TaskId:   taskId,
	}
	if !utils.IsFilePresent(sourceFileName) {
		result.ErrorInfo = "source file is not present or output name is empty"
		return result, fmt.Errorf("source file is not present or output name is empty")
	}

	objChan := make(chan channelObj)
	//initializing error
	var err error
	//errChan := make(chan error)

	go func() {
		var outputBytes bytes.Buffer
		var runStderr bytes.Buffer
		obj := channelObj{
			err: nil,
		}

		cmd := exec.Command(utils.PYTHON_RUN_CMD, sourceFileName)
		cmd.Stdin = bytes.NewBuffer([]byte(input))
		cmd.Stdout = &outputBytes
		cmd.Stderr = &runStderr

		startTime := time.Now()
		err := cmd.Run()
		endTime := time.Now()
		if err != nil {
			obj.err = fmt.Errorf(runStderr.String() + "\n" + err.Error())
			objChan <- obj
			//errChan <- err
		}

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

	return result, nil

}
