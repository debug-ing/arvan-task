package service

import (
	"net/http"

	"github.com/debug-ing/arvan-task/discount/internal"
	"github.com/debug-ing/arvan-task/discount/internal/dto"
	"github.com/debug-ing/arvan-task/discount/pkg/repository"
	"github.com/labstack/echo/v4"
)

type ChargeCodeService struct {
	ChargeCodeRepository      *repository.ChargeCodeRepository
	ChargeCodeUsageRepository *repository.ChargeCodeUsageRepository
	WalletRepository          *repository.WalletRepository
}

type IChargeCodeService interface {
	RedeemCode(c echo.Context) error
	GetCodeUsage(c echo.Context) error
}

var _ IChargeCodeService = &ChargeCodeService{}

func NewChargeCodeService(chargeCodeRepository *repository.ChargeCodeRepository, chargeCodeUsageRepository *repository.ChargeCodeUsageRepository, walletRepository *repository.WalletRepository) *ChargeCodeService {
	return &ChargeCodeService{ChargeCodeRepository: chargeCodeRepository, ChargeCodeUsageRepository: chargeCodeUsageRepository, WalletRepository: walletRepository}
}

func (cs *ChargeCodeService) GetAllCode(c echo.Context) error {
	all, err := cs.ChargeCodeRepository.GetAllChargeCode(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, all)
}

func (cs *ChargeCodeService) CreateCode(c echo.Context) error {
	body := c.Get("user").(dto.CreateChargeCodeDto)
	_, err := cs.ChargeCodeRepository.Create(c, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponseMessage("Success create code"))
}
func (cs *ChargeCodeService) UpdateCode(c echo.Context) error {
	body := c.Get("user").(dto.CreateChargeCodeDto)
	id := c.Param("id")
	r, err := cs.ChargeCodeRepository.Update(c, id, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponse(r))
}

func (cs *ChargeCodeService) RedeemCode(c echo.Context) error {
	body := c.Get("user").(dto.RedeemCodeDto)
	item, err := cs.ChargeCodeRepository.GetByCode(c, body.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	used, err := cs.ChargeCodeUsageRepository.CheckUserUsedCode(c, body.UserId, item.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	if used {
		return c.JSON(http.StatusBadRequest, internal.ErrorResponse("This code has been used before"))
	}
	r, err := cs.ChargeCodeRepository.CountDown(c, body.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	_, err = cs.ChargeCodeUsageRepository.CreateUsage(c, r.ID, body.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	_, err = cs.WalletRepository.ChargeWallet(c, body.UserId, r.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, internal.SuccessResponseMessage("Success redeem code"))
}

func (cs *ChargeCodeService) GetCodeUsage(c echo.Context) error {
	code := c.Param("code")
	all, err := cs.ChargeCodeUsageRepository.GetUsage(c, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponse(all))
}
