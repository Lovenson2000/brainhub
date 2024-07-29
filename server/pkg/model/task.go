package model

type Status string

const (
	ToDo       Status = "To Do"
	InProgress Status = "In Progress"
	Done       Status = "Done"
)

type Priority string

const (
	Low    Priority = "Low"
	Normal Priority = "Normal"
	High   Priority = "High"
)

type Task struct {
	ID          string   `json:"id"`
	UserID      int      `json:"user_id" db:"user_id"`
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	StartTime   string   `json:"start_time" db:"start_time"`
	DueDate     string   `json:"due_date" db:"due_date"`
	Status      Status   `json:"status" db:"status"`
	Priority    Priority `json:"priority" db:"priority"`
}
