package commands

import (
	"fmt"
	"sort"

	"github.com/bwmarrin/discordgo"
)

type CommandManager struct {
	commands   map[string]*Command
	aliases    map[string]string
	categories map[CategoryType]*Category
	Intent     discordgo.Intent
}

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Category    CategoryType
	Run         func(*discordgo.Session, *discordgo.MessageCreate, ...string) error
	Intent      discordgo.Intent
}

type Category struct {
	Type     CategoryType
	Name     string
	Commands []*Command
}

func (c *Category) AddCommand(command *Command) (err error) {
	for _, v := range c.Commands {
		if v.Name == command.Name {
			err = fmt.Errorf("command already exists in this category: %s", command.Name)
			return
		}
	}

	c.Commands = append(c.Commands, command)

	sort.Slice(c.Commands, func(i, j int) bool {
		return c.Commands[i].Name < c.Commands[j].Name
	})
	return
}
