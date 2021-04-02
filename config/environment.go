package config

import "github.com/spf13/viper"

func setupDefaults() {
	// Bot
	viper.SetDefault(NAME, "GruutBot")
	viper.SetDefault(TOKEN, "")
	viper.SetDefault(PREFIX, "gruut")
	viper.SetDefault(COMMANDS_PER_GUILD, false)

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
	viper.BindEnv(COMMANDS_PER_GUILD)

	// Log
	viper.BindEnv(LOG_LEVEL)
}
