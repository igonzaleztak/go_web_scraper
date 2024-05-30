package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/scraper"
	"intelygenz/pkg/scraper/hackerNews"
	"intelygenz/pkg/scraper/spaceFlight"
)

var rootCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Intelygenz Scraper aims to obtain several news from the Hacker News website",
	Long:  `Intelygenz Scraper aims to obtain several news from the Hacker News website`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init logs
		if err := logs.InitLogger(); err != nil {
			return fmt.Errorf("failed to init logs: %v", err)
		}

		// start scrapers
		hackerNewsScraper := hackerNews.NewScraper(config.AppConfig.HackersNewsAPI)
		spaceFlightScraper := spaceFlight.NewScraper(config.AppConfig.SpaceFlightNewsAPI)

		ctx := cmd.Context()
		errs, ctx := errgroup.WithContext(ctx)

		errs.Go(func() error {
			return scraper.StartScraperProcess(hackerNewsScraper)
		})
		errs.Go(func() error {
			return scraper.StartScraperProcess(spaceFlightScraper)
		})

		return errs.Wait()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Verbose, "verbose", "v", "Enable verbose mode. Supported modes Debug: 0, Info: 1")
	rootCmd.PersistentFlags().IntVarP(&config.CmdFlags.MaxStories, "max-stories", "n", config.CmdFlags.MaxStories, "Defines the number of news that will be fetched from the sources")
	rootCmd.PersistentFlags().IntVarP(&config.CmdFlags.NumWords, "num-words", "w", config.CmdFlags.NumWords, "Indicates the number of words that a title must have to be considered long")
}
