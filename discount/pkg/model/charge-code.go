package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChargeCode struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Code       string             `bson:"code"`
	Price      float32            `bson:"price"`
	UsageLimit int16              `bson:"usageLimit"`
}
