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

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection, ", err)
	}

	cm := commands.Manager()
	cm.ProcessQueue(dg)

	dg.Identify.Intents = cm.Intent

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Infoln("Shutting down bot")

	cm.RemoveCommands(dg)

	err = dg.Close()
	if err != nil {
		log.Fatalln("Error closing connection", err)
	}
}
