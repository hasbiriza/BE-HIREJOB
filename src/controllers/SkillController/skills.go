package skillcontroller

import (
	"be_hiring_app/src/helper"
	models "be_hiring_app/src/models/SkillModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllSkills(c *fiber.Ctx) error {
	helper.EnableCors(c)
	skill := models.SelectAllSkill()

	response := fiber.Map{
		"Message": "Success",
		"data":    skill,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetSkillById(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectSkillById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostSkill(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var Skill models.Skill
		if err := c.BodyParser(&Skill); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Skill{
			Name:   Skill.Name,
			UserId: Skill.UserId,
		}
		models.PostSkill(&item)

		return c.JSON(fiber.Map{
			"Message": "Skill Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateSkill(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Skill models.Skill
		if err := c.BodyParser(&Skill); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newSkill := models.Skill{
			Name:   Skill.Name,
			UserId: Skill.UserId,
		}
		models.UpdateSkill(id, &newSkill)

		return c.JSON(fiber.Map{
			"Message": "Skill Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteSkill(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteSkill(id)

	return c.JSON(fiber.Map{
		"Message": "Skill Deleted",
	})

}
