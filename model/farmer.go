package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Farmer struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PackageID primitive.ObjectID `json:"package" bson:"package"`
	FirstName string             `json:"firstname" bson:"firstname" validate:"required"`
	LastName  string             `json:"lastname" bson:"lastname" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required"`
	Location  string             `json:"location" bson:"location" validate:"required"`
	Password  string             `json:"password" bason:"password" validate:"required"`
	Phone     string             `json:"phone" bson:"phone" validate:"required"`
	Verified  bool               `json:"verified" bson:"verified,omitempty"`
	Active    bool               `json:"active" bson:"active,omitempty"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	UpdatedAt primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreateNewFarmer(email, packae, fname, lname, phone string) Farmer {
	return Farmer{
		ID:        primitive.NewObjectID(),
		PackageID: utils.IDHex(packae),
		FirstName: fname,
		LastName:  lname,
		Phone:     phone,
		Verified:  false,
		Active:    false,
		Email:     email,
		CreatedAt: utils.TimeLocal(),
		UpdatedAt: utils.TimeLocal(),
	}
}
