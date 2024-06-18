package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	School    string `json:"school"`
	Major     string `json:"major"`
	Bio       string `json:"bio"`
}

type Comment struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	PostID int    `json:"post_id"`
	Text   string `json:"text"`
}

type Post struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
}

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Text       string `json:"text"`
}

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Event struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type StudySession struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Location     string    `json:"location"`
	Participants []int     `json:"participants"` // List of user IDs
}

// TODO: REMOVE WHEN CONNECTED TO POSTGRESQL DB
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

var posts = []Post{
	{ID: 1, UserID: 1, Title: "How to learn Go?", Content: "I'm new to Go. Any tips?"},
	{ID: 2, UserID: 2, Title: "Favorite Biology Resources", Content: "What are your favorite resources for learning biology?"},
	// Add more posts as needed
}

var comments = []Comment{
	{ID: 1, UserID: 2, PostID: 1, Text: "Check out the Go documentation."},
	{ID: 2, UserID: 1, PostID: 2, Text: "I like Khan Academy for biology."},
	// Add more comments as needed
}

var studySessions = []StudySession{
	{ID: 1, Title: "Go Programming Basics", Description: "Introduction to Go programming", StartTime: time.Now(), EndTime: time.Now().Add(2 * time.Hour), Location: "Library Room 101", Participants: []int{1, 2}},
	{ID: 2, Title: "Biology Study Group", Description: "Study group for biology majors", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour), Location: "Biology Lab", Participants: []int{2}},
}

func main() {

	fmt.Println("Hello World")

	app := fiber.New()

	// User routes
	app.Get("/api/users", getUsers)
	app.Get("/api/users/:id", getUser)
	app.Post("/api/users", createUser)
	app.Delete("/api/users/:id", deleteUser)
	app.Patch("/api/users/:id", updateUser)

	// Post routes
	app.Get("/api/posts", getPosts)
	app.Get("/api/posts/:id", getPost)
	// app.Post("/api/posts", createPost)
	// app.Delete("/api/posts/:id", deletePost)
	// app.Patch("/api/posts/:id", updatePost)

	log.Fatal(app.Listen(":5001"))
}

// USERS HANLDERS
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

// POST HANLDERS

func getPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}

func getPost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	for _, post := range posts {
		if post.ID == id {
			return c.JSON(post)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}
