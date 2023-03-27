package utils

import "net/http"

func FormatResponse(d interface{}, success bool) map[string]interface{} {
	if success {
		return map[string]interface{}{"internal_response_code": 0, "data": d, "message": "Success"}
	}
	return map[string]interface{}{"internal_response_code": 1, "data": d, "message": "Failed"}
}

func FormatResponseMessage(d interface{}, err error, successStatusCode int) (int, map[string]interface{}) {
	if err != nil {
		data := FormatResponse(d, false)
		return http.StatusInternalServerError, data
	}
	data := FormatResponse(d, true)
	return successStatusCode, data
}
