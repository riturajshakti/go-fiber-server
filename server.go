package main

import (
	"fmt"
	"log"
	"time"

	"server/src/apis/users"
	"server/src/utils"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.InitEnv()

	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024,
	})
	app.Use(logger.New(logger.Config{
		Format: "${status} ${method} ${path}${latency}\n",
	}))
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(recover.New())
	app.Static("/", "./public")

	users.InitRoutes(app, "/api/users")

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s Server started!\n", green("Info:"))
	log.Fatal(app.Listen(fmt.Sprintf(":%s", utils.Env.Port)))
}
