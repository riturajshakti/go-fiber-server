package users

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App, path string) {
	router := app.Group(path)

	router.Get("/", GetUsersValidation, GetUsersController)
	router.Get("/:id", GetUserValidation, GetUserController)
	router.Post("/", PostUserValidation, PostUserController)
	router.Put("/:id", PutUserController)
	router.Patch("/:id", PatchUserValidation, PatchUserController)
	router.Delete("/:id", DeleteUserValidation, DeleteUserController)
}
