package database

import (
	"TEST_SERVER/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertCompany(companyData model.Company) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var returned interface{}

	// for clients & users emails are the defacto...
	err := Company.FindOne(ctx, bson.D{{Key: "name", Value: companyData.Name}}).Decode(&returned)
	if err != nil {
		if !(err == mongo.ErrNoDocuments) {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("company Name alredy taken")
	}
	// meter count
	companyData.MeterCount = 0
	result, err := Company.InsertOne(ctx, companyData)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func InsertUser(UserData model.User) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := User.InsertOne(ctx, UserData)
	if err != nil {
		fmt.Println("Error is ", err)
		return nil, err
	}
	return result.InsertedID, nil
}

func FindCompany(searchID primitive.ObjectID) (model.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data model.Company
	filter := bson.M{"_id": bson.M{"$eq": searchID}}
	err := Company.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return model.Company{}, err
	}
	return data, nil
}

func UpdateCompanyMeterCount(col *mongo.Collection, userId primitive.ObjectID, newValue uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"metercount": newValue}}

	// Update password in the collection
	_, err := col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}
	return true, nil
}
