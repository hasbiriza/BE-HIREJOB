package experiencecontroller

import (
	models "be_hiring_app/src/models/ExperienceModel"
	"be_hiring_app/src/services"
	"be_hiring_app/src/dtos"
	"encoding/json"
	"net/http"

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
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
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
	if c.Method() == fiber.MethodPost {
		var Experience models.Experience
		if err := c.BodyParser(&Experience); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Experience{
			Position:    Experience.Position,
			Photo:       Experience.Photo,
			StartDate:   Experience.StartDate,
			EndDate:     Experience.EndDate,
			Description: Experience.Description,
		}
		models.PostExperience(&item)

		return c.JSON(fiber.Map{
			"Message": "Experience Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
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

func FileUpload(c *fiber.Ctx) error {
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a file to upload"},
			})
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"data": uploadUrl},
		})
}

func RemoteUpload(c *fiber.Ctx) error {
	var url models.Experience

	if err := c.BodyParser(&url); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Error uploading file"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"data": uploadUrl},
		})
}
