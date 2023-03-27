package models

import (
	"github.com/lib/pq"
	"time"
)

type Problem struct {
	ProblemId         string         `gorm:"column:problem_id;primary_key" json:"question_id"`
	Title             string         `gorm:"column:title;NOT NULL" json:"title"`
	QuestionCreatedAt time.Time      `gorm:"column:question_created_at;NOT NULL" json:"question_created_at"`
	ProblemStatement  string         `gorm:"column:problem_statement" json:"problem_statement"`
	Input             string         `gorm:"column:input" json:"input"`
	ExpectedOutput    string         `gorm:"column:output" json:"output"`
	InputStatement    string         `gorm:"column:input_statement" json:"input_statement"`
	OutputStatement   string         `gorm:"column:output_statement" json:"output_statement"`
	Level             string         `gorm:"column:level" json:"level"`
	Constraints       pq.StringArray `gorm:"type:text[]" json:"constraints"`
	SampleInputs      pq.StringArray `gorm:"type:text[]" json:"sample_inputs"`
	SampleOutputs     pq.StringArray `gorm:"type:text[]" json:"sample_outputs"`
}
