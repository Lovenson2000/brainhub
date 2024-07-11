package model

type Message struct {
	ID             int    `json:"id" db:"id"`
	StudySessionID int    `json:"study_session_id" db:"study_session_id"`
	UserID         int    `json:"user_id" db:"user_id"`
	Message        string `json:"message" db:"message"`
	CreatedAt      string `json:"created_at" db:"created_at"`
}
