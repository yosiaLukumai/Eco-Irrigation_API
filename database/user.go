package database

import (
	"TEST_SERVER/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindEmail(email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var data model.User

	err := User.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&data)
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}

func FindEmailClient(email string) (model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data model.Client

	err := Client.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Client{}, nil
		} else {
			return model.Client{}, err

		}
	}
	return data, nil
}
func UpdatePassword(userId primitive.ObjectID, hashedPassword string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data model.User
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"password": hashedPassword, "verified": true}}

	// Update password in the collection
	_, err := User.UpdateOne(ctx, filter, update)

	if err != nil {
		return model.User{}, err
	}
	return data, nil
}

func UpdatePasswordClients(userId primitive.ObjectID, hashedPassword string) (model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data model.Client
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"password": hashedPassword, "verified": true}}

	// Update password in the collection
	_, err := Client.UpdateOne(ctx, filter, update)

	if err != nil {
		return model.Client{}, err
	}
	return data, nil
}

func Verify(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"verified": true}}

	_, err := User.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	// if found update the user
	return true, nil
}
