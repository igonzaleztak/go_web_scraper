package scraper

import (
	"fmt"
	"intelygenz/pkg/logs"
)

// StartScraperProcess is the main function that starts the scraping process. The flow is as follows:
// 1. Scrape the latest stories from the Hacker News API.
// 2. Split the stories by the length of the title. Short stories are those with a title length of less than 5 words. Long stories are those with a title length of more than 5 words.
// 3. Sort stories accordingly. Short stories must be sorted by score. Long stories must be sorted by number of comments.
// 4. Print stories in console in order. First long stories and then sort stories.
func StartScraperProcess(s Scraper) error {
	logs.Logger.Info("[scraper] starting scraping function")
	if err := s.Scrap(); err != nil {
		return fmt.Errorf("failed to scrape stories: %v", err)
	}

	// print data according to the requirements
	s.Print()

	logs.Logger.Info("[scraper] scraping function finished")
	return nil
}
