package notifiers

import (
	"DummyAlerts/messages"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

const (
	DISCORD_TOKEN = "DISCORD_TOKEN"
)

type DiscordNotifier struct {
	client *discordgo.Session
}

func (dn *DiscordNotifier) InitDiscordClient() error {
	token := os.Getenv(DISCORD_TOKEN)
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return err
	}
	dn.client = session
	return nil
}

func (d *DiscordNotifier) Notify(msg *messages.Message) error {
	// TODO: get channel id from app config
	d.client.ChannelMessageSend("<channelID>", "Test")
	return nil
}
