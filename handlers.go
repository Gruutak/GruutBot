package gruutbot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself and that doesn't have the command prefix
	isValidCommand, command := isValidCommand(s, m)

	if !isValidCommand {
		return
	}

	if command == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if command == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func isValidCommand(s *discordgo.Session, m *discordgo.MessageCreate) (bool, string) {

	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, "~") {
		return false, m.Content
	}

	return true, strings.TrimPrefix(m.Content, "~")
}