package routes

import "github.com/gofiber/fiber/v2"

func ApiV1(app *fiber.App) {
	api := app.Group("api/v1/organize")

	// dividas
	api.Post("/debts/create", CreateDebt)
	api.Get("/debts/show", ShowDebts)

	// contas
	api.Post("/accounts/create", CreateUser)
	api.Post("/accounts/login", Login)

}
