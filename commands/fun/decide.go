package fun

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
)

func init() {
	cm := commands.Manager()

	uc := &commands.Command{
		Name:        "decide",
		Description: "Decide between two alternatives",
		Run:         RunDecide,
		Intent:      discordgo.IntentsGuildMessages,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "Alternative1",
				Description: "The first alternative to choose from",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "Alternative2",
				Description: "The second alternative to choose from",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			}},
	}

	cm.AddToRegistrationQueue(uc)
}

func RunDecide(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) (err error) {
	rand.Seed(time.Now().UnixNano())

	decided := rand.Intn(len(options))
	response := fmt.Sprintf("You should go with alternative `%s`", strings.TrimSpace(options[decided].StringValue()))

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: response,
		},
	})

	return
}
