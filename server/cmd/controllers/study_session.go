package controllers

import (
	"log"
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetStudySessions(db *sqlx.DB, c *fiber.Ctx) error {
	var studySessions []model.StudySession

	query := "SELECT id, user_id, title, description, start_time, end_time FROM study_sessions"
	err := db.Select(&studySessions, query)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get study sessions",
		})
	}

	return c.JSON(studySessions)
}

func GetStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	var studySession model.StudySession
	query := "SELECT id, user_id, title, description, start_time, end_time, FROM study_sessions WHERE id=$1"
	err = db.Get(&studySession, query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Study Session not found",
		})
	}
	return c.JSON(studySession)
}

func GetStudySessionsByUserId(db *sqlx.DB, c *fiber.Ctx) error {
	userID, err := util.ExtractUserIDFromJwtToken(c)
	if err != nil {
		return err
	}

	var studySessions []model.StudySession

	query := "SELECT id, user_id, title, description, start_time, end_time FROM study_sessions WHERE user_id=$1"
	err = db.Select(&studySessions, query, userID)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get study sessions",
		})
	}

	return c.JSON(studySessions)
}

func CreateStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	log.Println("Starting CreateStudySession")

	newStudySession := new(model.StudySession)

	if err := c.BodyParser(newStudySession); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	userID, err := util.ExtractUserIDFromJwtToken(c)
	if err != nil {
		log.Println("Error extracting user ID from JWT token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	newStudySession.UserID = userID

	log.Printf("Creating study session for user ID %d with title %s", newStudySession.UserID, newStudySession.Title)

	query := `INSERT INTO study_sessions (user_id, title, description, start_time, end_time) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.QueryRow(query, newStudySession.UserID, newStudySession.Title, newStudySession.Description, newStudySession.StartTime, newStudySession.EndTime).Scan(&newStudySession.ID)
	if err != nil {
		log.Println("Error inserting study session into database:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create study session",
		})
	}

	log.Printf("Successfully created study session with ID %d for user ID %d", newStudySession.ID, newStudySession.UserID)
	return c.Status(fiber.StatusCreated).JSON(newStudySession)
}

func UpdateStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	updatedStudySession := new(model.StudySession)
	if err := c.BodyParser(updatedStudySession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	query := `UPDATE study_sessions SET title=$1, description=$2, start_time=$3, end_time=$4 WHERE id=$5`
	_, err = db.Exec(query, updatedStudySession.Title, updatedStudySession.Description, updatedStudySession.StartTime, updatedStudySession.EndTime, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update study session",
		})
	}
	return c.JSON(updatedStudySession)
}

func DeleteStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	_, err = db.Exec("DELETE FROM study_sessions WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete study session",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
