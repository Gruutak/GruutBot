package commands

import "github.com/bwmarrin/discordgo"

type CommandManager struct {
	commands   map[string]*Command
	aliases    map[string]string
	categories map[Category]string
	Intent     discordgo.Intent
}

type Command struct {
	Name        string
	Command     string
	Description string
	Aliases     []string
	Category    Category
	Run         func(*discordgo.Session, *discordgo.MessageCreate, ...string) error
	Intent      discordgo.Intent
}
