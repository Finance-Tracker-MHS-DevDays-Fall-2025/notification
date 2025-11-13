package errmw

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/mfx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"error_handler_middleware",
		fx.Provide(
			provideErrorHandlerMiddleware,
			mfx.ProvideConfig[appConfig](configSectionName),
		),
	)
}
