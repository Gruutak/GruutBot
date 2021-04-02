package info

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
)

const permissions = "536084054"

const urlFormat = "https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%s&scope=bot"

func init() {
	cm := commands.Manager()

	ic := &commands.Command{
		Name:        "invite",
		Description: "Shows the invite link for the bot",
		Run:         RunInvite,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(ic)
}

func RunInvite(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) (err error) {

	url := fmt.Sprintf(urlFormat, s.State.User.ID, permissions)

	response := fmt.Sprintf("You can invite the bot to your guild using the following url: %s", url)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: response,
		},
	})

	return
}
