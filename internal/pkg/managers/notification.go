package managers

import (
	"context"
	"strconv"

	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/internal/pkg/generated/notification/api"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type NotificationManager interface {
	SendNotification(ctx context.Context, request *api.SendNotificationRequest) error
}

type notificationManager struct {
	bot *tgbotapi.BotAPI
}

func ProvideNotificationManager(
	bot *tgbotapi.BotAPI,
) NotificationManager {
	return &notificationManager{
		bot: bot,
	}
}

func (mgr *notificationManager) SendNotification(_ context.Context, request *api.SendNotificationRequest) error {
	chatID, err := strconv.ParseInt(request.UserId, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	msg := tgbotapi.NewMessage(chatID, request.Message)

	_, err = mgr.bot.Send(msg)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
