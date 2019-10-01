package gruutbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	token := os.Getenv("GRUUTBOT_TOKEN")

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
