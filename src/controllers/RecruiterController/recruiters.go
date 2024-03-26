package recruitercontroller

import (
	models "be_hiring_app/src/models/RecruiterModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRecruiters(c *fiber.Ctx) error {
	recruiter := models.SelectAllRecruiter()

	response := fiber.Map{
		"Message": "Success",
		"data":    recruiter,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetRecruiterById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectRecruiterById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostRecruiter(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var Recruiter models.Recruiter
		if err := c.BodyParser(&Recruiter); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Recruiter{
			CompanyName: Recruiter.CompanyName,
			Position:    Recruiter.Position,
		}
		models.PostRecruiter(&item)

		return c.JSON(fiber.Map{
			"Message": "Recruiter Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateRecruiter(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Recruiter models.Recruiter
		if err := c.BodyParser(&Recruiter); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newRecruiter := models.Recruiter{
			CompanyName: Recruiter.CompanyName,
			Position:    Recruiter.Position,
		}
		models.UpdateRecruiter(id, &newRecruiter)

		return c.JSON(fiber.Map{
			"Message": "Recruiter Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteRecruiter(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteRecruiter(id)

	return c.JSON(fiber.Map{
		"Message": "Recruiter Deleted",
	})

}
