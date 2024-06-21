package repository

import (
	"context"
	"testing"

	"github.com/debug-ing/arvan-task/wallet/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTransactionRepository یک ساختار شبیه‌سازی شده برای TransactionRepository است
type MockTransactionRepository struct {
	mockCreate func(ctx context.Context, userID string, amount float32) (*model.Transaction, error)
	mockGet    func(ctx context.Context, userID string) ([]*model.Transaction, error)
}

// پیاده‌سازی رابط ITransactionRepository برای MockTransactionRepository
func (m *MockTransactionRepository) Create(ctx context.Context, userID string, amount float32) (*model.Transaction, error) {
	return m.mockCreate(ctx, userID, amount)
}

func (m *MockTransactionRepository) Get(ctx context.Context, userID string) ([]*model.Transaction, error) {
	return m.mockGet(ctx, userID)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := &MockTransactionRepository{
		mockCreate: func(ctx context.Context, userID string, amount float32) (*model.Transaction, error) {
			return &model.Transaction{
				ID:     primitive.NewObjectID(),
				UserID: userID,
				Amount: amount,
			}, nil
		},
	}

	userID := "123"
	amount := float32(100.0)
	tx, err := mockRepo.Create(context.Background(), userID, amount)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tx == nil {
		t.Fatal("expected transaction, got nil")
	}
	if tx.UserID != userID {
		t.Errorf("expected userID %s, got %s", userID, tx.UserID)
	}
	if tx.Amount != amount {
		t.Errorf("expected amount %f, got %f", amount, tx.Amount)
	}
	if tx.ID == primitive.NilObjectID {
		t.Error("expected valid ObjectID, got NilObjectID")
	}
}

func TestGetTransactions(t *testing.T) {
	mockRepo := &MockTransactionRepository{
		mockGet: func(ctx context.Context, userID string) ([]*model.Transaction, error) {
			return []*model.Transaction{
				{ID: primitive.NewObjectID(), UserID: userID, Amount: 50.0},
				{ID: primitive.NewObjectID(), UserID: userID, Amount: 75.0},
			}, nil
		},
	}

	userID := "123"
	txs, err := mockRepo.Get(context.Background(), userID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(txs) != 2 {
		t.Fatalf("expected 2 transactions, got %d", len(txs))
	}
	expectedAmounts := []float32{50.0, 75.0}
	for i, tx := range txs {
		if tx.Amount != expectedAmounts[i] {
			t.Errorf("expected amount %f, got %f", expectedAmounts[i], tx.Amount)
		}
	}
}
