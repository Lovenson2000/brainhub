package model

type Document struct {
	ID             int    `json:"id" db:"id"`
	StudySessionID int    `json:"study_session_id" db:"study_session_id"`
	Title          string `json:"title" db:"title"`
	URL            string `json:"url" db:"url"`
	UploadedBy     int    `json:"uploaded_by" db:"uploaded_by"`
	UploadedAt     string `json:"uploaded_at" db:"uploaded_at"`
}
