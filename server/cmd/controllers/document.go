package controllers

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/Lovenson2000/brainhub/pkg/model"
	"github.com/Lovenson2000/brainhub/pkg/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UploadDocument(db *sqlx.DB, c *fiber.Ctx) error {

	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error retrieving file",
		})
	}

	// SETUP AWS S3 UPLOADER
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to load AWS config",
		})
	}

	client := s3.NewFromConfig(config)
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
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file to S3",
		})
	}

	userID, err := util.ExtractUserIDFromJwtToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	studySessionID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	title := c.FormValue("title")

	newDocument := model.Document{
		StudySessionID: studySessionID,
		Title:          title,
		URL:            result.Location,
		UploadedBy:     userID,
		UploadedAt:     time.Now(),
	}

	// INSERTING INTO DB
	query := `INSERT INTO documents(study_session_id, title, url, uploaded_by, uploaded_at) 
		VALUES ($1, $2, $3, $4, $5) Returning id`
	err = db.QueryRow(query, newDocument.StudySessionID, newDocument.Title, newDocument.URL, newDocument.UploadedBy, newDocument.UploadedAt).Scan(&newDocument.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save document details to database",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newDocument)
}
