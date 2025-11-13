package corsmw

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type CORSMiddleware struct {
	fx.Out

	Middleware echo.MiddlewareFunc `name:"cors"`
}

func provideCORSMiddleware() CORSMiddleware {
	return CORSMiddleware{
		Middleware: middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{http.MethodGet, http.MethodPost},
			},
		),
	}
}
