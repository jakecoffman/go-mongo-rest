package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Author struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// This is the only required field in the spec (see author_routes)
	Name string `json:"name" bson:"name"`

	// Non-primitives like time.Time must have pointers to be considered "empty"
	Born *time.Time `json:"born" bson:"born,omitempty"`

	// Slices can be null, so this can be "empty"
	Books []Book `json:"books" bson:"books,omitempty"`

	Created *time.Time `json:"created" bson:"created,omitempty"`
	Updated *time.Time `json:"updated" bson:"updated,omitempty"`
}
