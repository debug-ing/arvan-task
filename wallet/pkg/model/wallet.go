package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"userId"`
	Balance float32            `bson:"balance"`
}
