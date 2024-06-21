package dto

type ChargeDTO struct {
	Amount float32 `json:"amount" validate:"required,numeric"`
}
