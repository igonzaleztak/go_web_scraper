package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"intelygenz/pkg/enums"
	"os"
)

// Flags holds all the values required to execute the program
type Flags struct {
	Verbose    enums.VerboseMode // Verbose defines de log level
	MaxStories int               // MaxStories defines the maximum number of stories that the server can fetch
	NumWords   int               // NumWords defines the number of words that a title must have to be considered long
}

func (f *Flags) ToString() string {
	return fmt.Sprintf("Verbose: %v, MaxStories: %d, NumWords: %d", f.Verbose, f.MaxStories, f.NumWords)
}

var CmdFlags = &Flags{}

// SetDefaultFlags  defines the default flag values and the app configuration
func SetDefaultFlags() error {
	viper.SetDefault("CONFIG_PATH", "config/config.yaml")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	configPath := viper.Get("CONFIG_PATH").(string)
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file from path '%s' : %v", configPath, err)
	}

	err = viper.ReadConfig(bytes.NewReader(configBytes))
	if err != nil {
		return fmt.Errorf("viper failed to read config file from path '%s' : %v", configPath, err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file from path '%s' : %v", configPath, err)
	}

	// set default flags
	CmdFlags.Verbose = enums.VerboseModeInfo
	CmdFlags.MaxStories = 30
	CmdFlags.NumWords = 5

	return nil
}
