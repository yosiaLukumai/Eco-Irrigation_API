package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `bson:"_id"`
	To        primitive.ObjectID `bson:"to"`
	Method    string             `bson:"method"`
	Amount    float64            `bson:"amount"`
	Status    string             `bson:"status"` // pending, complete // errored
	
	CreatedAt primitive.DateTime `bson:"createdat"`
	UpdatedAt primitive.DateTime `bson:"updatedat"`
}

var PaymentMethods []string = []string{
	"M-pesa",
	"Ezy-pesa",
	"Tigo-pesa",
	"Airtel-Money",
	"Lipa-kwa-simu",
	"Bank",
	"T-pesa",
	"Azam-pay",
	"Others",
}


func CheckPaymentMethodExists(payMethd string) bool {
	return utils.Includes(PaymentMethods, payMethd)
}