package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func ProvideBot(
	config BotConfig,
) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bot, nil
}
