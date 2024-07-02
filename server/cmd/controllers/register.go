package controllers

import (
	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterUser(db *sqlx.DB, c *fiber.Ctx) error {

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	var existingUser model.User
	err := db.Get(&existingUser, "SELECT id FROM users WHERE email=$1", user.Email)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "email already in use",
		})
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot hash password",
		})
	}

	user.Password = hashedPassword

	query := `INSERT INTO users(firstname, lastname, email, password, school, major, bio) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(query, user.Firstname, user.Lastname, user.Email, user.Password, user.School, user.Major, user.Bio)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create user"})
	}

	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(user)

}
