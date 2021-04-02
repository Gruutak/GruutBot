package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func InjectHandlers(session *discordgo.Session) {
	log.Trace("Injecting handlers...")

	session.AddHandler(messageCreate)
	log.Trace("messageCreate injected...")

	session.AddHandler(ready)
}
