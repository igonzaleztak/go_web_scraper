package spaceFlight

import (
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"time"
)

// printStory prints stories to console
func printStory(story schemas.SpaceFlightNews) {
	logs.Logger.Info("--------------------------------------------------------------------")
	logs.Logger.Infof("[spaceflight.scraper] story ID: %d", story.ID)
	logs.Logger.Infof("[spaceflight.scraper] story timestamp: %s", story.PublishedAt.Format(time.RFC3339))
	logs.Logger.Infof("[spaceflight.scraper] story title: %s", story.Title)
	logs.Logger.Infof("[spaceflight.scraper] story URL: %s", story.URL)
	logs.Logger.Infof("[spaceflight.scraper] story summary: %s", story.Summary)
	logs.Logger.Infof("--------------------------------------------------------------------\n")
}

type DataSpaceFlight []schemas.SpaceFlightNews

// Len returns the length of the short stories.
func (s DataSpaceFlight) Len() int {
	return len(s)
}

// Sort is an empty function since in the spaceFlight API is not defined how to sort the news.
func (s DataSpaceFlight) Sort() {}

// Print prints the stories.
func (s DataSpaceFlight) Print() {
	for _, story := range s {
		printStory(story)
	}
}
