package controllers

import (
	"strconv"
	"time"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/gofiber/fiber/v2"
)

var studySessions = []model.StudySession{
	{ID: 1, Title: "Go Programming Basics", Description: "Introduction to Go programming", StartTime: time.Now(), EndTime: time.Now().Add(2 * time.Hour), Location: "Library Room 101", Participants: []int{1, 2}},
	{ID: 2, Title: "Biology Study Group", Description: "Study group for biology majors", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour), Location: "Biology Lab", Participants: []int{2}},
}



func GetStudySessions(c *fiber.Ctx) error {
	return c.JSON(studySessions)
}

func GetStudySession(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	for _, studySession := range studySessions {
		if studySession.ID == id {
			return c.JSON(studySession)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Study Session not found",
	})
}

func CreateStudySession(c *fiber.Ctx) error {
	newStudySession := new(model.StudySession)

	if err := c.BodyParser(newStudySession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "couldn't create Study Session",
		})
	}

	newStudySession.ID = len(studySessions) + 1
	studySessions = append(studySessions, *newStudySession)

	return c.Status(fiber.StatusCreated).JSON(newStudySession)
}

func UpdateStudySession(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	updatedStudySession := new(model.StudySession)
	if err := c.BodyParser(updatedStudySession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing Study Session body",
		})
	}

	for i, studySession := range studySessions {
		if studySession.ID == id {
			studySessions[i].Title = updatedStudySession.Title
			studySessions[i].Description = updatedStudySession.Description
			studySessions[i].StartTime = updatedStudySession.StartTime
			studySessions[i].EndTime = updatedStudySession.EndTime
			studySessions[i].Location = updatedStudySession.Location
			studySessions[i].Participants = updatedStudySession.Participants
			return c.JSON(studySessions[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Study Session not found",
	})
}

func DeleteStudySession(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Study Session ID",
		})
	}

	for i, studySession := range studySessions {
		if studySession.ID == id {
			studySessions = append(studySessions[:i], studySessions[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Study Session not found",
	})
}
