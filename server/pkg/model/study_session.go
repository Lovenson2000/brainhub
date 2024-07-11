package model

import (
	"github.com/Lovenson2000/brainhub/pkg/model"
)

type StudySession struct {
	ID           int             `json:"id" db:"id"`
	Title        string          `json:"title" db:"title"`
	Description  string          `json:"description" db:"description"`
	StartTime    string          `json:"start_time" db:"start_time"`
	EndTime      string          `json:"end_time" db:"end_time"`
	Location     string          `json:"location" db:"location"`
	Participants []int           `json:"participants" db:"-"`
	Documents    []Document      `json:"documents,omitempty" db:"-"`
	ChatMessages []model.Message `json:"chat_messages,omitempty" db:"-"`
}
