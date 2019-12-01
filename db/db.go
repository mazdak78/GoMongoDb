package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client
}



type Hero struct {
	Id primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
	Age int   `json:"age"`
}


func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Hero {
	var heroes []*Hero


	collection := client.Database("civilact").Collection("heroes")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero Hero
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

func ReturnOneHero(client *mongo.Client, filter bson.M) Hero {
	var hero Hero
	collection := client.Database("civilact").Collection("heroes")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func InsertNewHero(client *mongo.Client, hero Hero) interface{} {
	collection := client.Database("civilact").Collection("heroes")
	insertResult, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	return insertResult.InsertedID
}

func RemoveOneHero(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heroes")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Hero", err)
	}
	return deleteResult.DeletedCount
}

func UpdateHero(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heroes")
	atualizacao := bson.D{ {Key: "$set", Value: updatedData} }
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}
	return updatedResult.ModifiedCount
}
