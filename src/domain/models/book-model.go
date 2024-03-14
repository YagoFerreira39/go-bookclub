package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookModel struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	Author    string              `bson:"author"`
	ISBN      string              `bson:"isbn"`
	Published float64             `bson:"published"`
}
