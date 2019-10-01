package gruutbot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself and that doesn't have the command prefix
	if m.Author.ID == s.State.User.ID && !strings.HasPrefix(m.Content, "~") {
		return
	}

	command := strings.TrimPrefix(m.Content, "~")

	if command == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if command == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
