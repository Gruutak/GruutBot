package info

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/commands"
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

const helpTemplate = `{{$blockQuote := "` + "```" + `"}}
{{$blockQuote}}md
{{range $val := .}}# {{$val.Name}} 
{{range $command := $val.Commands}}* {{$command.Name}} {{$command.ArgsFormat}}
> {{$command.Description}}
{{end}}
{{end}}
{{$blockQuote}}`

const commandHelpTemplate = `{{$blockQuote := "` + "```" + `"}}
{{$blockQuote}}md
{{.message}}
{{$blockQuote}}
`

func RunHelp(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	cm := commands.Manager()

	var response string

	isCommandSpecific := len(args) > 1

	if isCommandSpecific {
		response, err = commandHelp(cm, args...)
	} else {
		response, err = helpResponse(cm)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, response)

	return

}

func helpResponse(cm *commands.CommandManager) (response string, err error) {
	categories := cm.Categories()

	var categoriesSlice []*commands.Category

	for _, v := range categories {
		if len(v.Commands) == 0 {
			continue
		}
		categoriesSlice = append(categoriesSlice, v)
	}

	sort.Slice(categoriesSlice, func(i, j int) bool {
		return categoriesSlice[i].Name < categoriesSlice[j].Name
	})

	t := template.Must(template.New("").Parse(helpTemplate))

	var r bytes.Buffer

	if err = t.Execute(&r, categoriesSlice); err != nil {
		return
	}

	response = r.String()

	return
}

func commandHelp(cm *commands.CommandManager, args ...string) (response string, err error) {
	command := cm.Command(args[1])

	t := template.Must(template.New("").Parse(commandHelpTemplate))

	var r bytes.Buffer

	var message string

	if len(command.Help) > 0 {
		message = command.Help
	} else {
		message = command.Description
	}

	if err = t.Execute(&r, map[string]string{"message": message}); err != nil {
		return
	}

	response = r.String()

	return
}
