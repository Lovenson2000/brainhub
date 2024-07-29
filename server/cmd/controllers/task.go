package controllers

import (
	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetTasks(db *sqlx.DB, c *fiber.Ctx) error {
	var tasks []model.Task

	query := "SELECT * FROM tasks ORDER BY id ASC"
	err := db.Select(&tasks, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving tasks",
		})
	}
	return c.JSON(tasks)
}

func GetTasksByUserID(db *sqlx.DB, c *fiber.Ctx) error {
	userID, err := util.ExtractUserIDFromJwtToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	var tasks []model.Task

	query := "SELECT id, user_id, title, description, start_time, due_date, status, priority FROM tasks WHERE user_id=$1"
	err = db.Select(&tasks, query, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tasks for the user",
		})
	}

	return c.JSON(tasks)
}

func GetTask(db *sqlx.DB, c *fiber.Ctx) error {
	taskID := c.Params("id")
	var task model.Task

	query := "SELECT * FROM tasks WHERE id=$1"
	err := db.Get(&task, query, taskID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	return c.JSON(task)
}

func UpdateTask(db *sqlx.DB, c *fiber.Ctx) error {
	taskID := c.Params("id")
	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	query := `UPDATE tasks SET title=$1, description=$2, start_time=$3, due_date=$4, status=$5, priority=$6 WHERE id=$7`
	_, err := db.Exec(query, query, task.Title, task.Description, task.StartTime, task.DueDate, task.Status, task.Priority, taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.JSON(fiber.Map{"message": "Task updated successfully"})
}
