package commands

import (
	log "github.com/sirupsen/logrus"
)

var cm *CommandManager

func init() {
	categories := emptyCategories()

	categoryMap := map[CategoryType]*Category{}

	for _, c := range categories {
		categoryMap[c.Type] = c
	}

	cm = &CommandManager{
		queue:      []*Command{},
		commands:   make(map[string]*Command),
		aliases:    make(map[string]string),
		categories: categoryMap,
	}
}

func Manager() *CommandManager {
	return cm
}

func (cm *CommandManager) AddToRegistrationQueue(command *Command) {
	cm.queue = append(cm.queue, command)
}
func (cm *CommandManager) Register(command *Command) {
	commandString := command.Name

	c := cm.commands[commandString]

	if c != nil {
		log.Error("Name string already registered: ", commandString)
		return
	}

	cm.commands[commandString] = command
	cm.Intent = cm.Intent | command.Intent

	if err := cm.categories[command.Category].AddCommand(command); err != nil {
		log.Error(err)
		return
	}

	log.Info("Registered command ", command.Name)

	cm.registerAliases(command)
}

func (cm *CommandManager) registerAliases(command *Command) {
	commandString := command.Name

	cm.aliases[commandString] = commandString

	for _, commandAlias := range command.Aliases {
		c := cm.aliases[commandAlias]

		if len(c) > 0 {
			log.Error("Name aliases already registered: ", commandAlias)
			continue
		}

		cm.aliases[commandAlias] = commandString
	}
}

func (cm *CommandManager) Command(alias string) (command *Command) {
	commandString := cm.aliases[alias]

	if len(commandString) > 0 {
		command = cm.commands[commandString]
	}

	return
}

func (cm *CommandManager) Categories() map[CategoryType]*Category {
	return cm.categories
}

func (cm *CommandManager) Commands() map[string]*Command {
	return cm.commands
}

func (cm *CommandManager) ProcessQueue() {
	for _, c := range cm.queue {
		cm.Register(c)
	}
}
