package utils

import (
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TimeLocal() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now().Local())
}

func IDHex(id string) primitive.ObjectID {
	idp, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Failed to convert to ID string")
		log.Fatal(" Failed to parse to hex string...")
		// this can't be allowed to happen
		panic("ID_HEX Conversion Error")
	}
	return idp
}

func IDHexErr(id string) (primitive.ObjectID, error) {
	idp, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return idp, err
	}
	return idp, nil
}

// func FacetCreator() bson.A {
// 	return bson.A{
// 		bson.D{{"$facet", }}}
// 	}
// }

// func CountData() bson.A {
// 	return    bson.A{
// 		bson.D{{"$count", "count"}},
// 	},

// }

func FacetCreator(filter bson.A) bson.A {

	return bson.A{
		bson.D{
			{"$facet",
				bson.D{
					{"data", filter},
					{"count", DocCounter()},
				},
			},
		},
	}
}

func FacetCreatorMain(skipper ...bson.D) bson.D {
	skippPipe := mongo.Pipeline{}
	skippPipe = append(skippPipe, skipper...)
	return bson.D{{"$facet", MDs(ME("data", skippPipe), ME("count", DocCounter()))}}
}

func DocCounter() bson.A {
	return bson.A{
		bson.D{{"$count", "count"}},
	}
}

func LimiterSkipper(skip int, limit int, filter bson.A) bson.A {
	skipp := SkipValues(skip, limit)
	return append(filter, bson.D{{"$skip", skipp}}, bson.D{{"$limit", limit}})
}

func SkipValues(currentPage int, perPage int) int {
	return (currentPage - 1) * perPage
}

func LimitOnly(limit int, filter bson.A) bson.A {
	return append(filter, bson.D{{"$limit", limit}})
}

func AggregationFilter(filter ...bson.D) bson.A {
	var result = make(bson.A, 0)
	for _, value := range filter {
		result = append(result, value)
	}
	return result
}

func ME(key string, val interface{}) bson.E {
	return bson.E{Key: key, Value: val}
}

func MD(key string, val interface{}) bson.D {
	return bson.D{{Key: key, Value: val}}
}

func MDs(bsD ...bson.E) bson.D {
	var res bson.D
	res = append(res, bsD...)
	return res
}

func MA(vals ...any) bson.A {
	bsonArray := bson.A{}
	bsonArray = append(bsonArray, vals...)
	return bsonArray
}
