package httpsrv

import (
	middlewarepkg "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/httpsrv/middleware"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/httpsrv"
	"github.com/labstack/echo/v4"
)

const (
	baseGroupPrefix = "/v0/notification-service"
)

func ProvideHttpServerSetupFunc(
	middleware middlewarepkg.Params,
) httpsrv.HttpServerSetupFunc {
	return func(e *echo.Echo) {
		registerRoutes(e, middleware)
	}
}

func registerRoutes(
	e *echo.Echo,
	middleware middlewarepkg.Params,
) {
	base := e.Group(baseGroupPrefix)
	base.Use(middleware.PanicHandler)
	base.Use(middleware.ErrorHandler)
	base.Use(middleware.CORS)
	base.GET(httpsrv.HealthCheckPath, httpsrv.HealthCheckFunc)
	base.Use(middleware.RequestIDHandler)
}
