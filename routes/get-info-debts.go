package routes

import (
	"api-go/database"
	"api-go/types"
	"api-go/utils"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetInfoDebts(ctx *fiber.Ctx) error {

	var debt types.Debts
	db := database.Connection()

	err := db.Collection("debts").FindOne(context.Background(), bson.D{{Key: "uuid", Value: "x"}}).Decode(&debt)
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	return utils.Response(ctx, fiber.StatusOK, debt)
}
