package commands

import (
	"container/list"

	"github.com/bwmarrin/discordgo"
	"github.com/gruutak/gruutbot/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cm *CommandManager

func init() {
	categories := emptyCategories()

	categoryMap := map[CategoryType]*Category{}

	for _, c := range categories {
		categoryMap[c.Type] = c
	}

	cm = &CommandManager{
		queue:               list.New(),
		commands:            make(map[string]*Command),
		applicationCommands: []*discordgo.ApplicationCommand{},
		Intent:              discordgo.IntentsNone,
	}
}

func Manager() *CommandManager {
	return cm
}

func (cm *CommandManager) AddToRegistrationQueue(command *Command) {
	cm.queue.PushBack(command)
}
func (cm *CommandManager) register(command *Command) {
	if command.Initialize != nil {
		if err := command.Initialize(command); err != nil {
			log.Error(err)
			return
		}
	}

	cm.applicationCommands = append(cm.applicationCommands, &discordgo.ApplicationCommand{
		Name:        viper.GetString(config.PREFIX) + command.Name,
		Description: command.Description,
		Options:     command.Options,
	})

	cm.commands[command.Name] = command

	cm.Intent = cm.Intent | command.Intent

	log.Info("Registered command ", command.Name)
}

func (cm *CommandManager) Command(name string) (command *Command) {
	if len(name) > 0 {
		command = cm.commands[name]
	}

	return
}

func (cm *CommandManager) Commands() map[string]*Command {
	return cm.commands
}

func (cm *CommandManager) ProcessQueue(s *discordgo.Session) {
	for cm.queue.Len() > 0 {
		element := cm.queue.Front()

		cm.register(element.Value.(*Command))

		cm.queue.Remove(element)
	}

	for _, v := range cm.applicationCommands {
		var cmd *discordgo.ApplicationCommand
		var err error

		if viper.GetBool(config.COMMANDS_PER_GUILD) {
			for _, g := range s.State.Guilds {
				cmd, err = s.ApplicationCommandCreate(s.State.User.ID, g.ID, v)
			}
		} else {
			cmd, err = s.ApplicationCommandCreate(s.State.User.ID, "", v)
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Debug("Registered slash command", cmd)

		v.ID = cmd.ID
		v.Version = cmd.Version
		v.ApplicationID = cmd.ApplicationID
	}
}

func (cm *CommandManager) RemoveCommands(s *discordgo.Session) {
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range cmds {
		if viper.GetBool(config.COMMANDS_PER_GUILD) {
			for _, g := range s.State.Guilds {
				if _, err = s.ApplicationCommandCreate(s.State.User.ID, g.ID, v); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID); err != nil {
				log.Fatal(err)
			}
		}
	}
}
