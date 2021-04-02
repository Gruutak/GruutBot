package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
	"github.com/gruutak/gruutbot/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	mu.Lock()
	handlers = append(handlers, messageCreate)
	mu.Unlock()
}

func messageCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandName := strings.TrimPrefix(i.Data.Name, viper.GetString(config.PREFIX))

	log.Debug("Received command ", i.Data.Name)

	options := i.Data.Options

	cm := commands.Manager()

	command := cm.Command(commandName)

	if command == nil {
		return
	}

	if err := command.Run(s, i, options); err != nil {
		log.Error(err)
	}

}
