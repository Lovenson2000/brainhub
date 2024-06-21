package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
)

var posts = []model.Post{
	{ID: 1, UserID: 1, Title: "How to learn Go?", Content: "I'm new to Go. Any tips?"},
	{ID: 2, UserID: 2, Title: "Favorite Biology Resources", Content: "What are your favorite resources for learning biology?"},
	{ID: 3, UserID: 3, Title: "Is TailwindCSS better than vanilla css ?", Content: "What are your thoughts regarding the comparisons between TailwindCSS and Bootstrap ?"},
}

// POST HANDLERS

func GetPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
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

func CreatePost(c *fiber.Ctx) error {
	newPost := new(model.Post)

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newPost.ID = len(posts) + 1
	posts = append(posts, *newPost)

	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func DeletePost(c *fiber.Ctx) error {
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

func UpdatePost(c *fiber.Ctx) error {
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
