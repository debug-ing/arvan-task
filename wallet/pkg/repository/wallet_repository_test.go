package repository

import (
	"context"
	"testing"

	"github.com/debug-ing/arvan-task/wallet/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockWalletRepository یک ساختار شبیه‌سازی شده برای WalletRepository است
type MockWalletRepository struct {
	mockCreate func(ctx context.Context, userID string) (*model.Wallet, error)
	mockCharge func(ctx context.Context, userID string, balance float32) (model.Wallet, error)
	mockGet    func(ctx context.Context, userID string) (*model.Wallet, error)
}

// پیاده‌سازی رابط IWalletRepository برای MockWalletRepository
func (m *MockWalletRepository) Create(ctx context.Context, userID string) (*model.Wallet, error) {
	return m.mockCreate(ctx, userID)
}

func (m *MockWalletRepository) Charge(ctx context.Context, userID string, balance float32) (model.Wallet, error) {
	return m.mockCharge(ctx, userID, balance)
}

func (m *MockWalletRepository) Get(ctx context.Context, userID string) (*model.Wallet, error) {
	return m.mockGet(ctx, userID)
}

func TestCreateWallet(t *testing.T) {
	mockRepo := &MockWalletRepository{
		mockCreate: func(ctx context.Context, userID string) (*model.Wallet, error) {
			return &model.Wallet{
				ID:      primitive.NewObjectID(),
				UserID:  userID,
				Balance: 0,
			}, nil
		},
	}

	userID := "123"
	wallet, err := mockRepo.Create(context.Background(), userID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if wallet == nil {
		t.Fatal("expected wallet, got nil")
	}
	if wallet.UserID != userID {
		t.Errorf("expected userID %s, got %s", userID, wallet.UserID)
	}
	if wallet.Balance != 0 {
		t.Errorf("expected balance %f, got %f", 0.0, wallet.Balance)
	}
	if wallet.ID == primitive.NilObjectID {
		t.Error("expected valid ObjectID, got NilObjectID")
	}
}

func TestChargeWallet(t *testing.T) {
	mockRepo := &MockWalletRepository{
		mockCharge: func(ctx context.Context, userID string, balance float32) (model.Wallet, error) {
			return model.Wallet{
				ID:      primitive.NewObjectID(),
				UserID:  userID,
				Balance: balance,
			}, nil
		},
	}

	userID := "123"
	balance := float32(100.0)
	wallet, err := mockRepo.Charge(context.Background(), userID, balance)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if wallet.UserID != userID {
		t.Errorf("expected userID %s, got %s", userID, wallet.UserID)
	}
	if wallet.Balance != balance {
		t.Errorf("expected balance %f, got %f", balance, wallet.Balance)
	}
}

func TestGetWallet(t *testing.T) {
	mockRepo := &MockWalletRepository{
		mockGet: func(ctx context.Context, userID string) (*model.Wallet, error) {
			return &model.Wallet{
				ID:      primitive.NewObjectID(),
				UserID:  userID,
				Balance: 50.0,
			}, nil
		},
	}

	userID := "123"
	wallet, err := mockRepo.Get(context.Background(), userID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if wallet == nil {
		t.Fatal("expected wallet, got nil")
	}
	if wallet.UserID != userID {
		t.Errorf("expected userID %s, got %s", userID, wallet.UserID)
	}
	if wallet.Balance != 50.0 {
		t.Errorf("expected balance %f, got %f", 50.0, wallet.Balance)
	}
	if wallet.ID == primitive.NilObjectID {
		t.Error("expected valid ObjectID, got NilObjectID")
	}
}
