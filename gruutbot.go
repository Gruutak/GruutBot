package gruutbot

import (
	"gruutbot/pkg/structs"
	"strings"
)

type GruutBot struct {
	prefix string
	token  string
}

func New(ic structs.InitConfig) (g *GruutBot) {
	g = &GruutBot{}

	g.initialize(ic)

	return
}

func (g *GruutBot) initialize(ic structs.InitConfig) {
	setupViper()

	if len(strings.TrimSpace(ic.Prefix)) > 0 {
		g.prefix = ic.Prefix
	}

}

func (g *GruutBot) Start() {

}
