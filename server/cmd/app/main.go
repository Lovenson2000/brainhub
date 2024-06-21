package main

import (
	"fmt"
	"log"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/gofiber/fiber/v2"
)

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
	app.Get("/api/posts", controllers.GetPosts)
	app.Get("/api/posts/:id", controllers.GetPost)
	app.Post("/api/posts", controllers.CreatePost)
	app.Delete("/api/posts/:id", controllers.DeletePost)
	app.Patch("/api/posts/:id", controllers.UpdatePost)

	// Comment routes
	app.Get("/api/comments", controllers.GetAllComments)
	app.Get("/api/posts/:postId/comments", controllers.GetPostComments)
	app.Post("/api/comments", controllers.CreateComment)
	app.Delete("/api/comments/:id", controllers.DeleteComment)
	app.Patch("/api/comments/:id", controllers.UpdateComment)

	// // StudySession routes
	app.Get("/api/study-sessions", controllers.GetStudySessions)
	app.Get("/api/study-sessions/:id", controllers.GetStudySession)
	app.Post("/api/study-sessions", controllers.CreateStudySession)
	app.Delete("/api/study-sessions/:id", controllers.DeleteStudySession)
	app.Patch("/api/study-sessions/:id", controllers.UpdateStudySession)

	log.Fatal(app.Listen(":5001"))
}
