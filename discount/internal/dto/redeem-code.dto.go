package dto

type RedeemCodeDto struct {
	Code   string `json:"code" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}
