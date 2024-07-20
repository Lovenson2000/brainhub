package model

import "time"

type Note struct {
	ID             int       `json:"id" db:"id"`
	StudySessionID int       `json:"study_session_id" db:"study_session_id"`
	UserID         int       `json:"user_id" db:"user_id"`
	Content        string    `json:"content" db:"content"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"uploaded_at" db:"uploaded_at"`
}
