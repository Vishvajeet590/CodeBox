package models

// TODO use these status to update state of compile ->  run -> success/fail/err/TLE
const (
	SE  = "SYSTEM_ERROR"
	WA  = "WRONG_ANSWER"
	AC  = "ACCEPTED"
	TLE = "TIME_LIMIT_EXCEED"
	OLE = "OUTPUT_LIMIT_EXCEED"
	MLE = "MEMORY_LIMIT_EXCEED"
	RE  = "RUNTIME_ERROR"
	PE  = "PRESENTATION_ERROR"
	CE  = "COMPILE_ERROR"
)
