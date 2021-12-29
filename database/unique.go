package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	db *mongo.Database
}

func newCollectionIndex(db *mongo.Database) *collection {
	return &collection{db}
}

func (c *collection) createIndexCategory() {
	var err error
	_, err = c.db.Collection("categories").
		Indexes().
		CreateOne(
			context.TODO(),
			mongo.IndexModel{
				Keys: bson.D{
					{Key: "ean", Value: 1},
					{Key: "name", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *collection) createIndexProduct() {
	var err error
	_, err = c.db.Collection("products").
		Indexes().
		CreateOne(
			context.TODO(),
			mongo.IndexModel{
				Keys: bson.D{
					{Key: "name", Value: 1},
					{Key: "ean", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		)
	if err != nil {
		log.Fatalln(err)
	}
}
