package repository

import (
	"fmt"

	"github.com/debug-ing/arvan-task/discount/internal/dto"
	"github.com/debug-ing/arvan-task/discount/pkg/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChargeCodeRepository struct {
	client               *mongo.Client
	ChargeCodeCollection *mongo.Collection
}

func NewChargeCodeRepository(client *mongo.Client) *ChargeCodeRepository {
	repo := &ChargeCodeRepository{client: client}
	repo.Init()
	return repo
}

func (c *ChargeCodeRepository) Init() {
	c.ChargeCodeCollection = c.client.Database("code").Collection("charge_code")
}

func (c *ChargeCodeRepository) Client() *mongo.Client {
	return c.client
}

func (c *ChargeCodeRepository) GetAllChargeCode(ctx echo.Context) ([]model.ChargeCode, error) {
	var chargeCodes []model.ChargeCode
	mongoCtx := ctx.Request().Context()
	cursor, err := c.ChargeCodeCollection.Find(mongoCtx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(mongoCtx)
	for cursor.Next(mongoCtx) {
		var chargeCode model.ChargeCode
		err := cursor.Decode(&chargeCode)
		if err != nil {
			return nil, err
		}
		chargeCodes = append(chargeCodes, chargeCode)
	}
	return chargeCodes, nil
}

func (c *ChargeCodeRepository) GetByCode(ctx echo.Context, code string) (model.ChargeCode, error) {
	mongoCtx := ctx.Request().Context()
	var chargeCode model.ChargeCode
	err := c.ChargeCodeCollection.FindOne(mongoCtx, bson.M{"code": code}).Decode(&chargeCode)
	if err != nil {
		return model.ChargeCode{}, err
	}
	return chargeCode, nil
}

func (c *ChargeCodeRepository) Create(ctx echo.Context, body dto.CreateChargeCodeDto) (model.ChargeCode, error) {
	mongoCtx := ctx.Request().Context()
	chargeCode := model.ChargeCode{Code: body.Code, UsageLimit: body.UsageLimit, Price: body.Price}
	_, err := c.ChargeCodeCollection.InsertOne(mongoCtx, chargeCode)
	if err != nil {
		return model.ChargeCode{}, err
	}
	return chargeCode, nil
}

func (c *ChargeCodeRepository) Update(ctx echo.Context, id string, body dto.CreateChargeCodeDto) (model.ChargeCode, error) {
	mongoCtx := ctx.Request().Context()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.ChargeCode{}, err
	}
	chargeCode := model.ChargeCode{Code: body.Code, UsageLimit: body.UsageLimit, Price: body.Price}
	result := c.ChargeCodeCollection.FindOneAndUpdate(
		mongoCtx,
		bson.M{"_id": objectID},
		bson.M{"$set": chargeCode},
	)
	if result.Err() != nil {
		return model.ChargeCode{}, err
	}
	return chargeCode, nil
}

func (c *ChargeCodeRepository) CountDown(ctx echo.Context, code string) (model.ChargeCode, error) {
	mongoCtx := ctx.Request().Context()
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var chargeCode model.ChargeCode
	result := c.ChargeCodeCollection.FindOneAndUpdate(
		mongoCtx,
		bson.M{"code": code},
		bson.M{"$inc": bson.M{"usageLimit": -1}},
		options,
	)
	if err := result.Decode(&chargeCode); err != nil {
		return model.ChargeCode{}, fmt.Errorf("error decoding result: %w", err)
	}
	return chargeCode, nil
}
