package routes

import (
	"context"

	"github.com/igormpb/api-go-debts/database"
	"github.com/igormpb/api-go-debts/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ShowDebts(ctx *fiber.Ctx) error {

	db := database.Connection()

	// filtro
	filter := bson.D{{Key: "account_uuid", Value: "xx"}}

	cursor, err := db.Collection("debts").Find(context.Background(), filter)
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
	}

	var items []bson.M
	if err = cursor.All(context.Background(), &items); err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
	}

	return utils.Response(ctx, fiber.StatusOK, items)
}
