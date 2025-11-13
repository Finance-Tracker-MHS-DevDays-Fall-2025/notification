package httpsrv

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/mfx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"httpsrv",
		fx.Provide(
			provideHttpServer,
			mfx.ProvideConfig[httpServerConfig](configSectionName),
		),
	)
}
