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
	product struct {
		collection *mongo.Collection
	}

	Iproduct interface {
		Create(_product *entity.Product) (*mongo.InsertOneResult, error)
		Update(_product *entity.Product) (*mongo.UpdateResult, error)
		FindAll() (entity.Products, error)
		FindByEan(_ean string) (entity.Product, error)
		DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error)
	}
)

func NewProduct(_db *mongo.Database) *product {
	return &product{collection: _db.Collection("products")}
}

func (p *product) Create(_product *entity.Product) (*mongo.InsertOneResult, error) {
	_product.Active = true
	_product.CreatedAt = time.Now()
	result, err := p.collection.InsertOne(context.TODO(), _product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *product) Update(_product *entity.Product) (*mongo.UpdateResult, error) {
	_product.UpdatedAt = time.Now()
	update := bson.M{
		"$set": _product,
	}
	result, err := p.collection.UpdateByID(context.TODO(), _product.Id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *product) FindAll() (entity.Products, error) {
	cursor, err := p.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	products := new(entity.Products)
	if err := cursor.All(context.TODO(), products); err != nil {
		return nil, err
	}

	return *products, nil
}

func (p *product) FindByEan(_ean string) (entity.Product, error) {
	product := new(entity.Product)
	if err := p.collection.FindOne(context.TODO(), bson.M{
		"ean": _ean,
	}).Decode(product); err != nil {
		return entity.Product{}, err
	}

	return *product, nil
}

func (p *product) DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error) {
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
