package dto

type CreateChargeCodeDto struct {
	Code       string  `json:"code" validate:"required"`
	Price      float32 `json:"price" validate:"required"`
	UsageLimit int16   `json:"usage_limit" validate:"required"`
}
