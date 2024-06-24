package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// GET ALL USERS
func GetUsers(db *sqlx.DB, c *fiber.Ctx) error {
	var users []model.User
	err := db.Select(&users, "SELECT id, firstname, lastname, email, school, major, bio FROM users ORDER BY id ASC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return c.JSON(users)
}

// GET ONE USER
func GetUser(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var user model.User
	err = db.Get(&user, "SELECT id, firstname, lastname, email, school, major, bio FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

// CREATE A USER
func CreateUser(db *sqlx.DB, c *fiber.Ctx) error {
	newUser := new(model.User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var numberOfUsers int
	err := db.Get(&numberOfUsers, "SELECT COUNT(*) from users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count users",
		})
	}

	newUser.ID = numberOfUsers + 1

	query := `INSERT INTO users (id, firstname, lastname, email, school, major, bio) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(query, newUser.ID, newUser.Firstname, newUser.Lastname, newUser.Email, newUser.School, newUser.Major, newUser.Bio)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

// DELETE A USER
func DeleteUser(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{"message": "User deleted"})
}

// UPDATE A USER
func UpdateUser(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	updatedUser := new(model.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	query := `UPDATE users SET firstname=$1, lastname=$2, email=$3, school=$4, major=$5, bio=$6 WHERE id=$7`
	_, err = db.Exec(query, updatedUser.Firstname, updatedUser.Lastname, updatedUser.Email, updatedUser.School, updatedUser.Major, updatedUser.Bio, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(updatedUser)
}
