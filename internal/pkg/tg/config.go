package tg

const BotConfigSectionName string = "tg.bot"

type BotConfig struct {
	Token string `yaml:"token"`
}
