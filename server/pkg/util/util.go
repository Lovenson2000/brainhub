package util

import (
	"github.com/jmoiron/sqlx"
)

func UserExists(db *sqlx.DB, userID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)"
	err := db.Get(&exists, query, userID)
	return exists, err
}
