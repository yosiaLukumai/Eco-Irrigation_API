// package database

// import (
// 	"TEST_SERVER/model"
// 	"context"
// 	"fmt"
// 	"log"
// 	"testing"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// func NewMongoClient() *mongo.Client {

// 	mongoclient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://scopeelius:mHpF65cCZGW3O9fz@ppm.oxcmgvj.mongodb.net/?retryWrites=true&w=majority&appName=PPM"))
// 	if err != nil {
// 		log.Fatal("Error connecting ", err)
// 	}
// 	fmt.Println("Succesfully connected to DB")
// 	err = mongoclient.Ping(context.Background(), readpref.Primary())
// 	if err != nil {
// 		log.Fatal("Error Pingigng ", err)
// 	}
// 	return mongoclient
// }
// func TestMongoOperations(t *testing.T) {
// 	mongoTestClient := NewMongoClient()
// 	defer mongoTestClient.Disconnect(context.Background())
// 	coll := mongoTestClient.Database("CompanyDB").Collection("CompanyTest")
// 	companyRepo := CompanyRepo{MongoCollection: coll}
// 	t.Run("Inserrt Employ One", func(t *testing.T) {
// 		comp := model.Company{
// 			Name:     "EWURA",
// 			Phone:    "Ewura@gmail.com",
// 			Location: "Dar-es-Salaam",
// 		}
// 		result, err := companyRepo.Insert(&comp, 1)
// 		if err != nil {
// 			t.Fatal("Error inserting in company", err)
// 		}
// 		t.Log("Operation insert Succeed", result)

// 	})
// 	t.Run("Finding Company", func(t *testing.T) {

// 		result, err := companyRepo.Find("name", "EWURA")
// 		if err != nil {
// 			t.Log("Error Finding Company", err)
// 		}
// 		t.Log("Found Company ", result)
// 	})
// 	t.Run("Finding all", func(t *testing.T) {
// 		var result []model.Company
// 		 err := companyRepo.FindAll(&result)
// 		if err != nil {
// 			t.Log("Error finding all ", err)
// 		}
// 		t.Log("Got companies ", result)
// 	})
// 	t.Run("Updating", func(t *testing.T) {
// 		comp := model.Company{
// 			Name:     "EWURA",
// 			Phone:    "Ewura@gmail.com",
// 			Location: "Arusha",
// 		}
// 		resulr, err := companyRepo.Update(&comp)
// 		if err != nil {
// 			t.Log("Error Updating company ", err)
// 		}
// 		t.Log("Updated Employee ", resulr)
// 	})
// 	t.Run("Deleting Company ", func(t *testing.T) {
// 		result, err := companyRepo.Delete("EWURA")
// 		if err != nil {
// 			t.Log("Error deleting ", err)
// 		}
// 		t.Log("Deleted companie ", result)
// 	})

// 	t.Run("Deleting all database values ", func(t *testing.T) {
// 		result, err := companyRepo.DeleteAll()
// 		if err != nil {
// 			t.Log("Error deleting all data ", err)
// 		}
// 		t.Log("Deleted all data ", result)

// 	})
// }

package database
