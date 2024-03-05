package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBset() *mongo.Client{

	client, err := mongo.NewCLient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil{
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()


	err = client.Connect(ctx)

	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)

	if err != nil{
		log.Println("failed to connect to mongodb :(")
		return nil
	}

	fmt.Println("successfully connected to mongodb")
	return client

}

     var Client *mongo.Client = DBset()

func UserData(client *mongo.client, collectionName string)*mongo.Collection{

	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection



}

func ProductData(client *mongo.Client,collectionName string) *mongo.Collection{

	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection

}
