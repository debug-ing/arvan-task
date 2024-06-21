package repository

import (
	"context"

	"github.com/debug-ing/arvan-task/wallet/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	client                *mongo.Client
	TransactionCollection *mongo.Collection
}

type ITransactionRepository interface {
	Create(ctx context.Context, userID string, amount float32) (*model.Transaction, error)
	Get(ctx context.Context, userID string) ([]*model.Transaction, error)
}

var _ ITransactionRepository = &TransactionRepository{}

func NewTransactionRepository(client *mongo.Client) *TransactionRepository {
	repo := &TransactionRepository{
		client: client,
	}
	repo.Init()
	return repo
}

func (tr *TransactionRepository) Init() {
	tr.TransactionCollection = tr.client.Database("wallet").Collection("transaction")
}

func (tr *TransactionRepository) Create(ctx context.Context, userID string, amount float32) (*model.Transaction, error) {
	item := model.Transaction{
		UserID: userID,
		Amount: amount,
	}
	result, err := tr.TransactionCollection.InsertOne(ctx, item, nil)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID
	if oid, ok := insertedID.(primitive.ObjectID); ok {
		item.ID = oid
	}
	return &item, nil
}

func (tr *TransactionRepository) Get(ctx context.Context, userID string) ([]*model.Transaction, error) {
	// filter := model.Transaction{
	// 	UserID: userID,
	// }
	cursor, err := tr.TransactionCollection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	var results []*model.Transaction
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
