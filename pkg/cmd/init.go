package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"intelygenz/pkg/config"
	"intelygenz/pkg/enums"
	"os"
)

// setDefaultFlags  set the default values of the flags
func setDefaultFlags() error {
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

	err = viper.Unmarshal(&config.AppConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file from path '%s' : %v", configPath, err)
	}

	// set default flags
	config.CmdFlags.Mode = enums.ModeTypeAPI
	config.CmdFlags.Verbose = enums.VerboseModeLog
	config.CmdFlags.MaxStories = 30
	config.CmdFlags.NumWords = 5
	config.CmdFlags.Section = enums.SectionTypeNew

	return nil
}
