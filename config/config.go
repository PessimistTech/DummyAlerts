package config

type Config struct {
	App          AppConfig
	Interpreters map[string]InterpreterConfig
	Notifiers    map[string]NotifierConfig
}

type AppConfig struct {
}

type InterpreterConfig struct {
	Enabled bool
}

type NotifierConfig struct {
	Channels map[string]string
}

var config *Config

func SetConfig(cfg *Config) {
	config = cfg
}

func GetConfig() *Config {
	return config
}
