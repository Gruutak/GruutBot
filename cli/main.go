package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bwmarrin/discordgo"

	"github.com/gruutak/gruutbot/commands"
	"github.com/gruutak/gruutbot/config"
	"github.com/gruutak/gruutbot/handlers"

	_ "github.com/gruutak/gruutbot/commands/fun"
	_ "github.com/gruutak/gruutbot/commands/info"
)

func main() {
	config.Initialize()

	log.Info("Starting Bot :)")

	dg, err := discordgo.New("Bot " + viper.GetString(config.TOKEN))
	if err != nil {
		log.Fatal("Error creating Discord session, ", err)
	}

	handlers.InjectHandlers(dg)

	cm := commands.Manager()
	cm.ProcessQueue()

	dg.Identify.Intents = cm.Intent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection, ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	if err = dg.Close(); err != nil {
		log.Error("Error closing connection, ", err)
	}
}
