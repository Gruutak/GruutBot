package main

import (
	"github.com/gruutak/gruutbot"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	gruutbot.Start(setupConfig())
}

func setupConfig() *gruutbot.GruutbotConfig {
	return &gruutbot.GruutbotConfig{
		Logger: setupLogger(),
		Token:  os.Getenv("GRUUTBOT_TOKEN"),
	}
}

func setupLogger() *logrus.Logger {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(logLevel("info"))
	log.Infof("Logger set to level %+v\n", logLevel("info"))

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
