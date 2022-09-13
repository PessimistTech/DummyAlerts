package main

import (
	"DummyAlerts/api"
	"DummyAlerts/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./.config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatalf("unable to read in config")
	}

	var cfg config.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to marshall config")
	}

	logrus.Infof("Config: %+v", cfg)

	config.SetConfig(&cfg)
}

func main() {

	webApi := api.NewApi()

	webApi.Run(":8080")
}
