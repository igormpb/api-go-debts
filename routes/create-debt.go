package routes

import (
	"api-go/database"
	"api-go/types"
	"api-go/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateDebt(ctx *fiber.Ctx) error {
	var debts types.Debts

	db := database.Connection()
	if err := ctx.BodyParser(&debts); err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	// Pensar em como authenticar o user
	debts.AccountUUID = "XXXX-XXXX-XXXX-XXXX"
	debts.UUID = uuid.New().String()
	debts.CreatedAt = time.Now()
	debts.UpdatedAt = time.Now()
	_, err := db.Collection("debts").InsertOne(ctx.Context(), debts)
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	return utils.Response(ctx, 200, nil)
}
