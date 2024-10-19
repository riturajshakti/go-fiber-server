package users

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetUsersController(c *fiber.Ctx) error {
	fmt.Println("in GetUsersController controller")
	return c.JSON([]any{
		map[string]any{
			"name": "raj",
			"age":  30,
		},
		map[string]any{
			"name": "john",
			"age":  27,
		},
	})
}

func GetUserController(c *fiber.Ctx) error {
	println("in GetUserController controller")
	return c.JSON(map[string]any{
		"name": "raj",
		"age":  30,
	})
}

func PostUserController(c *fiber.Ctx) error {
	println("in PostUserController controller")
	return c.JSON(map[string]any{
		"name": "raj",
		"age":  30,
	})
}

func PatchUserController(c *fiber.Ctx) error {
	println("in PatchUserController controller")
	session, ok := c.Locals("sessions").(string)
	if !ok {
		return c.Status(401).JSON(map[string]any{"message": "authorization failed"})
	}
	return c.JSON(map[string]any{
		"name":    "raj",
		"age":     30,
		"session": session,
	})
}

func PutUserController(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": "unable to parse form",
		})
	}
	image, ok := form.File["image"]
	if !ok {
		return c.Status(400).JSON(map[string]string{
			"message": "image is required",
		})
	}
	fmt.Println(image)

	uploadsDir := "./uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err = os.Mkdir(uploadsDir, 0755)
		if err != nil {
			log.Println(err)
			return c.Status(500).JSON(map[string]string{
				"message": "internal server error",
			})
		}
	}

	// Save image to uploads directory
	for _, file := range image {
		filename := filepath.Base(file.Filename)
		if err := c.SaveFile(file, filepath.Join(uploadsDir, filename)); err != nil {
			log.Println(err)
			return c.Status(500).JSON(map[string]string{
				"message": "failed to save image",
			})
		}
	}

	return c.JSON(form)
}

func DeleteUserController(c *fiber.Ctx) error {
	println("in DeleteUserController controller")
	panic("exiting")
}
