package config

import "github.com/spf13/viper"

func setupDefaults() {
	// Bot
	viper.SetDefault(NAME, "GruutBot")
	viper.SetDefault(TOKEN, "MzE1MTc1ODYyMTQ2MTcwODgx.WR8oZQ.iR7ho5Qvdh1Js_TFf9g1cXT6Kl0")
	viper.SetDefault(PREFIX, "^")

	// Log
	viper.SetDefault(LOG_LEVEL, "info")
}

func setupEnvironment() {
	viper.SetEnvPrefix(ApplicationName)
	viper.AutomaticEnv()

	// Bot
	viper.BindEnv(NAME)
	viper.BindEnv(TOKEN)
	viper.BindEnv(PREFIX)

	// Log
	viper.BindEnv(LOG_LEVEL)
}
