package main

import (
	"log"
	"os"

	"github.com/Lovenson2000/brainhub/cmd/controllers"
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

	// Drop tables if they exist (for testing purposes)
	_, err = db.Exec(`DROP TABLE IF EXISTS posts, users CASCADE`)
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	// Create the 'users' table if it does not exist, matching the User struct fields
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			firstname TEXT,
			lastname TEXT,
			email TEXT UNIQUE,
			password TEXT,
			school TEXT,
			major TEXT,
			bio TEXT
		)`)

	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	// Create the 'posts' table if it does not exist, matching the User struct fields
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		content TEXT,
		image TEXT,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`)

	if err != nil {
		log.Fatal("Failed to create posts table:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Change this to specific origins if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

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
	app.Get("/api/study-sessions", controllers.GetStudySessions)
	app.Get("/api/study-sessions/:id", controllers.GetStudySession)
	app.Post("/api/study-sessions", controllers.CreateStudySession)
	app.Delete("/api/study-sessions/:id", controllers.DeleteStudySession)
	app.Patch("/api/study-sessions/:id", controllers.UpdateStudySession)

	log.Fatal(app.Listen(":8080"))
}
