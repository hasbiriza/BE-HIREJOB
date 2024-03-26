package helper

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func Upload(c *fiber.Ctx, file multipart.File, handler *multipart.FileHeader) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("src/uploads/%s_%s", timestamp, handler.Filename)
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
