package grpcsrv

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/generated/notification/api"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/event"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/event/sender"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type NotificationEventSender struct {
	fx.In

	Sender sender.EventSender `name:"notification_event_sender"`
}

type service struct {
	api.UnimplementedNotificationServiceServer

	notificationEventSender sender.EventSender
	logger                  *logrus.Logger
}

func (service *service) SendNotification(_ context.Context, req *api.SendNotificationRequest) (*api.SendNotificationResponse, error) {
	rawEvent, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	wrapper := event.Wrapper{
		RawEvent:  rawEvent,
		RequestID: uuid.New(),
		Timestamp: time.Now(),
	}

	err = service.notificationEventSender.SendEvent(wrapper)
	if err != nil {
		return nil, err
	}

	return &api.SendNotificationResponse{}, nil
}

func newNotificationService(
	notificationEventSender NotificationEventSender,
	logger *logrus.Logger,
) *service {
	return &service{
		notificationEventSender: notificationEventSender.Sender,
		logger:                  logger,
	}
}
