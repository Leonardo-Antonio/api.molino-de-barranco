package database

import (
	"context"
	"log"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	log.Println(env.MONGO_URI)
	clientOptions := options.Client().ApplyURI(env.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	db := client.Database(env.DB_NAME)
	unique := newCollectionIndex(db)
	unique.createIndexCategory()
	unique.createIndexProduct()

	return db
}
