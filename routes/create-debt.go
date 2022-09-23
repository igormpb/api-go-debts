package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/api-go-debts/database"
	"github.com/igormpb/api-go-debts/types"
	"github.com/igormpb/api-go-debts/utils"
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
