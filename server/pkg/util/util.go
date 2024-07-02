package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func UserExists(db *sqlx.DB, userID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)"
	err := db.Get(&exists, query, userID)
	return exists, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(storedPassword, newPassWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(newPassWord))
	return err == nil
}

func CreateJwtToken(userID int, secretKey string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
