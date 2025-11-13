package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grpcsetup "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/grpcsrv"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/handlers"
	httpsetup "github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/httpsrv"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/httpsrv/middleware"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/jobs"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/managers"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/receivers"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/senders"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/tg"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/grpcsrv"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/httpsrv"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/logging"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/mfx"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func App() *fx.App {
	return fx.New(
		fx.Provide(
			provideAppConfig,
		),
		fx.Options(
			httpsrv.Module(),
			grpcsrv.Module(),
			jobs.Module(),
			logging.Module(),
			middleware.Module(),
		),
		fx.Provide(
			httpsetup.ProvideHttpServerSetupFunc,
			grpcsetup.ProvideGrpcServerSetupFunc,
			receivers.ProvideNotificationEventReceiver,
			mfx.ProvideConfig[receivers.NotificationEventReceiverConfig](receivers.NotificationEventReceiverConfigSectionName),
			senders.ProvideNotificationEventSender,
			mfx.ProvideConfig[senders.NotificationEventSenderConfig](senders.NotificationEventSenderConfigSectionName),
			handlers.ProvideNotificationEventHandler,
			managers.ProvideNotificationManager,
			mfx.ProvideConfig[tg.BotConfig](tg.BotConfigSectionName),
			tg.ProvideBot,
		),
		fx.Invoke(
			func(jobs jobs.Params, logger *logrus.Logger) {
				ctx, cancel := context.WithCancel(context.Background())

				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

				go func() {
					<-sigChan
					cancel()
				}()

				for _, job := range jobs.List() {
					go func() {
						err := job.Run(ctx)
						if err != nil {
							logger.Fatalf("failed to start background job: %+v", err)
						}
					}()
				}
			},
			func(srv grpcsrv.GrpcServer, logger *logrus.Logger) error {
				go func() {
					err := srv.Run()
					if err != nil {
						logger.Fatalf("failed to start grpc server: %+v", err)
					}
				}()

				return nil
			},
			func(srv httpsrv.HttpServer) error {
				return srv.Run()
			},
		),
	)
}
