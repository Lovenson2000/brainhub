package main

import (
	"log"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=brainhub sslmode=disable password=Blatter@2000 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	app := fiber.New()

	// User routes
	app.Get("/api/users", func(c *fiber.Ctx) error {
		return controllers.GetUsers(db, c)
	})
	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		return controllers.GetUser(db, c)
	})
	app.Post("/api/users", func(c *fiber.Ctx) error {
		return controllers.CreateUser(db, c)
	})
	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteUser(db, c)
	})
	app.Patch("/api/users/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateUser(db, c)
	})

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
