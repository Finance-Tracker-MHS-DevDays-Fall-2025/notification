package middleware

import (
	corsmw "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/middleware/cors"
	errmw "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/middleware/error"
	panicmw "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/middleware/panic"
	reqidmw "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/middleware/requestid"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"middleware",
		fx.Options(
			corsmw.Module(),
			errmw.Module(),
			panicmw.Module(),
			reqidmw.Module(),
		),
	)
}

type Params struct {
	fx.In

	CORS             echo.MiddlewareFunc `name:"cors"`
	ErrorHandler     echo.MiddlewareFunc `name:"error_handler"`
	PanicHandler     echo.MiddlewareFunc `name:"panic_handler"`
	RequestIDHandler echo.MiddlewareFunc `name:"request_id_handler"`
}
