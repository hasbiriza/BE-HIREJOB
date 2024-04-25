package projectcontroller

import (
	"be_hiring_app/src/dtos"
	"be_hiring_app/src/helper"
	models "be_hiring_app/src/models/ProjectModel"
	"be_hiring_app/src/services"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllProjects(c *fiber.Ctx) error {
	helper.EnableCors(c)
	projects := models.SelectAllProject()

	response := fiber.Map{
		"Message": "Success",
		"data":    projects,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetProjectById(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectProjectById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostProject(c *fiber.Ctx) error {
	helper.EnableCors(c)
	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20 // 2 MB
	)
	if c.Method() == fiber.MethodPost {

		var Project models.Project
		if err := c.BodyParser(&Project); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		formHeader, err := c.FormFile("photo")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a photo to upload"},
			})
		}
		formFile, err := formHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
		}
		defer formFile.Close()

		ext := filepath.Ext(formHeader.Filename)
		ext = strings.ToLower(ext)
		allowedExts := strings.Split(AllowedExtensions, ",")
		validExtension := false
		for _, allowedExt := range allowedExts {
			if ext == allowedExt {
				validExtension = true
				break
			}
		}
		if !validExtension {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": "Invalid file extension"},
			})
		}

		// Validate file size
		fileSize := formHeader.Size
		if fileSize > MaxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": "File size exceeds the allowed limit"},
			})
		}

		// Upload the file
		uploadUrl, err := services.NewMediaUpload().FileUploadProject(models.File{File: formFile})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
		}

		// Create Experience object with uploaded file URL
		item := models.Project{
			Title:       Project.Title,
			Photo:       uploadUrl,
			Repository:  Project.Repository,
			ProjectType: Project.ProjectType,
			UserId:      Project.UserId,
		}
		models.PostProject(&item)

		return c.JSON(fiber.Map{
			"Message": "Project Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateProject(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Project models.Project
		if err := c.BodyParser(&Project); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newProject := models.Project{
			Title:       Project.Title,
			Photo:       Project.Photo,
			Repository:  Project.Repository,
			ProjectType: Project.ProjectType,
		}
		models.UpdateProject(id, &newProject)

		return c.JSON(fiber.Map{
			"Message": "Project Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteProject(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteProject(id)

	return c.JSON(fiber.Map{
		"Message": "Project Deleted",
	})

}
