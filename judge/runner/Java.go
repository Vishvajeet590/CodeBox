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

func RunJava(sourceFileName, input, output string, maxRunTime int, taskId string) (models.JudgementResult, error) {

	var resObj channelObj
	result := models.JudgementResult{
		ResultId: uuid.New().String(),
		TaskId:   taskId,
	}

	if !utils.IsFilePresent(sourceFileName) {
		result.ErrorInfo = "source file is not present or output name is empty"
		return result, fmt.Errorf("source file is not present or output name is empty")
	}

	//Compiling file in a separate folder as class name will be same for every file

	//Creating a folder with task id
	compileFolder := utils.CodeDIR + "/" + taskId
	err := os.Mkdir(compileFolder, 0777)
	if err != nil {
		return result, err
	}

	var compileOutput bytes.Buffer
	var stderr bytes.Buffer
	compileCmd := exec.Command(utils.JAVA_COMPILE_CMD, "-d", compileFolder+"/", sourceFileName)
	compileCmd.Stdout = &compileOutput
	compileCmd.Stderr = &stderr
	err = compileCmd.Run()
	if err != nil {
		result.ErrorInfo = utils.RESULT_COMPILE_ERROR
		result.WrongLine = stderr.Bytes()
		return result, err
	}

	//Checking is Codebox.class name file is present at ./downloads/
	if !utils.IsFilePresent(compileFolder + "/" + utils.JAVA_CLASS_NAME) {
		result.ErrorInfo = "something happened at compilation %v file is absent"
		return result, fmt.Errorf("something happened at compilation %v file is absent", utils.JAVA_CLASS_NAME)
	}

	//running
	objChan := make(chan channelObj)
	go func() {
		var outputBytes bytes.Buffer
		var runStderr bytes.Buffer
		obj := channelObj{
			err: nil,
		}

		cmd := exec.Command(utils.JAVA_RUN_CMD, "-cp", compileFolder, utils.JAVA_RUN_FILE)
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

	//Cleanup
	if err = os.RemoveAll(compileFolder); err != nil {
		return result, err
	}
	if err = os.Remove(sourceFileName); err != nil {
		return result, err
	}

	return result, nil
}
