package models

import (
	"encoding/json"
	"time"
)

type JudgementTask struct {
	TaskId      string `json:"id" gorm:"primaryKey" gorm:"column:task_id" validate:"required"`
	ProblemId   string `json:"problem_Id" gorm:"column:problem_id" validate:"required"`
	UserId      string `json:"user_id" gorm:"column:user_id" validate:"required"`
	Language    string `json:"language" gorm:"column:language" validate:"required"`
	TimeLimit   int    `json:"time_limit" gorm:"column:time_limit" validate:"required"`
	MemoryLimit int    `json:"memory_limit" gorm:"column:memory_limit" validate:"required"`
	SourceCode  string `json:"SourceCode" gorm:"column:source_code" validate:"required"`
}

type JudgementResult struct {
	ResultId      string `json:"result_id" gorm:"primaryKey" gorm:"column:result_id" validate:"required"`
	TaskId        string `json:"task_id"  gorm:"column:task_id" validate:"required"`
	Succeed       bool   `json:"succeed" gorm:"column:succeed" validate:"required"`
	Status        string `json:"status" gorm:"column:status" validate:"required"`
	RuntimeTime   int64  `json:"runtime_time" gorm:"column:runtime" validate:"required"`
	RuntimeMemory int64  `json:"runtime_memory" gorm:"column:runtime_memory" validate:"required"`

	WrongLine      []byte `json:"wrong_line" gorm:"column:wrong_line"`
	LastInput      string `json:"last_input" gorm:"column:last_input"`
	LastOutput     string `json:"last_output" gorm:"column:last_output"`
	ExpectedOutput string `json:"expected_output"  gorm:"column:expected_output"`
	ErrorInfo      string `json:"error_info"  gorm:"column:extra_info"`

	Timestamp time.Time `json:"timestamp"  gorm:"column:timestamp"`
}

type TaskPackage struct {
	TaskId      string `json:"id" gorm:"primaryKey" gorm:"column:task_id" validate:"required"`
	Language    string `json:"language" gorm:"column:language" validate:"required"`
	TimeLimit   int    `json:"time_limit" gorm:"column:time_limit" validate:"required"`
	MemoryLimit int    `json:"memory_limit" gorm:"column:memory_limit" validate:"required"`
	SourceCode  string `json:"SourceCode" gorm:"column:source_code" validate:"required"`
	Input       string `gorm:"column:input" json:"input"`
	Output      string `gorm:"column:output" json:"output"`
}

func (result *JudgementResult) ToJsonString() []byte {
	data, _ := json.Marshal(result)
	return data
}
