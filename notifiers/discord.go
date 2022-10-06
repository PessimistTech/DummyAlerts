package notifiers

import (
	"DummyAlerts/config"
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

var discord *DiscordNotifier

func (dn *DiscordNotifier) InitDiscordClient() error {
	token := os.Getenv(DISCORD_TOKEN)
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return err
	}
	dn.client = session
	return nil
}

func GetDiscordNotifier() (*DiscordNotifier, error) {
	if discord == nil {
		newNotifier := DiscordNotifier{}
		err := newNotifier.InitDiscordClient()
		if err != nil {
			return nil, err
		}
		discord = &newNotifier
	}

	return discord, nil
}

func (d *DiscordNotifier) Notify(msg *messages.Message) error {
	cfg := config.GetConfig()
	notifierCfg, ok := cfg.Notifiers["discord"]
	if !ok {
		return fmt.Errorf("notifier not configured")
	}
	val, ok := notifierCfg.Levels[msg.Level]
	if !val || !ok {
		return nil
	}
	channel, ok := notifierCfg.Channels["primary"]
	if !ok {
		return fmt.Errorf("no primary channel found for discord notifier")
	}
	_, err := d.client.ChannelMessageSend(channel, fmt.Sprintf("%s\n%s", msg.Title, msg.Content))
	return err
}
