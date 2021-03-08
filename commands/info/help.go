package info

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
	"github.com/gruutak/gruutbot/config"
	"github.com/spf13/viper"
)

func init() {
	cm := commands.Manager()

	hc := &commands.Command{
		Name:        "help",
		Description: "Shows this list of commands",
		Category:    commands.InfoCategory,
		Aliases:     []string{},
		Run:         RunHelp,
		Intent:      discordgo.IntentsGuildMessages,
	}

	cm.AddToRegistrationQueue(hc)
}

const commandTemplate = `
* {{.prefix}}{{.name}} {{.argsFormat}}
> {{.description}}
`

const helpTemplate = `{{$blockQuote := "` + "```" + `"}}
{{$blockQuote}}md
{{.}}
{{$blockQuote}}`

func RunHelp(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	cm := commands.Manager()

	var fields []*discordgo.MessageEmbedField

	isCommandSpecific := len(args) > 1

	var message *discordgo.MessageEmbed

	if isCommandSpecific {
		message, err = commandHelp(cm, args...)

		if message == nil {
			return
		}

	} else {
		fields, err = helpResponse(cm)

		message = &discordgo.MessageEmbed{
			Color:  3447003,
			Type:   discordgo.EmbedTypeArticle,
			Title:  "Help",
			Fields: fields,
			Description: "This is the list of commands that are currently available to use. " +
				"If you need help with a specific command, you can use `" + viper.GetString(config.PREFIX) +
				"help <command>` to find out more about that command.",
		}
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, message)

	return

}

func helpResponse(cm *commands.CommandManager) (response []*discordgo.MessageEmbedField, err error) {
	categories := cm.Categories()

	var categoriesSlice []*commands.Category

	for _, v := range categories {
		if len(v.Commands) == 0 {
			continue
		}
		categoriesSlice = append(categoriesSlice, v)
	}

	sort.Slice(categoriesSlice, func(i, j int) bool {
		return len(categoriesSlice[i].Commands) > len(categoriesSlice[j].Commands)
	})

	for _, v := range categoriesSlice {
		cmdText := ""
		t := template.Must(template.New("").Parse(helpTemplate))

		for _, c := range v.Commands {
			data := map[string]interface{}{
				"prefix":      viper.GetString(config.PREFIX),
				"name":        c.Name,
				"argsFormat":  c.ArgsFormat,
				"description": c.Description,
			}

			t := template.Must(template.New("").Parse(commandTemplate))

			var r bytes.Buffer

			if err = t.Execute(&r, data); err != nil {
				return
			}

			cmdText = cmdText + r.String()
		}

		var r bytes.Buffer

		if err = t.Execute(&r, cmdText); err != nil {
			return
		}

		field := &discordgo.MessageEmbedField{
			Name:   v.Name,
			Inline: true,
			Value:  r.String(),
		}

		response = append(response, field)
	}

	return
}

func commandHelp(cm *commands.CommandManager, args ...string) (response *discordgo.MessageEmbed, err error) {
	command := cm.Command(args[1])

	if command == nil {
		return
	}

	var message string

	if len(command.Help) > 0 {
		message = command.Help
	} else {
		message = command.Description
	}

	var fields []*discordgo.MessageEmbedField

	if len(command.ArgsFormat) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Usage",
			Inline: false,
			Value:  fmt.Sprintf("%s%s %s", viper.GetString(config.PREFIX), command.Name, command.ArgsFormat),
		})
	}

	response = &discordgo.MessageEmbed{
		Color:       3447003,
		Type:        discordgo.EmbedTypeArticle,
		Title:       "Help " + viper.GetString(config.PREFIX) + command.Name,
		Description: message,
		Fields:      fields,
	}

	return
}
