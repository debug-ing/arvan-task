package repository

import (
	"github.com/debug-ing/arvan-task/discount/pkg/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChargeCodeUsageRepository struct {
	client                    *mongo.Client
	ChargeCodeUsageCollection *mongo.Collection
}

func NewChargeCodeUsageRepository(client *mongo.Client) *ChargeCodeUsageRepository {
	repo := &ChargeCodeUsageRepository{client: client}
	repo.Init()
	return repo
}

func (c *ChargeCodeUsageRepository) Init() {
	c.ChargeCodeUsageCollection = c.client.Database("code").Collection("charge_code_usage")
}

func (c *ChargeCodeUsageRepository) CreateUsage(ctx echo.Context, ChargeCode primitive.ObjectID, UserID string) (bool, error) {
	mongoCtx := ctx.Request().Context()
	chargeCode := model.ChargeCodeUsage{ChargeCode: ChargeCode, UserID: UserID}
	_, err := c.ChargeCodeUsageCollection.InsertOne(mongoCtx, chargeCode)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChargeCodeUsageRepository) GetUsage(ctx echo.Context, ChargeCode string) ([]model.ChargeCodeUsage, error) {
	var chargeCodeUsages []model.ChargeCodeUsage
	mongoCtx := ctx.Request().Context()
	cursor, err := c.ChargeCodeUsageCollection.Find(mongoCtx, primitive.M{"chargeCode": ChargeCode})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(mongoCtx)
	for cursor.Next(mongoCtx) {
		var chargeCodeUsage model.ChargeCodeUsage
		err := cursor.Decode(&chargeCodeUsage)
		if err != nil {
			return nil, err
		}
		chargeCodeUsages = append(chargeCodeUsages, chargeCodeUsage)
	}
	return chargeCodeUsages, nil
}

func (c *ChargeCodeUsageRepository) GetChargeCodeUser(ctx echo.Context, UserID string) ([]model.ChargeCodeUsage, error) {
	var chargeCodeUsages []model.ChargeCodeUsage
	mongoCtx := ctx.Request().Context()
	cursor, err := c.ChargeCodeUsageCollection.Find(mongoCtx, bson.M{"userId": UserID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(mongoCtx)

	if err := cursor.All(mongoCtx, &chargeCodeUsages); err != nil {
		return nil, err
	}

	return chargeCodeUsages, nil
}
func (c *ChargeCodeUsageRepository) CheckUserUsedCode(ctx echo.Context, UserID string, ChargeCode primitive.ObjectID) (bool, error) {
	mongoCtx := ctx.Request().Context()
	cursor, err := c.ChargeCodeUsageCollection.Find(mongoCtx, bson.M{"userId": UserID, "chargeCode": ChargeCode})
	if err != nil {
		return false, err
	}
	defer cursor.Close(mongoCtx)
	if cursor.Next(mongoCtx) {
		return true, nil
	}
	return false, nil
}
