package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	Price      float64            `bson:"price,omitempty" json:"price,omitempty"`
	Src        string             `bson:"src,omitempty" json:"src,omitempty"`
	Ean        string             `bson:"ean,omitempty" json:"ean,omitempty"`
	Categories Categories         `bson:"categories,omitempty" json:"categories,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt  time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	Active     bool               `bson:"active,omitempty" json:"active,omitempty"`
}

type Products []*Product
