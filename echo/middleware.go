package echolog

import (
	"github.com/Bplotka/go-httplog"
	"github.com/labstack/echo"
)

func RequestLogger(logger *httplog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			logger.RequestLogger()(c.Response(), c.Request())
			return next(c)
		}
	}
}

func ResponseLogger(logger *httplog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if err := next(c); err != nil {
				c.Error(err)
			}
			logger.ResponseLogger()(c.Response(), c.Request())
			return nil
		}
	}
}