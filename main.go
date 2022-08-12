package main

import "DummyAlerts/api"

func main() {

	webApi := api.NewApi()

	webApi.Run(":8080")
}
