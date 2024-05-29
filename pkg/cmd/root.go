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
	"os"

	_ "go.uber.org/automaxprocs"
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

		// Do Stuff Here
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

func Execute() {
	// set default flags
	if err := setDefaultFlags(); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Mode, "mode", "m", "Mode in which the tool should work. Options: 'api', 'web'")
	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Verbose, "verbose", "v", "Enable verbose mode")
	rootCmd.PersistentFlags().IntVarP(&config.CmdFlags.MaxStories, "max-stories", "n", config.CmdFlags.MaxStories, "Indicates the maximum number of stories that the server can fetch")
	rootCmd.PersistentFlags().IntVarP(&config.CmdFlags.NumWords, "num-words", "w", config.CmdFlags.NumWords, "Indicates the number of words that a title must have to be considered long")
	rootCmd.PersistentFlags().VarP(&config.CmdFlags.Section, "section", "s", "Indicates the section of the stories to fetch. Options: 'new', 'past', 'comments', 'ask', 'show', 'jobs' or 'submit'")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
