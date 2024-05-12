package db

import (
	"context"
	"log"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() *mongo.Database {

	url := config.Environments().DatabaseMG.Url

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}

	log.Println("¡You successfully connected to MongoDB!")

	return client.Database("nlj-internal")

}
