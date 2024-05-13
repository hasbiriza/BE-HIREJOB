package usercontroller

import (
	"be_hiring_app/src/dtos"
	"be_hiring_app/src/helper"
	models "be_hiring_app/src/models/UserModel"
	"be_hiring_app/src/services"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
	MaxFileSize       = 2 << 20 // 2 MB
)

func RegisterWorker(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    Password,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Description: user.Description,
			UserToken:   "-",
			Instagram:   user.Instagram,
			Github:      user.Github,
			Linkedin:    user.Linkedin,
			Role:        "Worker",
		}
		models.PostUser(&item)

		return c.JSON(fiber.Map{
			"Message": "Worker Registered",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func RegisterRecruiter(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    Password,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Description: user.Description,
			UserToken:   "-",
			Instagram:   user.Instagram,
			Github:      "-",
			Linkedin:    user.Linkedin,
			Role:        "Recruiter",
		}
		models.PostUser(&item)

		return c.JSON(fiber.Map{
			"Message": "User Registered",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func Login(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		users, err := models.FindEmail(user.Email)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Email Not Found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password)); err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Password Not found")
		}

		jwtKey := os.Getenv("SECRETKEY")
		token, err := helper.GenerateToken(jwtKey, users.Email, users.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed To Generate Tokens")
		}

		payload := fiber.Map{
			"Message": "HI, " + users.Name + " as a " + users.Role,
			"Email":   user.Email,
			"Role":    users.Role,
			"Token":   token,
			"User_ID": users.ID,
		}
		return c.JSON(payload)
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func GetAllUsers(c *fiber.Ctx) error {
	helper.EnableCors(c)
	users := models.SelectAllUser()

	response := fiber.Map{
		"Message": "Success",
		"data":    users,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal Konversi Json")
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetUserById(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectUserById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func UpdateWorker(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var user models.User
		if err := c.BodyParser(&user); err != nil {
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

		fileSize := formHeader.Size
		if fileSize > MaxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": "File size exceeds the allowed limit"},
			})
		}

		uploadUrl, err := services.NewMediaUpload().FileUploadUser(models.File{File: formFile})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newUser := models.User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    Password,
			PhoneNumber: user.PhoneNumber,
			Photo:       uploadUrl,
			Address:     user.Address,
			Description: user.Description,
			Instagram:   user.Instagram,
			Github:      user.Github,
			Linkedin:    user.Linkedin,
		}
		models.UpdateUser(id, &newUser)

		return c.JSON(fiber.Map{
			"Message": "User Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateRecruiter(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var user models.User
		if err := c.BodyParser(&user); err != nil {
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

		fileSize := formHeader.Size
		if fileSize > MaxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": "File size exceeds the allowed limit"},
			})
		}

		uploadUrl, err := services.NewMediaUpload().FileUploadUser(models.File{File: formFile})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.MediaDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newUser := models.User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    Password,
			PhoneNumber: user.PhoneNumber,
			Photo:       uploadUrl,
			Address:     user.Address,
			Description: user.Description,
			Instagram:   user.Instagram,
			Linkedin:    user.Linkedin,
		}
		models.UpdateUser(id, &newUser)

		return c.JSON(fiber.Map{
			"Message": "User Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteUser(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteUser(id)

	return c.JSON(fiber.Map{
		"Message": "User Deleted",
	})

}
