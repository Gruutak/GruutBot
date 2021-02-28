package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setupLogrus() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	logLevel := logLevel()

	log.SetLevel(logLevel)
	log.Info("Log level set to ", logLevel)
}

func logLevel() (ll log.Level) {
	level := strings.ToLower(viper.GetString(LOG_LEVEL))

	ll = log.InfoLevel

	switch level {
	case "trace":
		ll = log.TraceLevel
	case "debug":
		ll = log.DebugLevel
	case "warn":
		ll = log.WarnLevel
	case "error":
		ll = log.ErrorLevel
	}

	return
}
