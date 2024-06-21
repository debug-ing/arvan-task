package service

import (
	"net/http"

	"github.com/debug-ing/arvan-task/wallet/internal"
	dto "github.com/debug-ing/arvan-task/wallet/internal/dto"
	"github.com/debug-ing/arvan-task/wallet/pkg/repository"
	"github.com/labstack/echo/v4"
)

type WalletService struct {
	RepoWallet      repository.IWalletRepository
	RepoTransaction repository.ITransactionRepository
}

type IWalletService interface {
	GetWallet(c echo.Context) error
	GetWalletHistory(c echo.Context) error
	ChargeWallet(c echo.Context) error
	InitWallet(c echo.Context) error
}

var _ IWalletService = &WalletService{}

func NewWalletService(repoWallet *repository.WalletRepository, repoWalletHistory *repository.TransactionRepository) *WalletService {
	return &WalletService{RepoWallet: repoWallet, RepoTransaction: repoWalletHistory}
}

func (wl *WalletService) ChargeWallet(c echo.Context) error {
	userId := c.Param("userId")
	chargeDTO := c.Get("user").(*dto.ChargeDTO)
	_, err := wl.RepoWallet.Charge(c.Request().Context(), userId, chargeDTO.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	_, err = wl.RepoTransaction.Create(c.Request().Context(), userId, chargeDTO.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponseMessage("Wallet charged successfully"))
}

func (wl *WalletService) InitWallet(c echo.Context) error {
	userId := c.Param("userId")
	_, err := wl.RepoWallet.Create(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponseMessage("Init Wallet successfully"))
}

func (wl *WalletService) GetWallet(c echo.Context) error {
	userId := c.Param("userId")
	me, err := wl.RepoWallet.Get(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponse(me))
}

func (wl *WalletService) GetWalletHistory(c echo.Context) error {
	userId := c.Param("userId")
	all, err := wl.RepoTransaction.Get(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internal.ErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, internal.SuccessResponse(all))
}
