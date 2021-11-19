package models

type Book struct {
	Title *string `json:"title" bson:"title"`
	Genre *string `json:"genre" bson:"genre"`
}
