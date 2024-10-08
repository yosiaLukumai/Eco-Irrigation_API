package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID            primitive.ObjectID `bson:"_id"`
	Kit           primitive.ObjectID `bson:"kit"`
	Amount        float64            `bson:"amount"`
	TransactionID string             `bson:"transactionId"`
	Phone         string             `bson:"phone"`
	Provider      string             `bson:"provider"`
	Status        bool               `bson:"status"` // pending, complete // errored
	CreatedAt     primitive.DateTime `bson:"createdat"`
	UpdatedAt     primitive.DateTime `bson:"updatedat"`
}

func CreatePayment(kit string, amount float64, tranID, provider string, phone string) Payment {
	return Payment{
		ID:            primitive.NewObjectID(),
		Kit:           utils.IDHex(kit),
		Amount:        amount,
		TransactionID: tranID,
		Phone:         phone,
		Provider:      provider,
		Status:        false,
		CreatedAt:     utils.TimeLocal(),
		UpdatedAt:     utils.TimeLocal(),
	}
}
