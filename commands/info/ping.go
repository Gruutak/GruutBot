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
		Run:         RunPing,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(pc)
}

func RunPing(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) (err error) {
	latency := s.HeartbeatLatency()

	response := fmt.Sprintf("Pong! I mean... %s", latency.Truncate(time.Millisecond))

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: response,
		},
	})

	return
}
