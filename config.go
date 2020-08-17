package gruutbot

import "github.com/spf13/viper"

var configDefaults = map[string]string{
	"PREFIX":    "^",
	"LOG_LEVEL": "INFO",
	"BOT_TOKEN": "",
}

func setupViper() {
	setViperDefaults(viper.GetViper())
}

func setViperDefaults(v *viper.Viper) {
	for key, value := range configDefaults {
		v.SetDefault(key, value)
	}
}

func setViperEnv(v *viper.Viper) {
	v.SetEnvPrefix("gruutbot")
	v.AutomaticEnv()

	for key := range configDefaults {
		if err := v.BindEnv(key); err != nil {
			logrus.Errorln("Error binding environment variable.", err)
		}
	}
}

func setViperFiles(v *viper.Viper) {
	v.SetConfigName("config")
	v.AddConfigPath("/etc/gruutbot/")
	v.AddConfigPath("$HOME/.gruutbot")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs/")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Warnln("No config files found.")
		} else {
			logrus.Debugln(err)
			logrus.Errorln("Config files were found, but couldn't be read.")
		}
	}
}
