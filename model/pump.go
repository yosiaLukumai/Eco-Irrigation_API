package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pump struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FarmerID  primitive.ObjectID `json:"farmer" bson:"farmer"`
	Disharge  float64            `json:"discharge" bson:"discharge"`
	Head      float64            `json:"head" bson:"head"`
	Status    bool               `json:"status" bson:"status"`
	Balance   float64            `json:"balance" bson:"balance"`
	Assigned  bool               `json:"assigned" bson:"assigned"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	UpdatedAt primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreateNewPump(discharge, head float64) Pump {
	return Pump{
		ID: primitive.NewObjectID(),
		// FarmerID:  utils.IDHex(farmerID),
		Disharge:  discharge,
		Head:      head,
		Status:    false,
		Assigned:  false,
		Balance:   0.0,
		CreatedAt: utils.TimeLocal(),
		UpdatedAt: utils.TimeLocal(),
	}
}
