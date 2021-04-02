package fun

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
)

func init() {
	cm := commands.Manager()

	uc := &commands.Command{
		Name:        "decide",
		Description: "Decide between provided options",
		Run:         RunDecide,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(uc)
}

func RunDecide(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	if len(args) < 3 {
		response := fmt.Sprintf("<@%s> You must provide at least 2 options", m.Author.ID)
		_, err = s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	options := args[1:]

	rand.Seed(time.Now().UnixNano())

	decided := rand.Intn(len(options))
	response := fmt.Sprintf("<@%s> You should go with `%s`", m.Author.ID, options[decided])

	_, err = s.ChannelMessageSend(m.ChannelID, response)

	return
}
