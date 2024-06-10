package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Student struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	School    string `json:"school"`
	Major     string `json:"major"`
}

var students = []Student{
	{ID: 1, Firstname: "John", Lastname: "Doe", School: "Yuan Ze University", Major: "Computer Science"},
	{ID: 2, Firstname: "Jane", Lastname: "Smith", School: "New York College", Major: "Biology"},
}

func main() {

	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/api/students", getStudents)
	app.Get("/api/students/:id", getStudent)

	log.Fatal(app.Listen(":5001"))
}

func getStudents(c *fiber.Ctx) error {
	return c.JSON(students)
}

func getStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	for _, student := range students {
		if student.ID == id {
			return c.JSON(student)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Student not found",
	})
}
