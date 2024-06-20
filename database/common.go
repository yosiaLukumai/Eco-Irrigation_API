package database

import (
	"TEST_SERVER/model"
	"TEST_SERVER/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TotalCount struct {
	Count int64 `bson:"count"`
}

func UpdateOne(col *mongo.Collection, filter bson.M, update bson.M) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Update password in the collection
	_, err := col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}
	return true, nil
}

func Find(mongoCollection *mongo.Collection, key string, value string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var data interface{}

	err := mongoCollection.FindOne(ctx, bson.D{{Key: key, Value: value}}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func FindFilter(mongoCollection *mongo.Collection, filter bson.D) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var data interface{}

	err := mongoCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func FindFilterMap(mongoCollection *mongo.Collection, filter bson.M) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var data interface{}

	err := mongoCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func SaveVerification(verificationDetails model.EmailVerification) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result, err := EmailVerification.InsertOne(ctx, verificationDetails)
	if err != nil {
		fmt.Println("Error is ", err)
		return nil, err
	}
	return result.InsertedID, nil
}

func FindKey(key string) (model.EmailVerification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var data model.EmailVerification

	err := EmailVerification.FindOne(ctx, bson.D{{Key: "token", Value: key}}).Decode(&data)
	if err != nil {
		return model.EmailVerification{}, err
	}

	return data, nil
}

func InsertOne(col *mongo.Collection, data interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	data, err := col.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FindColl(col *mongo.Collection, filter bson.M) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	// iterating over the cursor for the datax

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil

}

func FindCollArrayTable(col *mongo.Collection, filter bson.A, initial bool) (utils.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := col.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}
	if initial {
		var results bson.M
		utils.AggregatorDecoder(ctx, cursor, &results)
		data, ok := results["data"].(bson.A)
		if !ok {
			log.Panicln("failed to convert pipeline's data to bson.A")
		}
		count, ok := results["count"].(bson.A)
		if !ok {
			log.Panicln("failed to convert pipeline's data to bson.A")
		}

		var docFound int64
		if len(count) > 0 {
			var totalDocs TotalCount
			utils.Doconveter(count[0], &totalDocs)
			docFound = totalDocs.Count
		}
		return utils.Map{"data": data, "count": docFound}, nil
	} else {
		var results []bson.M
		if err = cursor.All(ctx, &results); err != nil {
			return nil, err
		}
		cursor.Close(ctx)
		return utils.Map{"data": results}, nil
	}
}

func FindCollArrayTableMain(col *mongo.Collection, filter mongo.Pipeline, initial bool) (utils.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := col.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}
	if initial {
		var results bson.M
		utils.AggregatorDecoder(ctx, cursor, &results)
		data, ok := results["data"].(bson.A)
		if !ok {
			log.Panicln("failed to convert pipeline's data to bson.A")
		}
		count, ok := results["count"].(bson.A)
		if !ok {
			log.Panicln("failed to convert pipeline's data to bson.A")
		}

		var docFound int64
		if len(count) > 0 {
			var totalDocs TotalCount
			utils.Doconveter(count[0], &totalDocs)
			docFound = totalDocs.Count
		}
		return utils.Map{"data": data, "count": docFound}, nil
	} else {
		var results []bson.M
		if err = cursor.All(ctx, &results); err != nil {
			return nil, err
		}
		cursor.Close(ctx)
		return utils.Map{"data": results}, nil
	}
}

func FindByID(col *mongo.Collection, id primitive.ObjectID) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"_id": bson.M{"$eq": id}}
	var data interface{}
	err := Company.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FindByMaps(col *mongo.Collection, filter bson.D) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var data []bson.M
	curs, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := curs.All(ctx, &data); err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func FindCollReturnArray(model *mongo.Collection, filter bson.A) (utils.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := model.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	cursor.Close(ctx)
	return utils.Map{"data": results}, nil
}

func PD(key string, v interface{}) (filter bson.D) {
	return bson.D{{Key: key, Value: v}}
}
