package commands

import (
	"container/list"

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
		queue:    list.New(),
		commands: make(map[string]*Command),
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

	commandString := command.Name

	c := cm.commands[commandString]

	if c != nil {
		log.Error("Name string already registered: ", commandString)
		return
	}

	cm.commands[commandString] = command
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

func (cm *CommandManager) ProcessQueue() {
	for cm.queue.Len() > 0 {
		element := cm.queue.Front()

		cm.register(element.Value.(*Command))

		cm.queue.Remove(element)
	}
}
