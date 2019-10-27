package main

import (
	"github.com/gruutak/gruutbot"
)

func main() {
	gruutbot.Start(setupConfig())
}

func setupConfig() *gruutbot.Config {
	v := gruutbot.ConfigViper()

	return &gruutbot.Config{
		Logger: gruutbot.SetupLogger(v.GetString("LOG_LEVEL")),
		Viper:  v,
	}
}
