package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	School    string `json:"school"`
	Major     string `json:"major"`
}

var users = []User{
	{ID: 1, Firstname: "John", Lastname: "Doe", School: "Yuan Ze University", Major: "Computer Science"},
	{ID: 2, Firstname: "Jane", Lastname: "Smith", School: "New York College", Major: "Biology"},
}

func main() {

	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/api/users", getUsers)
	app.Get("/api/users/:id", getUser)

	log.Fatal(app.Listen(":5001"))
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
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
		"error": "Student not found",
	})
}
