package controllers

import (
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// POST HANDLERS

func GetPosts(db *sqlx.DB, c *fiber.Ctx) error {
	var posts []model.Post
	query := "SELECT id, user_id, content, image, created_at FROM posts ORDER BY id ASC"

	err := db.Select(&posts, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}

	return c.JSON(posts)
}

func GetPost(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var post model.Post
	query := "SELECT id, user_id, content, image, created_at FROM posts WHERE id=$1"

	err = db.Get(&post, query, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(post)
}

func CreatePost(db *sqlx.DB, c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error retrieving file",
		})
	}

	// SETUP AWS S3 UPLOADER
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load AWS config",
		})
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer f.Close()

	// UPLOADING FILE TO AWS S3
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("brainhub-server"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file to S3",
		})
	}

	// Get other form values
	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	content := c.FormValue("content")

	// Create a new post
	newPost := &model.Post{
		UserID:  userID,
		Content: content,
		Image:   result.Location,
	}

	query := `INSERT INTO posts (user_id, content, image) VALUES ($1, $2, $3) RETURNING id, created_at`
	err = db.QueryRow(query, newPost.UserID, newPost.Content, newPost.Image).Scan(&newPost.ID, &newPost.CreatedAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func UpdatePost(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	updatedPost := new(model.Post)
	if err := c.BodyParser(updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	existingPost := model.Post{}
	err = db.Get(&existingPost, "SELECT id, user_id, content, image, created_at FROM posts WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	existingPost.Content = updatedPost.Content
	existingPost.Image = updatedPost.Image

	_, err = db.Exec("UPDATE posts SET content=$1, image=$2 WHERE id=$3",
		existingPost.Content, existingPost.Image, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(existingPost)
}

func DeletePost(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var existingPost model.Post
	err = db.Get(&existingPost, "SELECT id, user_id, content, image, created_at FROM posts WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	_, err = db.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	return c.JSON(fiber.Map{"message": "Post deleted"})
}
