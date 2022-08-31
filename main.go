package main

import (
	"DummyAlerts/api"
	"DummyAlerts/messages"
	"DummyAlerts/notifiers"
)

func main() {

	webApi := api.NewApi()

	// TODO: remove below discord test before merge/real use
	discord := notifiers.DiscordNotifier{}
	discord.InitDiscordClient()
	discord.Notify(&messages.Message{})

	webApi.Run(":8080")
}
