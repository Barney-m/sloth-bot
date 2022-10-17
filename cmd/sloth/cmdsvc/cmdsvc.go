package cmdsvc

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Sloth struct {
	Name string `json: "name"`
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func HelloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!sloth" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello!")
		if err != nil {
			fmt.Println(err)
		}
	}
}
