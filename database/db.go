package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DB       *mongo.Database
	DBClient *mongo.Client

	Client            *mongo.Collection
	User              *mongo.Collection
	EmailVerification *mongo.Collection
	Roles             *mongo.Collection
	Company           *mongo.Collection
	Branch            *mongo.Collection
	Pumps             *mongo.Collection
	Farmers           *mongo.Collection
	Package           *mongo.Collection
	Payment           *mongo.Collection
)

func InitDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error connecting to Database")
	}
	log.Println("Success Connecting the database")
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error Pinging the Database \n", err)
		panic("Failed to connect to the database...")
	}
	log.Println("===Success connected the database==")

	DB = client.Database("AFM")
	Client = DB.Collection("client")
	User = DB.Collection("user")
	Roles = DB.Collection("roles")
	EmailVerification = DB.Collection("emailVerfication")
	Company = DB.Collection("company")
	Pumps = DB.Collection("pumps")
	Farmers = DB.Collection("farmers")
	Package = DB.Collection("package")
	Payment = DB.Collection("payment")
}
