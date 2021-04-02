package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func init() {
	mu.Lock()
	handlers = append(handlers, ready)
	mu.Unlock()
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	guildsLen := len(m.Guilds)

	var form string

	if guildsLen == 1 {
		form = "guild"
	} else {
		form = "guilds"
	}

	log.Info(fmt.Sprintf("Bot is now ready! Currently on %d %s.", guildsLen, form))
}
