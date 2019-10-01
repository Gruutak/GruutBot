package gruutbot

import "github.com/sirupsen/logrus"

type GruutbotConfig struct {
	Logger *logrus.Logger
	Token string
}