package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// POST HANDLERS

func GetPosts(db *sqlx.DB, c *fiber.Ctx) error {
	var posts []model.Post
	query := "SELECT id, user_id, content, image FROM posts ORDER BY id ASC"

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
	query := "SELECT id, user_id, content, image FROM posts WHERE id=$1"

	err = db.Get(&post, query, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(post)
}

func CreatePost(db *sqlx.DB, c *fiber.Ctx) error {
	newPost := new(model.Post)

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	query := `INSERT INTO posts (user_id, content, image) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, newPost.UserID, newPost.Content, newPost.Image).Scan(&newPost.ID)
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
	err = db.Get(&existingPost, "SELECT id, user_id, content, image FROM posts WHERE id=$1", id)
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
	err = db.Get(&existingPost, "SELECT id, user_id content, image FROM posts WHERE id=$1", id)
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
