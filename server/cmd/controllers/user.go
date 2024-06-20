package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
)

var users = []model.User{
	{ID: 1, Firstname: "John", Lastname: "Doe", Email: "john.doe@example.com", School: "Yuan Ze University", Major: "Computer Science"},
	{ID: 2, Firstname: "Jane", Lastname: "Smith", Email: "jane.smith@example.com", School: "New York College", Major: "Biology"},
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}

func CreateUser(c *fiber.Ctx) error {
	newUser := new(model.User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newUser.ID = len(users) + 1
	users = append(users, *newUser)

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func UpdateUser(c *fiber.Ctx) error {
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

	for i, user := range users {
		if user.ID == id {
			users[i].Firstname = updatedUser.Firstname
			users[i].Lastname = updatedUser.Lastname
			users[i].Email = updatedUser.Email
			users[i].School = updatedUser.School
			users[i].Major = updatedUser.Major
			return c.JSON(users[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(fiber.Map{"message": "User deleted"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}
