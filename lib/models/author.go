package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Author struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// This is the only required field in the spec (see author_routes) so no need for omitempty
	Name string `json:"name" bson:"name"`

	// Non-primitives like time.Time must have pointers to be considered "empty"
	Born *time.Time `json:"born,omitempty" bson:"born,omitempty"`

	// Slices can be nil, so this can be "empty"
	Books []Book `json:"books,omitempty" bson:"books,omitempty"`

	// These values are always set by the server, so they're never not set
	Created time.Time `json:"created" bson:"created"`
	Updated time.Time `json:"updated" bson:"updated"`
}
