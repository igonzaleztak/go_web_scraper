package scraper

import (
	"fmt"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
)

// StartScraperProcess is the main function that starts the scraping process. The flow is as follows:
// 1. Scrape the latest stories from the Hacker News API.
// 2. Print data accordingly.
func StartScraperProcess(s Scraper) error {
	logs.Logger.Infof("[scraper] starting scraping function with the following args: %s", config.CmdFlags.ToString())
	if err := s.Scrap(); err != nil {
		return fmt.Errorf("failed to scrape stories: %v", err)
	}

	// print data according to the requirements
	s.Print()

	logs.Logger.Info("[scraper] scraping function finished")
	return nil
}
