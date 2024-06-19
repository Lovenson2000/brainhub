package model

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	School    string `json:"school"`
	Major     string `json:"major"`
	Bio       string `json:"bio"`
}
