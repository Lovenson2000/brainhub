package controllers

import (
	"strconv"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetStudySessions(db *sqlx.DB, c *fiber.Ctx) error {

	var studySessions []model.StudySession

	query := "SELECT id, title, description, start_time, end_time, location, participants FROM study_sessions ORDER BY id ASC"
	err := db.Select(&studySessions, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get study sessions",
		})

	}
	return c.JSON(studySessions)
}

func GetStudySession(db sqlx.DB, c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	var studySession model.StudySession
	query := "SELECT id, title, description, start_time, end_time, location, participants FROM study_sessions WHERE id=$1"
	err = db.Get(&studySession, query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Study Session not found",
		})
	}
	return c.JSON(studySession)
}

func CreateStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	newStudySession := new(model.StudySession)
	if err := c.BodyParser(newStudySession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	query := `INSERT INTO study_sessions (title, description, start_time, end_time, participants) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, newStudySession.Title, newStudySession.Description, newStudySession.StartTime, newStudySession.EndTime, pq.Array(newStudySession.Participants)).Scan(&newStudySession.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create study session",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newStudySession)
}
