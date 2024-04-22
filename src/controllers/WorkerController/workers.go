package workercontroller

import (
	"be_hiring_app/src/config"
	"be_hiring_app/src/helper"
	models "be_hiring_app/src/models/WorkerModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllWorkers(c *fiber.Ctx) error {
	helper.EnableCors(c)
	worker := models.SelectAllWorker()

	response := fiber.Map{
		"Message": "Success",
		"data":    worker,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetWorkerById(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectWorkerById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostWorker(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var Worker models.Worker
		if err := c.BodyParser(&Worker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Worker{
			JobDesc:     Worker.JobDesc,
			JobType:     Worker.JobType,
			Skill:       Worker.Skill,
			CompanyName: Worker.CompanyName,
			UserId:      Worker.UserId,
		}
		models.PostWorker(&item)

		return c.JSON(fiber.Map{
			"Message": "Worker Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateWorker(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)

		// Parse the request body to get the updated worker data
		var updatedWorker models.Worker
		if err := c.BodyParser(&updatedWorker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Retrieve the existing worker
		var existingWorker models.Worker
		if err := config.DB.Preload("User").First(&existingWorker, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Worker not found"})
		}

		// Update the worker fields
		existingWorker.JobDesc = updatedWorker.JobDesc
		existingWorker.JobType = updatedWorker.JobType
		existingWorker.CompanyName = updatedWorker.CompanyName
		existingWorker.Skill = updatedWorker.Skill

		// Update the associated user fields
		existingWorker.User.Name = updatedWorker.User.Name
		existingWorker.User.Address = updatedWorker.User.Address

		// Save the changes to both worker and associated user
		tx := config.DB.Begin()
		if err := tx.Save(&existingWorker.User).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
		}
		if err := tx.Save(&existingWorker).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update worker"})
		}
		tx.Commit()

		return c.JSON(fiber.Map{"message": "Worker updated successfully"})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed")
	}
}
func DeleteWorker(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteWorker(id)

	return c.JSON(fiber.Map{
		"Message": "Worker Deleted",
	})

}