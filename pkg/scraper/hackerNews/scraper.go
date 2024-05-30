package hackerNews

import (
	"encoding/json"
	"fmt"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"io"
	"net/http"
)

type Scraper struct {
	API  string
	Data *[]schemas.StoryHackerNews
}

// NewScraper creates a new instance of the Scraper struct for hacker news
func NewScraper(api string) *Scraper {
	return &Scraper{
		API:  api,
		Data: &[]schemas.StoryHackerNews{},
	}
}

// Scrap is the main function that starts the scraping process. It gets the latest stories from the hacker news API.
// The response from the API is sorted by the latest 500 stories. Therefore, we need to slice them according to the
// number of stories that we want to fetch.
//
// Hacker news API returns only the IDs of the stories. Therefore, we need to fetch the details of each story by making
// a request to the item endpoint.
func (s *Scraper) Scrap() error {
	logs.Logger.Info("[hackernews.scraper] starting hacker news scraper")
	uri := fmt.Sprintf("%s/newstories.json?print=pretty", s.API)
	logs.Logger.Debugf("[hackernews.scraper] got hacker news API host: %s", uri)

	logs.Logger.Infof("[hackernews.scraper] getting latest '%d' stories from: %s", config.CmdFlags.MaxStories, uri)
	logs.Logger.Debugf("[hackernews.scraper] executing get request to: %s", uri)
	resp, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("failed to execute get to '%s': %v", uri, err)
	}
	logs.Logger.Debugf("[hackernews.scraper] got response: %v", resp)

	logs.Logger.Debugf("[hackernews.scraper] reading response body")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	logs.Logger.Debugf("[hackernews.scraper] got response body: %v", string(body))

	logs.Logger.Debugf("[hackernews.scraper] unmarshalling response body")
	liveDataResponse := make(schemas.LiveDataHackerNews, 0)
	if err := json.Unmarshal(body, &liveDataResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}
	logs.Logger.Debugf("[hackernews.scraper] got unmarshalled response body: %v", liveDataResponse)

	// The response from the API is sorted by the latest stories. Therefore, we can slice the first 30 stories and fetch
	// their details.
	latest := liveDataResponse[0:config.CmdFlags.MaxStories]
	stories := make([]schemas.StoryHackerNews, 0)
	for _, id := range latest {
		story, err := getStoryDetails(id)
		if err != nil {
			logs.Logger.Errorf("[hackernews.scraper] failed to get story details for id: %d: %v", id, err)
			continue
		}
		stories = append(stories, *story)
	}

	*s.Data = stories
	return nil
}

// GetData returns the data fetched from the hacker news API
func (s *Scraper) GetData() *[]schemas.StoryHackerNews {
	return s.Data
}

// Print prints the stories in order: long stories first, then short stories. In this case we need to sort the stories
// by number of comments and score respectively.
func (s *Scraper) Print() {
	// split stories by length
	short, long := splitStoriesByLength(*s.Data)

	// sort long and short stories
	logs.Logger.Debugf("[hackernews.scraper] sorting short stories by score")
	short.Sort()
	logs.Logger.Debugf("[hackernews.scraper] short stories were sorted successfully")

	logs.Logger.Debugf("[hackernews.scraper] sorting long stories by number of comments")
	long.Sort()
	logs.Logger.Debugf("[hackernews.scraper] long stories were sorted successfully")

	// print long stories first and then short stories
	logs.Logger.Info("[hackernews.scraper] printing stories in order: long stories first, then short stories")
	logs.Logger.Infof("[hackernews.scraper] printing long stories \n")
	long.Print()
	logs.Logger.Infof("[hackernews.scraper] ---------------------------- \n")
	logs.Logger.Infof("[hackernews.scraper] ---------------------------- \n")

	logs.Logger.Infof("[hackernews.scraper] printing short stories \n")
	short.Print()
	logs.Logger.Infof("[hackernews.scraper] ---------------------------- \n")
	logs.Logger.Infof("[hackernews.scraper] ---------------------------- \n")
}
