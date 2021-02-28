package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gruutak/gruutbot/commands"
	"github.com/gruutak/gruutbot/config"

	_ "github.com/gruutak/gruutbot/commands/info"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := m.Content
	content = strings.TrimSpace(content)

	if !strings.HasPrefix(content, viper.GetString(config.PREFIX)) {
		return
	}

	log.Debug("Received message ", content)

	content = strings.TrimPrefix(content, viper.GetString(config.PREFIX))

	args := strings.Fields(content)

	cm := commands.Manager()

	command := cm.Command(args[0])

	if err := command.Run(s, m, args...); err != nil {
		log.Error(err)
	}

}
