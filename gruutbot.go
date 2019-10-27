package gruutbot

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func Start(config *Config) {
	parseConfig(config)

	discord, err := discordgo.New("Bot " + viper.GetString("BOT_TOKEN"))

	if err != nil {
		logger.Errorln("Error creating Discord session,", err)
		return
	}

	// Register the MessageCreate func as a callback for MessageCreate events.
	discord.AddHandler(MessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		logger.Errorln("Error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	logger.Infoln("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	if err = discord.Close(); err != nil {
		logger.Errorln("Error closing connection, ", err)
	}
}

func parseConfig(config *Config) {
	if config == nil {
		config = new(Config)
	}

	viper := config.Viper
	if viper == nil {
		viper = ConfigViper()
	}

	logger = config.Logger
	if logger == nil {
		logger = SetupLogger(viper.GetString("LOG_LEVEL"))
	}
}
