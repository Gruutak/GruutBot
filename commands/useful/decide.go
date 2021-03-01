package useful

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
		Description: "Decide between options",
		ArgsFormat:  "<option1> <option2> ... <optionN>",
		Help:        "This command will pick 1 of the options you provided randomly",
		Category:    commands.UsefulCategory,
		Aliases:     []string{},
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

	min := 0
	max := len(options)

	rand.Seed(time.Now().UnixNano())

	decided := rand.Intn(max-min+1) + min
	response := fmt.Sprintf("<@%s> You should go with `%s`", m.Author.ID, options[decided])

	_, err = s.ChannelMessageSend(m.ChannelID, response)

	return
}
