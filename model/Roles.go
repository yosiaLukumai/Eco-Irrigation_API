package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Roles struct {
	ID          primitive.ObjectID `bson:"_id"`
	Company     primitive.ObjectID `bson:"company"`
	Name        string             `bson:"name"`
	Description string             `bson:"desc"`
	Access      []string           `bson:"access"`
	CreatedAt   primitive.DateTime `bson:"createdat"`
	UpdatedAt   primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreateCreatorRole(companyID primitive.ObjectID) Roles {
	return Roles{
		ID:          primitive.NewObjectID(),
		Company:     companyID,
		Name:        "Creator",
		Description: "Register of the utility company to the system",
		Access:      []string{"*"},
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
}
