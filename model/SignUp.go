package model

type SignUp struct {
	Admin   User    `json:"user" bson:"user" validate:"required"`
	Company Company `json:"company" bson:"company"  validate:"required"`
}
