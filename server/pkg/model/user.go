package model

type User struct {
	ID        int    `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	School    string `json:"school" db:"school"`
	Major     string `json:"major" db:"major"`
	Bio       string `json:"bio" db:"bio"`
	Posts     []Post `json:"posts,omitempty" db:"-"`
}
