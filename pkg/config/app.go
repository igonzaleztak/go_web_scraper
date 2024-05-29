package config

// DefaultConfig stores the hacker news configuration including its URL and API URL
type DefaultConfig struct {
	HackersNewsAPI     string `yaml:"hackersNewsAPI"`
	SpaceFlightNewsAPI string `yaml:"spaceFlightNewsAPI"`
	LogsPath           string `yaml:"logsPath"`
}

var AppConfig = &DefaultConfig{}
