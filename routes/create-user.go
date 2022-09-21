package routes

import (
	"api-go/database"
	"api-go/types"
	"api-go/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user types.User

	if err := ctx.BodyParser(&user); err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	if ok := validateUser(&user); !ok {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
	}

	db := database.Connection()

	user.UUID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Printf("error em : %v", err)
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "Serviço indisponivel", "XXXX")
	}

	user.Password = string(hashedPassword)

	_, errDb := db.Collection("accounts").InsertOne(ctx.Context(), user)
	if errDb != nil {
		fmt.Printf("error em : %v", errDb)
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	return utils.Response(ctx, fiber.StatusOK, nil)

}

func validateUser(user *types.User) bool {
	if user.BirthDate == "" || len(user.BirthDate) == 0 {
		return false
	}

	if user.Email == "" || len(user.Email) == 0 {
		return false
	}

	if user.Name == "" || len(user.Name) == 0 {
		return false
	}

	if user.Password == "" || len(user.Password) == 0 {
		return false
	}

	if user.Username == "" || len(user.Username) == 0 {
		return false
	}
	return true
}
