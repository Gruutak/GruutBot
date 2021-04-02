package fun

import (
	_ "embed"
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

	fileUrl := "https://raw.githubusercontent.com/gruutak/gruutbot/here-we-go-again/commands/fun/coin/"

	if sides[decided] == "heads" {
		fileUrl = fileUrl + "heads.png"
	} else {
		fileUrl = fileUrl + "tails.png"
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Embeds: []*discordgo.MessageEmbed{{
				Type: discordgo.EmbedTypeImage,
				Image: &discordgo.MessageEmbedImage{
					URL: fileUrl,
				},
			}},
		},
	})

	return
}
