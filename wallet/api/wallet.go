package api

import (
	"github.com/debug-ing/arvan-task/wallet/internal"
	dto "github.com/debug-ing/arvan-task/wallet/internal/dto"
	"github.com/debug-ing/arvan-task/wallet/pkg/service"
	"github.com/labstack/echo/v4"
)

type WalletRoutes struct {
	WalletService *service.WalletService
}

func NewWalletRoutes(wl *service.WalletService) *WalletRoutes {
	return &WalletRoutes{WalletService: wl}
}

func Register(app *echo.Echo, wl *WalletRoutes) {
	app.POST("/wallet/:userId", wl.WalletService.InitWallet)
	app.PATCH("/wallet/:userId/charge", wl.WalletService.ChargeWallet, internal.ValidationMiddleware(func() interface{} { return &dto.ChargeDTO{} }))
	app.GET("/wallet/:userId", wl.WalletService.GetWallet)
	app.GET("/wallet/:userId/history", wl.WalletService.GetWalletHistory)
}
