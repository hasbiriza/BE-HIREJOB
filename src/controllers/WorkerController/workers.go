package workercontroller

import (
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
			Skills:      Worker.Skills,
			CompanyName: Worker.CompanyName,
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
		var Worker models.Worker
		if err := c.BodyParser(&Worker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newWorker := models.Worker{
			JobDesc:     Worker.JobDesc,
			JobType:     Worker.JobType,
			Skills:      Worker.Skills,
			CompanyName: Worker.CompanyName,
		}
		models.UpdateWorker(id, &newWorker)

		return c.JSON(fiber.Map{
			"Message": "Worker Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
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
