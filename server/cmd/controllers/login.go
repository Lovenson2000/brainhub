package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func LoginUser(db *sqlx.DB, c *fiber.Ctx) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the secret key from environment variables
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("jwt Secret key not found in .env file")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "jwt Secret key not found",
		})
	}

	var req model.User
	var user model.User

	// STEP 1 Parse the JSON request body into req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// STEP 2 Check if the user exists in the database
	query := "SELECT id, firstname, lastname, email, password, school, major, bio FROM users WHERE email=$1"
	err = db.Get(&user, query, req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	// STEP 3 Compare the stored hashed password with the provided password
	if !util.CheckPasswordHash(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	// STEP 3 Create JWT token
	token, err := util.CreateJwtToken(user.ID, jwtSecretKey)
	fmt.Println(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot generate token",
		})
	}

	// STEP 4 Send token and user as response
	return c.JSON(fiber.Map{
		"Token": token,
		"user":  user,
	})
}
