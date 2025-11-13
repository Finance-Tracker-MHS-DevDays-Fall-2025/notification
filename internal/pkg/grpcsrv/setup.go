package grpcsrv

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/generated/notification/api"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/grpcsrv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ProvideGrpcServerSetupFunc(
	notificationEventSender NotificationEventSender,
	logger *logrus.Logger,
) grpcsrv.GrpcServerSetupFunc {
	return func(server *grpc.Server) {
		api.RegisterNotificationServiceServer(
			server,
			newNotificationService(
				notificationEventSender,
				logger,
			),
		)

		reflection.Register(server)
	}
}
