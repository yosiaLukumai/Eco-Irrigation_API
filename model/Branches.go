package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Branches struct {
	ID        primitive.ObjectID `bson:"_id"`
	CompanyID primitive.ObjectID `bson:"company"`
	Name      string             `bson:"name"`
	ShortForm string             `bson:"shortform"`
	CreatedAt primitive.DateTime `bson:"createdat"`
	UpdatedAt primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func CreateNewBranch(compId primitive.ObjectID, name string, ShortForm string) Branches {
	return Branches{
		ID:        primitive.NewObjectID(),
		CompanyID: compId,
		Name:      name,
		ShortForm: ShortForm,
		CreatedAt: utils.TimeLocal(),
		UpdatedAt: utils.TimeLocal(),
	}

}
