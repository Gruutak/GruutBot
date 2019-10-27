package gruutbot

import (
	"github.com/sirupsen/logrus"
	v "github.com/spf13/viper"
)

var configKeys = map[string]string{
	"PREFIX":    "^",
	"LOG_LEVEL": "INFO",
	"BOT_TOKEN": "",
}

var viper *v.Viper

type Config struct {
	Logger *logrus.Logger
	Viper  *v.Viper
}

func ConfigViper() *v.Viper {
	vip := v.New()

	setViperDefaults(vip)
	setViperEnv(vip)
	setViperFiles(vip)

	return vip
}

func setViperDefaults(vip *v.Viper) {
	for key, value := range configKeys {
		vip.SetDefault(key, value)
	}
}

func setViperEnv(vip *v.Viper) {
	vip.SetEnvPrefix("gruutbot")
	vip.AutomaticEnv()

	for key := range configKeys {
		if err := vip.BindEnv(key); err != nil {
			logrus.Errorln("Error binding environment variable.", err)
		}
	}
}

func setViperFiles(vip *v.Viper) {
	vip.SetConfigName("config")
	vip.AddConfigPath("/etc/gruutbot/")
	vip.AddConfigPath("$HOME/.gruutbot")
	vip.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigFileNotFoundError); ok {
			logrus.Warnln("No config files found.")
		} else {
			logrus.Debugln(err)
			logrus.Errorln("Config files were found, but couldn't be read.")
		}
	}
}
