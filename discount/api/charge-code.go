package api

import (
	"github.com/debug-ing/arvan-task/discount/internal"
	"github.com/debug-ing/arvan-task/discount/internal/dto"
	"github.com/debug-ing/arvan-task/discount/pkg/service"
	"github.com/labstack/echo/v4"
)

type ChargeCodeRoutes struct {
	ChargeCodeService *service.ChargeCodeService
}

func NewChargeCodeRoutes(wl *service.ChargeCodeService) *ChargeCodeRoutes {
	return &ChargeCodeRoutes{ChargeCodeService: wl}
}

func Register(app *echo.Echo, wl *ChargeCodeRoutes) {

	app.POST("/charge-code/redeem", wl.ChargeCodeService.RedeemCode, internal.ValidationMiddleware(func() interface{} { return &dto.RedeemCodeDto{} }))
	app.GET("/charge-code/:code/usage", wl.ChargeCodeService.GetCodeUsage)
	app.GET("/charge-code", wl.ChargeCodeService.GetAllCode)
	app.POST("/charge-code", wl.ChargeCodeService.CreateCode, internal.ValidationMiddleware(func() interface{} { return &dto.CreateChargeCodeDto{} }))
	app.PUT("/charge-code/:id", wl.ChargeCodeService.UpdateCode)

}
