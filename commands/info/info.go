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

const runTemplate = "```md\n" +
	"{{.title}}\n" +
	"{{.titleDashes}}\n" +
	"> Gruutbut created by Gruutak#3335\n" +
	"\n" +
	"# Uptime\n" +
	"< {{.uptime}} >\n\n" +
	"# Shards\n" +
	"< {{.shards}} >\n\n" +
	"# Guilds per shard \n" +
	"< {{.guilds}} >\n\n" +
	"# Users per shard \n" +
	"< {{.users}} >\n\n" +
	"# Channels per shard\n" +
	"< {{.channelsCount}} >\n\n" +
	"Use {{.prefix}}help to see the command list." +
	"```"

func init() {
	cm := commands.Manager()

	pc := &commands.Command{
		Name:        "info",
		Description: "Shows information about the bot",
		Category:    commands.InfoCategory,
		Aliases:     []string{},
		Run:         RunInfo,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.Register(pc)
}

func RunInfo(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	title := "About " + viper.GetString(config.NAME)
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

	_, err = s.ChannelMessageSend(m.ChannelID, response.String())

	return
}
