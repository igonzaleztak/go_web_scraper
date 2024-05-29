package hackerNews

import (
	"encoding/json"
	"fmt"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"io"
	"net/http"
	"strings"
)

// getStoryDetails fetches the details of a story from the hacker news API.
func getStoryDetails(id int) (*schemas.StoryHackerNews, error) {
	logs.Logger.Debugf("[scraper.GetStoryDetails] getting story details for id: %d", id)
	uri := fmt.Sprintf("%s/item/%d.json?print=pretty", config.AppConfig.HackersNewsAPI, id)
	logs.Logger.Debugf("[scraper.GetStoryDetails] executing get request to: %s", uri)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get to '%s': %v", uri, err)
	}
	logs.Logger.Debugf("[scraper.GetStoryDetails] got response for item '%d': %v", id, resp)

	logs.Logger.Debugf("[scraper.GetStoryDetails] reading response body")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	logs.Logger.Debugf("[scraper.GetStoryDetails] got response body: %v", string(body))

	logs.Logger.Debugf("[scraper.GetStoryDetails] unmarshalling response body for item '%d'", id)
	var story schemas.StoryHackerNews
	if err := json.Unmarshal(body, &story); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body for item '%d': %v", id, err)
	}
	logs.Logger.Debugf("[scraper.GetStoryDetails] got unmarshalled response body for item '%d': %v", id, story)
	return &story, nil
}

// splitStoriesByLength splits the stories into two categories: short and long. Short stories are those with a title length of less than 5 words. Long stories are those with a title length of more than 5 words.
func splitStoriesByLength(stories []schemas.StoryHackerNews) (ShortData, LongData) {
	logs.Logger.Info("[scraper.SplitStoriesByLength] splitting stories by length")
	short := make(ShortData, 0)
	long := make(LongData, 0)

	for _, story := range stories {
		logs.Logger.Debugf("[scraper.SplitStoriesByLengths] checking story %d with title `%s`", story.ID, story.Title)
		if len(strings.Fields(story.Title)) < config.CmdFlags.NumWords {
			logs.Logger.Debugf("[scraper.SplitStoriesByLengths] story %d is short", story.ID)
			short = append(short, story)
		} else {
			logs.Logger.Debugf("[scraper.SplitStoriesByLengths] story %d is long", story.ID)
			long = append(long, story)
		}
	}
	logs.Logger.Infof("[scraper.SplitStoriesByLength] got %d short stories and %d long stories", short.Len(), long.Len())
	return short, long
}
