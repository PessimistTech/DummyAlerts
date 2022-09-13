package notifiers

import (
	"DummyAlerts/messages"
	"fmt"
)

type Notifier interface {
	Notify(*messages.Message) error
}

func GetNotifier(name string) (Notifier, error) {
	switch name {
	case "discord":
		discord, err := GetDiscordNotifier()
		if err != nil {
			return nil, err
		}
		return discord, nil
	}
	return nil, fmt.Errorf("no matching notifier found")
}
