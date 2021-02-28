package commands

import (
	log "github.com/sirupsen/logrus"
)

var cm *CommandManager

func init() {
	cm = &CommandManager{
		commands:   make(map[string]*Command),
		aliases:    make(map[string]string),
		categories: make(map[Category]string),
	}
}

func Manager() *CommandManager {
	return cm
}

func (cm *CommandManager) Register(command *Command) {
	commandString := command.Command

	c := cm.commands[commandString]

	if c != nil {
		log.Error("Command string already registered: ", commandString)
		return
	}

	cm.commands[commandString] = command
	cm.categories[command.Category] = commandString
	cm.Intent = cm.Intent | command.Intent

	cm.registerAliases(command)
}

func (cm *CommandManager) registerAliases(command *Command) {
	commandString := command.Command

	cm.aliases[commandString] = commandString

	for _, commandAlias := range command.Aliases {
		c := cm.aliases[commandAlias]

		if len(c) > 0 {
			log.Error("Command aliases already registered: ", commandAlias)
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
