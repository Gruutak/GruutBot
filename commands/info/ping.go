package info

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/gruutak/gruutbot/commands"
)

func init() {
	cm := commands.Manager()

	pc := &commands.Command{
		Name:        "ping",
		Description: "Pong!",
		Category:    commands.InfoCategory,
		Aliases:     []string{"pong"},
		Run:         RunPing,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(pc)
}

func RunPing(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	command := args[0]
	latency := s.HeartbeatLatency()

	var response string

	if command == "ping" {
		response = fmt.Sprintf("%s Pong! I mean... %s", m.Author.ID, latency.Truncate(time.Millisecond))
	}

	if command == "pong" {
		response = fmt.Sprintf("%s Ping!", m.Author.ID)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, response)

	return
}
