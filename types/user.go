package types

import "time"

type User struct {
	UUID      string    `bson:"uuid" json:"uuid"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	Username  string    `bson:"username" json:"username"`
	BirthDate string    `bson:"birthdate" json:"birthdate"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
