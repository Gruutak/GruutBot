package handlers

import (
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var mu sync.Mutex

var handlers []interface{}

func InjectHandlers(session *discordgo.Session) {
	log.Trace("Injecting handlers...")

	for _, h := range handlers {
		session.AddHandler(h)
		log.Tracef("%s handler injected", nameOf(h))
	}
}

func nameOf(handler interface{}) string {
	v := reflect.ValueOf(handler)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			s := strings.Split(rf.Name(), ".")
			name := s[len(s)-1]
			return name
		}
	}
	return v.String()
}
