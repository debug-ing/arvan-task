package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Amount    float32            `bson:"amount"`
	Timestamp time.Time          `bson:"timestamp"`
}
