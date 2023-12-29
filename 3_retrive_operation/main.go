package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
		FirstName: "Kunal",
		LastName:  "klsjdflkj",
		Awards:    4,
	}

	filter := bson.D{{"firstname", "Chetan"}}

	collection := client.Database("dvdstore").Collection("actordetail")

	if err != nil {
		log.Fatal(err)
	}

	collection.FindOne(context.TODO(), filter).Decode(&actor)

	fmt.Println(actor)

	// Retrieving mulitple object
	cur, _ := collection.Find(context.TODO(), filter)

	var result []Actor
	for cur.Next(context.TODO()) {
		var act Actor

		cur.Decode(&act)
		result = append(result, act)

	}
	fmt.Println(result)

	for _, value := range result {
		fmt.Println(value)
	}
	cur.Close(context.TODO())

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
