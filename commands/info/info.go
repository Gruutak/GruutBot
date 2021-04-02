package info

import (
	"bytes"
	"regexp"
	"strconv"
	"text/template"
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
	title := "About " + s.State.User.Username
	dashesRegex := regexp.MustCompile(".")
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

	templateInfo := map[string]string{
		"title":         title,
		"titleDashes":   dashesRegex.ReplaceAllString(title, "-"),
		"uptime":        time.Now().Sub(viper.GetTime(config.START_TIME)).Truncate(time.Second).String(),
		"shards":        strconv.Itoa(s.ShardCount),
		"guilds":        strconv.Itoa(guilds),
		"users":         strconv.Itoa(users),
		"channelsCount": strconv.Itoa(channelsCount),
		"prefix":        viper.GetString(config.PREFIX),
	}

	t := template.Must(template.New("").Parse(runTemplate))

	var response bytes.Buffer

	if err = t.Execute(&response, templateInfo); err != nil {
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessage,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: response.String(),
		},
	})

	return
}
