package database

import (
	"TEST_SERVER/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateVerification(verificationID primitive.ObjectID) (model.EmailVerification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data model.EmailVerification
	filter := bson.M{"_id": verificationID}
	update := bson.M{"$set": bson.M{"used": true}}

	// Update password in the collection
	_, err := EmailVerification.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.EmailVerification{}, err
	}

	return data, nil
}
