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

func RunHelp(s *discordgo.Session, m *discordgo.MessageCreate, args ...string) (err error) {
	cm := commands.Manager()

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

	var response bytes.Buffer

	if err = t.Execute(&response, categoriesSlice); err != nil {
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, response.String())

	return

}
