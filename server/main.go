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
	{ID: 1, Firstname: "John", Lastname: "Doe", Email: "john.doe@example.com", School: "Yuan Ze University", Major: "Computer Science"},
	{ID: 2, Firstname: "Jane", Lastname: "Smith", Email: "jane.smith@example.com", School: "New York College", Major: "Biology"},
	{ID: 3, Firstname: "Alice", Lastname: "Johnson", Email: "alice.johnson@example.com", School: "Stanford University", Major: "Physics"},
	{ID: 4, Firstname: "Bob", Lastname: "Brown", Email: "bob.brown@example.com", School: "Harvard University", Major: "Law"},
	{ID: 5, Firstname: "Charlie", Lastname: "Davis", Email: "charlie.davis@example.com", School: "MIT", Major: "Engineering"},
	{ID: 6, Firstname: "Diana", Lastname: "Evans", Email: "diana.evans@example.com", School: "UC Berkeley", Major: "Mathematics"},
	{ID: 7, Firstname: "Edward", Lastname: "Garcia", Email: "edward.garcia@example.com", School: "University of Oxford", Major: "Literature"},
	{ID: 8, Firstname: "Fiona", Lastname: "Hill", Email: "fiona.hill@example.com", School: "Columbia University", Major: "History"},
	{ID: 9, Firstname: "George", Lastname: "Ingram", Email: "george.ingram@example.com", School: "University of Cambridge", Major: "Economics"},
	{ID: 10, Firstname: "Hannah", Lastname: "Jones", Email: "hannah.jones@example.com", School: "Yale University", Major: "Political Science"},
}

func main() {

	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/api/users", getUsers)
	app.Get("/api/users/:id", getUser)
	app.Post("/api/users", createUser)
	app.Delete("/api/users/:id", deleteUser)
	app.Patch("/api/users/:id", updateUser)

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

func createUser(c *fiber.Ctx) error {
	newUser := new(User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newUser.ID = len(users) + 1
	users = append(users, *newUser)

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func deleteUser(c *fiber.Ctx) error {
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

func updateUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	updatedUser := new(User)
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
