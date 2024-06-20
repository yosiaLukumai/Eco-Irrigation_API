package model

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meter struct {
	ID          primitive.ObjectID `bson:"_id"`
	SerialNo    string             `bson:"serialno"`
	CompanyID   primitive.ObjectID `bson:"company"`
	Branch      primitive.ObjectID `bson:"branch"`
	MaxPressure float64            `bson:"maxpressure"`
	FlowRate    float64            `bson:"flowrate"`
	Status      bool               `bson:"status"`   // of or active
	Amount      float64            `bson:"amount"`   // currently unit of water
	Assigned    bool               `bson:"assigned"` // assigned to which user
	CreatedAt   primitive.DateTime `bson:"createdat"`
	UpdatedAt   primitive.DateTime `json:"updatedat" bson:"updatedat"`
}

func NewMeter(compId, branchs string, companyName string, branchShortForm string, meterCount uint64, flowrate float64, maxpressure float64) Meter {
	return Meter{
		ID:          primitive.NewObjectID(),
		CompanyID:   utils.IDHex(compId),
		SerialNo:    utils.GenerateSerialNumber(companyName, branchShortForm, meterCount),
		Branch:      utils.IDHex(branchs),
		Status:      false,
		Amount:      0.0,
		MaxPressure: maxpressure,
		FlowRate:    flowrate,
		Assigned:    false,
		CreatedAt:   utils.TimeLocal(),
		UpdatedAt:   utils.TimeLocal(),
	}
}
