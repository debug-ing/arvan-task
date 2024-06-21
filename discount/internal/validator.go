package internal

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidationMiddleware(structType func() interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			model := structType()
			if err := c.Bind(model); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request data"})
			}
			val := c.Get("Validator").(*validator.Validate)
			if err := val.Struct(model); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}
			c.Set("user", model)
			return next(c)
		}
	}
}
