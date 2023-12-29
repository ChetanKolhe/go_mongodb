package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("MongoDb connection ")

	// crednetial := options.Credential{
	// 	AuthMechanism: "SCRAM-SHA-1",
	// 	AuthSource:    "test",
	// 	Username:      "chetan",
	// 	Password:      "chetan",
	// }
	// clientOption := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(crednetial)

	// Setting up clint option
	clientOption := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	client, _ := mongo.Connect(context.TODO(), clientOption)

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection successful")

	type Actor struct {
		FirstName string
		LastName  string
		Awards    int16
	}

	actor := Actor{
		FirstName: "Chetan",
		LastName:  "Kolhe",
		Awards:    4,
	}

	collection := client.Database("dvdstore").Collection("actordetail")
	inserResult, err := collection.InsertOne(context.TODO(), actor)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inserResult)

	// Insert Manay example
	inserResults, _ := collection.InsertMany(context.TODO(), []interface{}{Actor{FirstName: "Chetan", LastName: "Kolhe2", Awards: int16(100)}})
	fmt.Println(inserResults)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
