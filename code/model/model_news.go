package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//News a
type News struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	Title, Body, Author, Date string
	React, Count              int64
}
