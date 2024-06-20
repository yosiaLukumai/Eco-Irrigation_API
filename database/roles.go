package database

import (
	"TEST_SERVER/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RoleByID(roleId primitive.ObjectID) (model.Roles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	var role model.Roles
	defer cancel()
	err := Roles.FindOne(ctx, bson.M{"_id": roleId}).Decode(&role)
	if err != nil {
		return model.Roles{}, err
	}
	return role, nil
}

func InsertNewRole(data model.Roles) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := Roles.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
