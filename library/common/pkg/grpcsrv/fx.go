package grpcsrv

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/mfx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"grpcsrv",
		fx.Provide(
			provideGrpcServer,
			mfx.ProvideConfig[grpcServerConfig](configSectionName),
		),
	)
}
