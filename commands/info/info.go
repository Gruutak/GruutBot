package info

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"github.com/gruutak/gruutbot/commands"
	"github.com/gruutak/gruutbot/config"
)

const runTemplate = `{{$blockQuote := "` + "```" + `"}}
{{$blockQuote}}md
{{.title}}
{{.titleDashes}}
> This bot was created using GruutBot by Gruutak#3335

# Uptime
< {{.uptime}} >
# Shards
< {{.shards}} >
# Guilds per shard 
< {{.guilds}} >
# Users per shard 
< {{.users}} >
# Channels per shard
< {{.channelsCount}} >

Use {{.prefix}}help to see the command list.
{{$blockQuote}}`

func init() {
	cm := commands.Manager()

	ic := &commands.Command{
		Name:        "info",
		Description: "Shows information about the bot",
		Run:         RunInfo,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(ic)
}

func RunInfo(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) (err error) {
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessage,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Color:       3447003,
					Type:        discordgo.EmbedTypeRich,
					Title:       "Help",
					Fields:      infoFields(s),
					Description: "This bot was created using GruutBot by Gruutak#3335",
					Footer: &discordgo.MessageEmbedFooter{
						IconURL: s.State.User.AvatarURL(""),
						Text:    s.State.User.Username,
					},
				},
			},
		},
	})

	return
}

func infoFields(s *discordgo.Session) []*discordgo.MessageEmbedField {
	guilds := len(s.State.Guilds)
	users := 0
	channelsCount := 0

	for _, g := range s.State.Guilds {
		users += len(g.Members)

		channels, _ := s.GuildChannels(g.ID)

		for _, c := range channels {
			if c.Type == discordgo.ChannelTypeGuildText {
				channelsCount++
			}
		}
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Uptime",
			Value: time.Now().Sub(viper.GetTime(config.START_TIME)).Truncate(time.Second).String(),
		},
		{
			Name:  "Shards",
			Value: strconv.Itoa(s.ShardCount),
		},
		{
			Name:  "Guilds per shard",
			Value: strconv.Itoa(guilds),
		},
		{
			Name:  "Users per shard ",
			Value: strconv.Itoa(users),
		},
		{
			Name:  "Channels per shard",
			Value: strconv.Itoa(channelsCount),
		},
	}

	return fields
}
