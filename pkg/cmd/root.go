package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/scraper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Intelygenz Scraper aims to obtain several from the Hacker News website",
	Long:  `Intelygenz Scraper aims to obtain several from the Hacker News website`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init logs
		if err := logs.InitLogger(); err != nil {
			return fmt.Errorf("failed to init logs: %v", err)
		}

		// Do Stuff Here
		scraper.StartScraperProcess()
		return nil
	},
}

func Execute() {
	// set default flags
	setDefaultFlags()

	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Mode, "mode", "m", "Mode in which the tool should work. Options: 'api', 'web'")
	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Verbose, "verbose", "v", "Enable verbose mode")
	rootCmd.PersistentFlags().StringArrayP("sources", "s", config.CmdFlags.Sources, "Indicates list of sections from hacker news that can be obtained")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
