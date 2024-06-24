package model

type Post struct {
	ID      int    `json:"id" db:"id"`
	UserID  int    `json:"user_id" db:"user_id"`
	Content string `json:"content" db:"content"`
	Image   string `json:"image" db:"image"`
}
