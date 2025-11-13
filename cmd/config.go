package main

import (
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/config"
)

func provideAppConfig() (config.AppConfig, error) {
	return config.ProvideAppConfig()
}
