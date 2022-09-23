package routes

import (
	"github.com/igormpb/api-go-debts/database"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/igormpb/api-go-debts/types"
	"github.com/igormpb/api-go-debts/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/api-go-debts/modules/rabbitmq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user types.User

	if err := ctx.BodyParser(&user); err != nil {
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, "Serviço indisponivel", "XXXX")
	}

	if ok, msgError := validateUser(&user); !ok {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, fmt.Sprintf("%s é obrigatório", msgError), "XXXX")
	}

	db := database.Connection()

	if existAccount := existAccount(db, &user); existAccount {
		return utils.ResponseError(ctx, fiber.StatusBadRequest, "E-mail ou usuário já estão sendo utilizado", "XXXX")
	}

	user.UUID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = "pending_confirm"

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

	var queue types.QueueEmail

	queue.AccountUUID = user.UUID
	item, err := json.Marshal(queue)
	if err != nil {
		log.Println(err.Error())
	}

	err = rabbitmq.Publish("accounts", "email", item)

	if err != nil {
		log.Println(err.Error())
		return utils.ResponseError(ctx, fiber.StatusInternalServerError, err.Error(), "XXXX")

	}
	return utils.Response(ctx, fiber.StatusOK, nil)

}

func validateUser(user *types.User) (bool, string) {
	if user.BirthDate == "" || len(user.BirthDate) == 0 {
		return false, "Data de nascimento"
	}

	if user.Email == "" || len(user.Email) == 0 {
		return false, "E-mail"
	}

	if user.Name == "" || len(user.Name) == 0 {
		return false, "Nome"
	}

	if user.Password == "" || len(user.Password) == 0 {
		return false, "Senha"
	}

	if user.Username == "" || len(user.Username) == 0 {
		return false, "Usuário"
	}
	return true, ""
}

func existAccount(db *mongo.Database, account *types.User) bool {
	ctx := context.Background()
	filter := bson.D{{"$or", bson.A{
		bson.D{{Key: "email", Value: account.Email}}, bson.D{{Key: "username", Value: account.Username}},
	}}}

	acc, err := db.Collection("accounts").Find(ctx, filter)
	if err != nil {
		return true
	}

	var items []bson.M
	err = acc.All(ctx, &items)
	if err != nil {
		return true
	}

	if len(items) > 0 || len(items) != 0 {
		return true
	}
	return false
}
