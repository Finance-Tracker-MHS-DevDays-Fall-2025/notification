package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/generated/notification/api"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/managers"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/request"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/event"
	"github.com/sirupsen/logrus"
)

const (
	notificationEventProcessingTimeout = time.Second * 5
)

type NotificationEventHandler interface {
	event.Handler
}

type notificationEventHandler struct {
	notificationManager managers.NotificationManager
	logger              *logrus.Logger
}

func ProvideNotificationEventHandler(
	notificationManager managers.NotificationManager,
	logger *logrus.Logger,
) NotificationEventHandler {
	return &notificationEventHandler{
		notificationManager: notificationManager,
		logger:              logger,
	}
}

func (handler *notificationEventHandler) Handle(wrapper event.Wrapper) {
	ctx, cancel := context.WithTimeout(
		request.CreateContext(context.Background(), wrapper.RequestID),
		notificationEventProcessingTimeout,
	)

	defer cancel()

	handler.logger.WithContext(ctx).Infof("start handling notification event")

	var notification api.SendNotificationRequest
	err := json.Unmarshal(wrapper.RawEvent, &notification)
	if err != nil {
		handler.logger.WithContext(ctx).Errorf("invalid notification event format, error: %+v", err)
	}

	err = handler.notificationManager.SendNotification(ctx, &notification)
	if err != nil {
		handler.logger.WithContext(ctx).Errorf("failed to send notification, error: %+v", err)
	}
}
