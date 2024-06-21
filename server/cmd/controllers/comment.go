package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
)

var comments = []model.Comment{
	{ID: 1, UserID: 2, PostID: 1, Text: "Check out the Go documentation."},
	{ID: 2, UserID: 1, PostID: 2, Text: "I like Khan Academy for biology."},
	{ID: 3, UserID: 3, PostID: 3, Text: "For me, Tailwind is better faster."},
}

func GetAllComments(c *fiber.Ctx) error {
	return c.JSON(comments)
}

func GetPostComments(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("postId"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Post ID",
		})
	}

	var postComments []model.Comment

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

func CreateComment(c *fiber.Ctx) error {
	newComment := new(model.Comment)

	if err := c.BodyParser(newComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing Comment body",
		})
	}

	newComment.ID = len(comments) + 1
	comments = append(comments, *newComment)

	return c.Status(fiber.StatusCreated).JSON(newComment)
}

func DeleteComment(c *fiber.Ctx) error {
	commentID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Comment ID",
		})
	}

	// Find the index of the comment with the given ID
	index := -1
	for i, comment := range comments {
		if comment.ID == commentID {
			index = i
			break
		}
	}

	if index == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Comment not found",
		})
	}

	// Remove the comment from the slice
	comments = append(comments[:index], comments[index+1:]...)

	return c.JSON(fiber.Map{
		"message": "Comment deleted successfully",
	})
}

func UpdateComment(c *fiber.Ctx) error {
	commentID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Comment ID",
		})
	}

	updatedComment := new(model.Comment)
	if err := c.BodyParser(updatedComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing Comment body",
		})
	}

	var foundComment *model.Comment
	for i := range comments {
		if comments[i].ID == commentID {
			foundComment = &comments[i]
			break
		}
	}

	if foundComment == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Comment not found",
		})
	}

	foundComment.Text = updatedComment.Text

	return c.JSON(foundComment)
}
