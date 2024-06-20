package model

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Phone      string             `json:"phone" bson:"phone" validate:"required"`
	Location   string             `json:"location" bson:"location" validate:"required"`
	MeterCount uint64             `json:"metercount,omitempty" bson:"metercount"` // global count to set value of the meters
	CreatedAt  primitive.DateTime `bson:"createdat" validate:"required"`
	UpdatedAt  primitive.DateTime `bson:"updatedat" validate:"required"`
}

func (s *SignUp) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
