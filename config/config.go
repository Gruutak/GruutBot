package config

import (
	"time"

	"github.com/spf13/viper"
)

const ApplicationName = "GruutBot"

func Initialize() {
	viper.Set(START_TIME, time.Now())
	setupDefaults()
	setupEnvironment()
	setupLogrus()
}
