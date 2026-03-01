package metrics

import "github.com/labstack/echo/v4"

func RequestCounter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			RequestsTotal.Inc()
			return next(c)
		}
	}
}
