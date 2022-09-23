package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/igormpb/api-go-debts/modules/rabbitmq"
	"github.com/igormpb/api-go-debts/routes"
)

func main() {
	app := fiber.New()

	routes.ApiV1(app)
	app.Use(logger.New())
	go rabbitmq.Sub()
	app.Listen(fmt.Sprintf(":%s", os.Getenv("GOPORT")))
}
