package model

type Comment struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	PostID int    `json:"post_id"`
	Text   string `json:"text"`
}
