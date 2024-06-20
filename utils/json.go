package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AggregatorDecoder(ctx context.Context, cursor *mongo.Cursor, data any) {
	for cursor.Next(ctx) {
		err := cursor.Decode(data)
		if err != nil {
			log.Panicln(err)
		}

		cursor.Close(ctx)
	}
}

func SkipValue(currentPage int16, perPage int16) int16 {
	return (currentPage - 1) * perPage
}

func FullName(firstName string, lastname string) string {
	return firstName + "-" + lastname
}

type Map map[string]interface{}

func Msg(msg string) (response Map) {
	return Map{"msg": msg}
}

func Doconveter(from, to any) {
	bsonBytes, err := bson.Marshal(from)
	if err != nil {
		log.Panicln(err)
	}

	err = bson.Unmarshal(bsonBytes, to)
	if err != nil {
		log.Panicln(err)
	}
}
