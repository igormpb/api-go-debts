package types

import "time"

type Debts struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	Name        string    `bson:"name" json:"name"`
	Parcel      int       `bson:"parcel,omitempty" json:"parcel"`
	AccountUUID string    `bson:"account_uuid,omitempty " json:"account_uuid"`
	Category    string    `bson:"category" json:"category"`
	Fix         bool      `bson:"fix" json:"fix"`
	DateDebt    string    `bson:"date_debt" json:"date_debt"`
	Price       int       `bson:"price" json:"price"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
