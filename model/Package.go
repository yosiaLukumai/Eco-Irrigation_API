package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Package struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	AmountPerDay  float64            `json:"amountperday" bson:"amountperday"`
	InitialAmount float64            `json:"initialamount" bson:"initialamount"`
	PowerSize     float64            `json:"powersize" bson:"powersize"`
	CreatedAt     primitive.DateTime `json:"createdat" bson:"createdat"`
	UpdatedAt     primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreatePackage(name string, amountpdy, initalA, powers float64) Package {
	return Package{
		ID:            primitive.NewObjectID(),
		Name:          name,
		AmountPerDay:  amountpdy,
		InitialAmount: initalA,
		PowerSize:     powers,
		CreatedAt:     utils.TimeLocal(),
		UpdatedAt:     utils.TimeLocal(),
	}
}
