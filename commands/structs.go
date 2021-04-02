package commands

import (
	"container/list"
	"fmt"
	"sort"

	"github.com/bwmarrin/discordgo"
)

type CommandManager struct {
	queue    *list.List
	commands map[string]*Command
	Intent   discordgo.Intent
}

type Command struct {
	ID          string
	Name        string
	Description string
	Version     string
	Options     []*discordgo.ApplicationCommandOption
	Initialize  func(*Command) error
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
