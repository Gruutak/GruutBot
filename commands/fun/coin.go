package fun

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
)

//go:embed coin/tails.png
var tailsFile []byte

//go:embed coin/heads.png
var headsFile []byte

func init() {
	cm := commands.Manager()

	cc := &commands.Command{
		Name:        "coin",
		Description: "Flips a coin!",
		Run:         RunCoin,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(cc)
}

func RunCoin(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) (err error) {
	sides := []string{"heads", "tails"}

	rand.Seed(time.Now().UnixNano())
	decided := rand.Intn(len(sides))

	var file *discordgo.File

	if sides[decided] == "heads" {
		file = &discordgo.File{
			Name:   "heads.png",
			Reader: bytes.NewReader(headsFile),
		}
	} else {
		file = &discordgo.File{
			Name:   "tails.png",
			Reader: bytes.NewReader(tailsFile),
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: fmt.Sprintf("<@%s> %s", i.Member.User.ID, sides[decided]),
			Embeds:  []*discordgo.MessageEmbed{{}},
		},
	})

	data := &discordgo.MessageSend{
		Content: fmt.Sprintf("<@%s> %s", i.Member.User.ID, sides[decided]),
		File:    file,
	}

	_, err = s.ChannelMessageSendComplex(i.ChannelID, data)

	return
}
