package util

import (
	"fmt"
	"os"
	"time"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

func CheckPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func CreateJwtToken(userID int, secretKey string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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

func InsertPostIntoDB(db *sqlx.DB, userID int, content string, imageUrl string) (*model.Post, error) {

	newPost := &model.Post{
		UserID:  userID,
		Content: content,
		Image:   imageUrl,
	}

	query := `INSERT INTO posts (user_id, content, image) VALUES($1, $2, $3) RETURNING id, created_at`
	err := db.QueryRow(query, newPost.UserID, newPost.Content, newPost.Image).Scan(&newPost.ID, &newPost.CreatedAt)
	if err != nil {
		return nil, err
	}

	return newPost, nil
}

func GetPostByID(db *sqlx.DB, id int) (*model.Post, error) {

	var post model.Post
	err := db.Get(&post, "SELECT id, user_id, content, image, created_at FROM posts WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func CreateTables(db *sqlx.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			firstname VARCHAR(50),
			lastname VARCHAR(50),
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			school VARCHAR(100),
			major VARCHAR(100),
			bio TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id) ON DELETE CASCADE,
			content TEXT,
			image TEXT,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS study_sessions (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			title VARCHAR(100) NOT NULL,
			description TEXT,
			start_time TIMESTAMP,
			end_time TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS documents (
			id SERIAL PRIMARY KEY,
			study_session_id INTEGER REFERENCES study_sessions(id) ON DELETE CASCADE,
			title VARCHAR(100),
			url VARCHAR(255) NOT NULL,
			uploaded_by INTEGER REFERENCES users(id) ON DELETE CASCADE,
			uploaded_at TIMESTAMP NOT NULL
		)`,
		`CREATE TYPE task_status AS ENUM ('To Do', 'In Progress', 'Done');`,
		`CREATE TYPE task_priority AS ENUM ('Low', 'Normal', 'High');`,
		`CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			title VARCHAR(100) NOT NULL,
			description TEXT,
			start_time TIMESTAMP,
			due_date TIMESTAMP,
			priority task_priority DEFAULT 'Normal',
			status task_status DEFAULT 'To Do',
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExtractUserIDFromJwtToken(c *fiber.Ctx) (int, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Missing authorization header")
	}

	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	if !token.Valid {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}

	userIDFloat64, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID in token")
	}
	userID := int(userIDFloat64)

	return userID, nil
}
