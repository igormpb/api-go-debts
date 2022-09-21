package routes

import (
	"api-go/database"
	"api-go/types"
	"api-go/utils"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {
	db := database.Connection()

	var body types.User
	var user types.User

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel 1", "XXXX")
	}

	if ok := validadeLogin(&body); !ok {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
	}

	filter := bson.D{{Key: "username", Value: body.Username}}

	err := db.Collection("accounts").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel 2", "XXXX")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Senha inválida.", "XXXX")
	}

	return utils.Response(ctx, fiber.StatusOK, nil)

}

func validadeLogin(user *types.User) bool {

	if user.Username == "" || len(user.Username) == 0 {
		return false
	}

	if user.Password == "" || len(user.Password) == 0 {
		return false
	}

	return true

}
