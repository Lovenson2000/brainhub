package main

import (
	"log"
	"os"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/Lovenson2000/brainhub/cmd/middleware"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	ngrokURL := os.Getenv("NGROK_URL")
	if ngrokURL == "" {
		log.Println("NGROK_URL is not set")
	} else {
		log.Printf("Using ngrok URL: %s\n", ngrokURL)
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS study_sessions, users CASCADE")
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	err = util.CreateTables(db)
	if err != nil {
		log.Fatalf("Failed to create database tables: %v\n", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	app := fiber.New()

	//Custom middlewares
	app.Use(middleware.CorsHandler())
	app.Use(middleware.Logger())
	app.Use(middleware.ErrorHandler())
	app.Use(middleware.Compresssor())
	app.Use(middleware.RateLimiter())

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

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return controllers.RegisterUser(db, c)
	})
	app.Post("/api/login", func(c *fiber.Ctx) error {
		return controllers.LoginUser(db, c)
	})

	app.Get("/api/study-sessions", func(c *fiber.Ctx) error {
		return controllers.GetStudySessions(db, c)
	})

	// Tasks routes (For testing purpose)
	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		return controllers.GetAllTasks(db, c)
	})
	app.Get("/api/users/:userID/tasks", func(c *fiber.Ctx) error {
		return controllers.GetTasksByUserID(db, c)
	})

	app.Use(middleware.Auth())

	// Post routes
	posts := app.Group("api/posts")
	posts.Use(middleware.Auth())

	posts.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetPosts(db, c)
	})
	posts.Get("/:id", func(c *fiber.Ctx) error {
		return controllers.GetPost(db, c)
	})
	posts.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreatePost(db, c)
	})
	posts.Delete("/:id", func(c *fiber.Ctx) error {
		return controllers.DeletePost(db, c)
	})
	posts.Patch("/:id", func(c *fiber.Ctx) error {
		return controllers.UpdatePost(db, c)
	})

	// Comment routes
	app.Get("/api/comments", controllers.GetAllComments)
	app.Get("/api/posts/:postId/comments", controllers.GetPostComments)
	app.Post("/api/comments", controllers.CreateComment)
	app.Delete("/api/comments/:id", controllers.DeleteComment)
	app.Patch("/api/comments/:id", controllers.UpdateComment)

	// StudySession routes (with Auth middleware)
	studySessions := app.Group("api/study-sessions")
	studySessions.Use(middleware.Auth())

	studySessions.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetStudySessionsByUserId(db, c)
	})
	studySessions.Get("/:id", func(c *fiber.Ctx) error {
		return controllers.GetStudySession(db, c)
	})
	studySessions.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateStudySession(db, c)
	})
	studySessions.Delete("/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteStudySession(db, c)
	})
	studySessions.Patch("/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateStudySession(db, c)
	})

	// Document routes
	studySessions.Post("/:id/documents", func(c *fiber.Ctx) error {
		return controllers.UploadDocument(db, c)
	})
	studySessions.Get("/:id/documents", func(c *fiber.Ctx) error {
		return controllers.GetDocumentsByStudySession(db, c)
	})
	studySessions.Delete("/:id/documents/:docId", func(c *fiber.Ctx) error {
		return controllers.DeleteDocument(db, c)
	})

	app.Get("/api/documents", func(c *fiber.Ctx) error {
		return controllers.GetAllDocuments(db, c)
	})
	app.Get("/api/users/:userID/documents", func(c *fiber.Ctx) error {
		return controllers.GetDocumentsByUserID(db, c)
	})

	tasks := app.Group("/api/tasks")
	tasks.Get("/:id", func(c *fiber.Ctx) error {
		return controllers.GetTask(db, c)
	})
	tasks.Get("/user/:userID", func(c *fiber.Ctx) error {
		return controllers.GetTasksByUserID(db, c)
	})
	tasks.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateTask(db, c)
	})
	tasks.Patch("/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateTask(db, c)
	})

	log.Fatal(app.Listen(":8080"))
}
