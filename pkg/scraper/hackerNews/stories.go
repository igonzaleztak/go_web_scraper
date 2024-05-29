package hackerNews

import (
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"sort"
	"time"
)

// printStory prints stories to console
func printStory(story schemas.StoryHackerNews) {
	logs.Logger.Info("--------------------------------------------------------------------")
	logs.Logger.Infof("[hackernews.scraper] story ID: %d", story.ID)
	logs.Logger.Infof("[hackernews.scraper] story timestamp: %s", story.Time.Format(time.RFC3339))
	logs.Logger.Infof("[hackernews.scraper] story title: %s", story.Title)
	logs.Logger.Infof("[hackernews.scraper] story URL: %s", story.URL)
	logs.Logger.Infof("[hackernews.scraper] story score: %d", story.Score)
	logs.Logger.Infof("[hackernews.scraper] story descendants: %d", story.Descendants)
	logs.Logger.Infof("[hackernews.scraper] story text: %v", story.Text)
	logs.Logger.Infof("--------------------------------------------------------------------\n")
}

type ShortData []schemas.StoryHackerNews
type LongData []schemas.StoryHackerNews

// Len returns the length of the short stories.
func (s ShortData) Len() int {
	return len(s)
}

// Sort sorts the short stories by number of points.
func (s ShortData) Sort() {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Score > s[j].Score
	})
}

// Get returns the short data
func (s ShortData) Get() []schemas.StoryHackerNews {
	return s
}

// Print prints the short stories.
func (s ShortData) Print() {
	for _, story := range s {
		printStory(story)
	}
}

// Len returns the length of the long stories.
func (l LongData) Len() int {
	return len(l)
}

// Sort the short stories by number of descendants. Descendants are the number of comments a story has
func (l LongData) Sort() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Descendants > l[j].Descendants
	})
}

// Get returns the long data
func (l LongData) Get() []schemas.StoryHackerNews {
	return l
}

// Print prints the long stories.
func (l LongData) Print() {
	for _, story := range l {
		printStory(story)
	}
}
