package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChargeCodeUsage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ChargeCode primitive.ObjectID `bson:"chargeCode,omitempty"`
	UserID     string             `bson:"userId"`
	Timestamp  time.Time          `bson:"timestamp"`
}
