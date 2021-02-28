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
		Category:    commands.InfoCategory,
		Aliases:     []string{},
		Run:         RunInvite,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.Register(ic)
}

func RunInvite(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {

	url := fmt.Sprintf(urlFormat, s.State.User.ID, permissions)

	_, err = s.ChannelMessageSendReply(m.ChannelID, url, m.MessageReference)

	return
}
