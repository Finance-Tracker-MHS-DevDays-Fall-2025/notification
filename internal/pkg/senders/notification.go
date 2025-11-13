package senders

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/config"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/event/sender"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

const (
	NotificationEventSenderConfigSectionName = "notification_event"

	notificationEventSenderName = "notification-event"
)

type NotificationEventSenderConfig struct {
	config.EventSenderConfig `yaml:",inline"`
}

type NotificationEventSender struct {
	fx.Out

	Sender sender.EventSender `name:"notification_event_sender"`
}

func ProvideNotificationEventSender(
	config NotificationEventSenderConfig,
	logger *logrus.Logger,
) (NotificationEventSender, error) {
	snd, err := sender.NewEventSender(config.EventSenderConfig, logger, notificationEventSenderName)

	return NotificationEventSender{
		Sender: snd,
	}, err
}
