package gruutbot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

const (
	BotId = 123456
)

func mockSession() *discordgo.Session {
	dg, _ := discordgo.New()
	dg.State.User = &discordgo.User{
		ID:            strconv.Itoa(BotId),
		Email:         "gruutbot@mock.com",
		Username:      "gruutbot",
		Discriminator: strconv.Itoa(1000),
		Bot:           true,
	}
	dg.State.User.ID = strconv.Itoa(BotId)

	return dg
}

func mockAuthor(authorId int) *discordgo.Member {
	name := "member"
	var roles []string

	return &discordgo.Member{
		GuildID: "1",
		Nick:    name + ":" + strconv.Itoa(authorId),
		User: &discordgo.User{
			ID:            strconv.Itoa(authorId),
			Email:         name + "@mock.com",
			Username:      name,
			Discriminator: strconv.Itoa(1001),
			Bot:           false,
		},
		Roles: roles,
	}
}

func MockMessage(content string, authorId int) *discordgo.MessageCreate {
	return &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "123",
			ChannelID: "1",
			Content:   content,
			Timestamp: discordgo.Timestamp(time.Unix(10000, 0).Format(time.RFC3339)),
			Author:    mockAuthor(authorId).User,
		},
	}
}

func TestMessageCreate(t *testing.T) {
	gviper = viper.New()
	gviper.Set("PREFIX", "~")
	mockedSession := mockSession()

	message := "ping"
	authorId := 1
	v, c := isValidCommand(mockedSession, MockMessage(message, authorId))
	assert.Equal(t, false, v)
	assert.Equal(t, message, c)

	message = "~ping"
	v, c = isValidCommand(mockedSession, MockMessage(message, authorId))
	assert.Equal(t, true, v)
	assert.Equal(t, "ping", c)

	message = "ping"
	v, c = isValidCommand(mockedSession, MockMessage(message, BotId))
	assert.Equal(t, false, v)
	assert.Equal(t, message, c)

	message = "~ping"
	v, c = isValidCommand(mockedSession, MockMessage(message, BotId))
	assert.Equal(t, false, v)
	assert.Equal(t, message, c)

	message = "^ping"
	v, c = isValidCommand(mockedSession, MockMessage(message, authorId))
	assert.Equal(t, false, v)
	assert.Equal(t, message, c)

	message = "^ping"
	v, c = isValidCommand(mockedSession, MockMessage(message, BotId))
	assert.Equal(t, false, v)
	assert.Equal(t, message, c)
}
