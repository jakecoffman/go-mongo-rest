package models

type Book struct {
	Title string `json:"title" bson:"title,omitempty"`
	Genre string `json:"genre" bson:"genre,omitempty"`
}
