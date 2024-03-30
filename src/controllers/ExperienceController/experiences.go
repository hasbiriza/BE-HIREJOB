package experiencecontroller

import (
	"be_hiring_app/src/dtos"
	models "be_hiring_app/src/models/ExperienceModel"
	"be_hiring_app/src/services"
	"encoding/json"

	"path/filepath"
	"strings"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllExperiences(c *fiber.Ctx) error {
	experience := models.SelectAllExperience()

	response := fiber.Map{
		"Message": "Success",
		"data":    experience,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetExperienceById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectExperienceById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostExperience(c *fiber.Ctx) error {
	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20 // 2 MB
	)

	if c.Method() == fiber.MethodPost {
		var Experience models.Experience
		if err := c.BodyParser(&Experience); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a file to upload"},
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
		uploadUrl, err := services.NewMediaUpload().FileUploadExperience(models.File{File: formFile})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
		}

		// Create Experience object with uploaded file URL
		item := models.Experience{
			Position:    Experience.Position,
			Photo:       uploadUrl,
			StartDate:   Experience.StartDate,
			EndDate:     Experience.EndDate,
			Description: Experience.Description,
			UserId:      Experience.UserId,
		}

		models.PostExperience(&item)

		return c.JSON(fiber.Map{
			"Message": "Experience Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed")
	}
}

func UpdateExperience(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Experience models.Experience
		if err := c.BodyParser(&Experience); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newExperience := models.Experience{
			Position:    Experience.Position,
			Photo:       Experience.Photo,
			StartDate:   Experience.StartDate,
			EndDate:     Experience.EndDate,
			Description: Experience.Description,
		}
		models.UpdateExperience(id, &newExperience)

		return c.JSON(fiber.Map{
			"Message": "Experience Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteExperience(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteExperience(id)

	return c.JSON(fiber.Map{
		"Message": "Experience Deleted",
	})

}

func RemoteUpload(c *fiber.Ctx) error {
	var url models.Experience

	if err := c.BodyParser(&url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dtos.MediaDto{
				StatusCode: fiber.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Error uploading file"},
			})
	}

	return c.Status(fiber.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: fiber.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"data": uploadUrl},
		})
}
