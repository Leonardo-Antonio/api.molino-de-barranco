package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Products   []Product          `bson:"products,omitempty" json:"products,omitempty"`
	InfoClient Client             `bson:"info_client,omitempty" json:"info_client,omitempty"`
	Type       string             `bson:"type,omitempty" json:"type,omitempty"`
	Nick       string             `bson:"nick,omitempty" json:"nick,omitempty" validmor:"required"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt  time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	Active     bool               `bson:"active,omitempty" json:"active,omitempty"`
}

type Sales []*Sale
