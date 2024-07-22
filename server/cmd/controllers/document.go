package controllers

import (
	"context"
	"log"
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
	log.Println("File sucessfully uploaded to aws S3")

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
	log.Println("document details saved to database")

	return c.Status(fiber.StatusCreated).JSON(newDocument)
}

func GetDocumentsByStudySession(db *sqlx.DB, c *fiber.Ctx) error {
	studySessionID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var documents []model.Document
	query := `SELECT id, study_session_id, title, url, uploaded_by, uploaded_at FROM documents WHERE study_session_id=$1`
	err = db.Select(&documents, query, studySessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve documents from database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(documents)
}

func GetAllDocuments(db *sqlx.DB, c *fiber.Ctx) error {
	var documents []model.Document

	query := `SELECT id, study_session_id, title, url, uploaded_by, uploaded_at FROM documents`
	err := db.Select(&documents, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve documents from database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(documents)
}

func GetDocumentsByUserID(db *sqlx.DB, c *fiber.Ctx) error {
	userID, err := util.ExtractUserIDFromJwtToken(c)
	if err != nil {
		return err
	}

	var documents []model.Document

	query := `SELECT id, study_session_id, title, url, uploaded_by, uploaded_at FROM documents WHERE uploaded_by = $1`
	err = db.Select(&documents, query, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve documents from database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(documents)
}

func DeleteDocument(db *sqlx.DB, c *fiber.Ctx) error {
	documentID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ID Invalid",
		})
	}

	var document model.Document
	query := `SELECT url FROM documents WHERE id = $1`
	err = db.Get(&document, query, documentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve document from database",
		})
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load AWS config",
		})
	}

	client := s3.NewFromConfig(cfg)
	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(document.URL),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete document from S3",
		})
	}

	query = `DELETE FROM documents WHERE id = $1`
	_, err = db.Exec(query, documentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete document from database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Document deleted successfully",
	})

}
