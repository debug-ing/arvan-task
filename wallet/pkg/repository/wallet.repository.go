package repository

import (
	"context"
	"fmt"

	"github.com/debug-ing/arvan-task/wallet/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository struct {
	client           *mongo.Client
	WalletCollection *mongo.Collection
}

type IWalletRepository interface {
	Create(ctx context.Context, userID string) (*model.Wallet, error)
	Charge(ctx context.Context, UserID string, balance float32) (model.Wallet, error)
	Get(ctx context.Context, userID string) (*model.Wallet, error)
}

var _ IWalletRepository = &WalletRepository{}

func NewWalletRepository(client *mongo.Client) *WalletRepository {
	repo := &WalletRepository{
		client: client,
	}
	repo.Init()
	return repo
}

func (wr *WalletRepository) Init() {
	wr.WalletCollection = wr.client.Database("wallet").Collection("wallet")
}

func (wr *WalletRepository) Create(ctx context.Context, userID string) (*model.Wallet, error) {
	fmt.Printf("Creating wallet for user %s\n", userID)
	item := model.Wallet{
		UserID:  userID,
		Balance: 0,
	}
	result, err := wr.WalletCollection.InsertOne(ctx, item, nil)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID
	if oid, ok := insertedID.(primitive.ObjectID); ok {
		item.ID = oid
	}
	return &item, nil
}

func (wr *WalletRepository) Charge(ctx context.Context, UserID string, balance float32) (model.Wallet, error) {
	result := wr.WalletCollection.FindOneAndUpdate(
		ctx,
		bson.M{"userId": UserID},
		bson.M{"$inc": bson.M{"balance": balance}},
	)
	if result.Err() != nil {
		return model.Wallet{}, result.Err()
	}
	var wallet model.Wallet
	result.Decode(wallet)
	return wallet, nil
}

func (wr *WalletRepository) Get(ctx context.Context, userID string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := wr.WalletCollection.FindOne(ctx, bson.M{"userId": userID}).Decode(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
