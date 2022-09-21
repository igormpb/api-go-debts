package routes

import (
	"api-go/database"
	"api-go/utils"
	"context"

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
		return err
	}

	var items []bson.M
	if err = cursor.All(context.Background(), &items); err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
		return err
	}

	utils.Response(ctx, fiber.StatusOK, items)
	return nil
}
