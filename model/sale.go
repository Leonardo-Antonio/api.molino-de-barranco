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
	sale struct {
		collection *mongo.Collection
	}

	Isale interface {
		Create(_sale *entity.Sale) (*mongo.InsertOneResult, error)
		Update(_sale *entity.Sale) (*mongo.UpdateResult, error)
		FindAll(_status bool) (entity.Sales, error)
		DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error)
		FindById(_id primitive.ObjectID) (entity.Sale, error)
	}
)

func NewSale(_db *mongo.Database) *sale {
	return &sale{collection: _db.Collection("sales")}
}

func (s *sale) Create(_sale *entity.Sale) (*mongo.InsertOneResult, error) {
	_sale.Active = true
	_sale.CreatedAt = time.Now()
	_sale.Status = "pedido"
	result, err := s.collection.InsertOne(context.TODO(), _sale)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sale) Update(_sale *entity.Sale) (*mongo.UpdateResult, error) {
	_sale.UpdatedAt = time.Now()
	update := bson.M{
		"$set": _sale,
	}
	result, err := s.collection.UpdateByID(context.TODO(), _sale.Id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sale) FindAll(_status bool) (entity.Sales, error) {
	cursor, err := s.collection.Find(context.TODO(), bson.M{
		"active": _status,
	})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	sales := new(entity.Sales)
	if err := cursor.All(context.TODO(), sales); err != nil {
		return nil, err
	}

	return *sales, nil
}

func (s *sale) FindById(_id primitive.ObjectID) (entity.Sale, error) {
	order := new(entity.Sale)
	if err := s.collection.FindOne(
		context.TODO(),
		bson.M{
			"_id":    _id,
			"active": true,
		}).Decode(&order); err != nil {
		return entity.Sale{}, err
	}

	return *order, nil
}

func (s *sale) DeleteById(_id primitive.ObjectID) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"active":     false,
			"status":     "finalizida",
		},
	}
	result, err := s.collection.UpdateByID(context.TODO(), _id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
