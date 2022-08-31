package config

// TODO: finish working out this config, and add yaml config file
type Config struct {
	Cfg          AppConfig
	Interpreters map[string]InterpreterConfig
	Notifiers    []NotifierConfig
}

type AppConfig struct {
}

type InterpreterConfig struct {
}

type NotifierConfig struct {
	Channels map[string]string
}
