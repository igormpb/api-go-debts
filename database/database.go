package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Database {
	ctx := context.TODO()

	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.n7loi.mongodb.net/?retryWrites=true&w=majority", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	println("DATABASE SUCCESS CONNECT!")
	return client.Database("finup")
}
