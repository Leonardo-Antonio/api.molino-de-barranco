package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	category struct {
		collection *mongo.Collection
	}

	Icategory interface {
		Create(_category *entity.Category) (*mongo.InsertOneResult, error)
		Update(_categoty *entity.Category) (*mongo.UpdateResult, error)
		FindAll() (entity.Categories, error)
		DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error)
	}
)

func NewCategory(_db *mongo.Database) *category {
	return &category{collection: _db.Collection("categories")}
}

func (p *category) Create(_category *entity.Category) (*mongo.InsertOneResult, error) {
	_category.Active = true
	_category.CreatedAt = time.Now()
	result, err := p.collection.InsertOne(context.TODO(), _category)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *category) Update(_categoty *entity.Category) (*mongo.UpdateResult, error) {
	_categoty.UpdatedAt = time.Now()
	update := bson.M{
		"$set": _categoty,
	}
	result, err := p.collection.UpdateByID(context.TODO(), _categoty.Id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *category) FindAll() (entity.Categories, error) {
	cursor, err := p.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	categories := new(entity.Categories)
	if err := cursor.All(context.TODO(), categories); err != nil {
		return nil, err
	}

	return *categories, nil
}

func (p *category) DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"active":     false,
		},
	}
	result, err := p.collection.UpdateByID(context.TODO(), _id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
