package projectcontroller

import (
	"be_hiring_app/src/helper"
	models "be_hiring_app/src/models/ProjectModel"
	"encoding/json"
	"net/http"
	"strconv"

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
	if c.Method() == fiber.MethodPost {
		var Project models.Project
		if err := c.BodyParser(&Project); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Project{
			Title:       Project.Title,
			Photo:       Project.Photo,
			Repository:  Project.Repository,
			ProjectType: Project.ProjectType,
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
