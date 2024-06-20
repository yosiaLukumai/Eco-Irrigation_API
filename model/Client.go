package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id"`
	CompanyID primitive.ObjectID `json:"companyId,omitempty" bson:"companyId,omitempty" validate:"required"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required"`
	Nida      string             `json:"nida" bson:"nida" validate:"required"`
	Contacts  string             `json:"contacts" bson:"contacts" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	MeterID   primitive.ObjectID `json:"meterID" bson:"meterID"`
	Balance   float64            `json:"balance" bson:"balance"`
	Verified  bool               `json:"verified" bson:"verified"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	UpdatedAt primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func NewClient(CompanyID primitive.ObjectID, firstName string, LastName string, Email string, Nida string, Contact string, MeterID primitive.ObjectID) Client {
	return Client{
		ID:        primitive.NewObjectID(),
		MeterID: MeterID,
		CompanyID: CompanyID,
		Name:      utils.FullName(firstName, LastName),
		Email:     Email,
		Nida:      Nida,
		Contacts:  Contact,
		Verified:  false,
		Balance:   0.0,
		CreatedAt: utils.TimeLocal(),
		UpdatedAt: utils.TimeLocal(),
	}
}
