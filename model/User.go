package model

import (
	"TEST_SERVER/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname" validate:"required"`
	LastName  string             `json:"lastname" bson:"lastname" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Roles     primitive.ObjectID `json:"roles,omitempty" bson:"roles,omitempty"`
	Admin     bool               `json:"admin,omitempty" bson:"admin"`
	Phone     string             `json:"phone" bson:"phone" validate:"required"`
	Verified  bool               `json:"verified" bson:"verified,omitempty"`
	Active    bool               `json:"active" bson:"active,omitempty"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	UpdatedAt primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreateAdmin(pass, email, fname, lname, phone string) User {
	fmt.Println(pass, email, lname, fname, phone)
	return User{
		ID:        primitive.NewObjectID(),
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  pass,
		Phone:     phone,
		Verified:  true,
		Active:    true,
		Admin:     true,
		CreatedAt: utils.TimeLocal(),
		UpdatedAt: utils.TimeLocal(),
	}
}
