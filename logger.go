package gruutbot

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var logger *logrus.Logger

func SetupLogger(level string) *logrus.Logger {
	level = strings.ToLower(level)
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if len(level) < 1 {
		level = "info"
	}

	log.SetLevel(logLevel(level))
	log.Infof("Logger set to level %+v\n", log.Level)

	return log
}

func logLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
