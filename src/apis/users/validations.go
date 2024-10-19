package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUsersValidation(c *fiber.Ctx) error {
	println("in GetUsersValidation validation")
	fmt.Println(c.Query("limit"))
	return c.Next()
}

func GetUserValidation(c *fiber.Ctx) error {
	println("in GetUserValidation validation")
	fmt.Println(c.Params("id"))
	return c.Next()
}

func PostUserValidation(c *fiber.Ctx) error {
	println("in PostUserValidation validation")
	var out any
	c.BodyParser(&out)
	fmt.Println(out)
	return c.Next()
}

func PatchUserValidation(c *fiber.Ctx) error {
	println("in PatchUserValidation validation")
	fmt.Println(c.Get("authorization"))
	c.Locals("sessions", "sessionX")
	return c.Next()
}

func DeleteUserValidation(c *fiber.Ctx) error {
	println("in DeleteUserValidation validation")
	var out any
	c.BodyParser(&out)
	fmt.Println(out)
	return c.Next()
}
