package main

import (
	"log"
	"os"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
	"github.com/Lovenson2000/brainhub/cmd/middleware"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
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

	// TODO: TO BE CHANGED IN PRODUCTION
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(middleware.Logger())
	app.Use(middleware.ErrorHandler())

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

	// Auth routes
	app.Post("/api/register", func(c *fiber.Ctx) error {
		return controllers.RegisterUser(db, c)
	})
	app.Post("/api/login", func(c *fiber.Ctx) error {
		return controllers.LoginUser(db, c)
	})

	// Comment routes
	app.Get("/api/comments", controllers.GetAllComments)
	app.Get("/api/posts/:postId/comments", controllers.GetPostComments)
	app.Post("/api/comments", controllers.CreateComment)
	app.Delete("/api/comments/:id", controllers.DeleteComment)
	app.Patch("/api/comments/:id", controllers.UpdateComment)

	// // StudySession routes
	app.Get("/api/study-sessions", func(c *fiber.Ctx) error {
		return controllers.GetStudySessions(db, c)
	})

	app.Get("/api/study-sessions/:id", func(c *fiber.Ctx) error {
		return controllers.GetStudySession(db, c)
	})

	app.Post("/api/study-sessions", func(c *fiber.Ctx) error {
		return controllers.CreateStudySession(db, c)
	})
	app.Delete("/api/study-sessions/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteStudySession(db, c)
	})
	app.Patch("/api/study-sessions/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateStudySession(db, c)
	})

	log.Fatal(app.Listen(":8080"))
}
