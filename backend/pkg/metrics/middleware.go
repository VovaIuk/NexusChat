package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func HTTPMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() == "/api/v1/metrics" {
				return next(c)
			}

			start := time.Now()
			err := next(c)

			status := c.Response().Status
			if status == 0 {
				status = http.StatusOK
			}

			path := c.Path()
			if path == "" {
				path = "unknown"
			}

			method := c.Request().Method
			statusStr := strconv.Itoa(status)

			RequestsTotal.WithLabelValues(method, path, statusStr).Inc()
			RequestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())

			return err
		}
	}
}
