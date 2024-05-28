package config

// DefaultConfig stores the hacker news configuration including its URL and API URL
type DefaultConfig struct {
	APIURL   string `yaml:"apiURL"`
	WebURL   string `yaml:"webURL"`
	LogsPath string `yaml:"logsPath"`
}

var AppConfig = &DefaultConfig{}
