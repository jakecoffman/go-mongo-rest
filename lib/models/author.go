package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Author struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name *string    `json:"name" bson:"name,omitempty"`
	Born *time.Time `json:"born" bson:"born,omitempty"`

	Books []Book `json:"books" bson:"books,omitempty"`

	Created *time.Time `json:"created" bson:"created,omitempty"`
	Updated *time.Time `json:"updated" bson:"updated,omitempty"`
}
