package main

import (
	"api-go/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	routes.ApiV1(app)
	app.Use(logger.New())

	app.Listen(fmt.Sprintf(":%s", os.Getenv("GOPORT")))
}
