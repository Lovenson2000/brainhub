package main

import (
	"log"
	"os"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connectionString := "user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable password=" + os.Getenv("DB_PASSWORD") +
		" host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT")

	db, err := sqlx.Connect("postgres", connectionString)
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
		return controllers.GetUserWithPosts(db, c)
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
	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return controllers.GetPosts(db, c)
	})

	app.Get("/api/posts/:id", func(c *fiber.Ctx) error {
		return controllers.GetPost(db, c)
	})

	app.Post("/api/posts", func(c *fiber.Ctx) error {
		return controllers.CreatePost(db, c)
	})

	app.Delete("/api/posts/:id", func(c *fiber.Ctx) error {
		return controllers.DeletePost(db, c)
	})

	app.Patch("/api/posts/:id", func(c *fiber.Ctx) error {
		return controllers.UpdatePost(db, c)
	})

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
