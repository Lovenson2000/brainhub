package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
)

type Post = model.Post
type Comment = model.Comment
type StudySession = model.StudySession

var posts = []Post{
	{ID: 1, UserID: 1, Title: "How to learn Go?", Content: "I'm new to Go. Any tips?"},
	{ID: 2, UserID: 2, Title: "Favorite Biology Resources", Content: "What are your favorite resources for learning biology?"},
	{ID: 3, UserID: 3, Title: "Is TailwindCSS better than vanilla css ?", Content: "What are your thoughts regarding the comparisons between TailwindCSS and Bootstrap ?"},
}

var comments = []Comment{
	{ID: 1, UserID: 2, PostID: 1, Text: "Check out the Go documentation."},
	{ID: 2, UserID: 1, PostID: 2, Text: "I like Khan Academy for biology."},
	{ID: 3, UserID: 3, PostID: 3, Text: "For me, Tailwind is better faster."},
}

var studySessions = []StudySession{
	{ID: 1, Title: "Go Programming Basics", Description: "Introduction to Go programming", StartTime: time.Now(), EndTime: time.Now().Add(2 * time.Hour), Location: "Library Room 101", Participants: []int{1, 2}},
	{ID: 2, Title: "Biology Study Group", Description: "Study group for biology majors", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour), Location: "Biology Lab", Participants: []int{2}},
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	// User routes
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	app.Patch("/api/users/:id", controllers.UpdateUser)

	// Post routes
	app.Get("/api/posts", getPosts)
	app.Get("/api/posts/:id", getPost)
	app.Post("/api/posts", createPost)
	app.Delete("/api/posts/:id", deletePost)
	app.Patch("/api/posts/:id", updatePost)

	// Comment routes
	app.Get("/api/comments", getAllComments)
	app.Get("/api/posts/:postId/comments", getPostComments)
	app.Post("/api/comments", createComment)
	// app.Delete("/api/comments/:id", deleteComment)
	// app.Patch("/api/comments/:id", updateComment)

	// // StudySession routes
	// app.Get("/api/study-sessions", getStudySessions)
	// app.Get("/api/study-sessions/:id", getStudySession)
	// app.Post("/api/study-sessions", createStudySession)
	// app.Delete("/api/study-sessions/:id", deleteStudySession)
	// app.Patch("/api/study-sessions/:id", updateStudySession)

	log.Fatal(app.Listen(":5001"))
}

// POST HANDLERS

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

func createPost(c *fiber.Ctx) error {
	newPost := new(Post)

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newPost.ID = len(posts) + 1
	posts = append(posts, *newPost)

	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func deletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return c.JSON(fiber.Map{"message": "Post deleted"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}

func updatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	updatedPost := new(Post)
	if err := c.BodyParser(updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i].UserID = updatedPost.UserID
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			return c.JSON(posts[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Post not found",
	})
}

func getAllComments(c *fiber.Ctx) error {
	return c.JSON(comments)
}

func getPostComments(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("postId"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Post ID",
		})
	}

	var postComments []Comment

	for _, comment := range comments {
		if comment.PostID == postId {
			postComments = append(postComments, comment)
		}
	}

	if len(postComments) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No comments found for the post",
		})
	}

	return c.JSON(postComments)
}

func createComment(c *fiber.Ctx) error {
	newComment := new(Comment)

	if err := c.BodyParser(newComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing Comment body",
		})
	}

	newComment.ID = len(comments) + 1
	comments = append(comments, *newComment)

	return c.Status(fiber.StatusCreated).JSON(newComment)
}
