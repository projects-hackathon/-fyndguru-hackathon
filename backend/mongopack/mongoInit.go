package mongopack

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	MongoDb *mongo.Database
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.43.41:27017"))
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://adminpass:adminpass@cluster0.d1egb.mongodb.net/fyndguru?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Mongo initialised")
	MongoDb = client.Database("fyndguru")

}
