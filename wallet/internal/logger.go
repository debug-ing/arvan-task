package internal

import (
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func CustomLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c) // Execute request handler
			if err != nil {
				c.Error(err)
			}

			// After request is processed
			latency := time.Since(start)
			status := c.Response().Status
			path := c.Request().URL.Path

			// Determine the color based on the status code
			var statusColor *color.Color
			if status >= 200 && status < 300 {
				statusColor = color.New(color.FgGreen) // Green for 2xx success statuses
			} else if status >= 400 && status < 500 {
				statusColor = color.New(color.FgRed) // Red for 4xx client errors
			} else if status >= 500 {
				statusColor = color.New(color.FgRed) // Red for 5xx server errors
			} else {
				statusColor = color.New(color.FgWhite) // White for other statuses
			}

			// Log with color
			// statusColor.Printf("time=%v status=%d method=%s path=%s latency=%v\n",
			// 	time.Now().Format("2006-01-02 15:04:05"), status, c.Request().Method, path, latency)
			statusColor.Printf("%s %s %s %v\n", c.Request().Method, strconv.Itoa(status), path, latency)

			return nil
		}
	}
}
