package spaceFlight

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
	Data *[]schemas.SpaceFlightNews
}

// NewScraper creates a new instance of the Scraper struct for space flight news
func NewScraper(api string) *Scraper {
	return &Scraper{
		API:  api,
		Data: &[]schemas.SpaceFlightNews{},
	}
}

// Scrap is the main function that starts the scraping process. It gets the latest stories from the space flight API.
func (s *Scraper) Scrap() error {
	logs.Logger.Info("[spaceflight.scraper] starting space flight scraper")
	uri := fmt.Sprintf("%s/articles/?format=json&limit=%d", s.API, config.CmdFlags.MaxStories)
	logs.Logger.Debugf("[spaceflight.scraper] got space flight API host: %s", uri)

	logs.Logger.Infof("[spaceflight.scraper] getting latest stories from: %s", uri)
	logs.Logger.Debugf("[spaceflight.scraper] executing get request to: %s", uri)
	resp, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("failed to execute get to '%s': %v", uri, err)
	}
	logs.Logger.Debugf("[spaceflight.scraper] got response: %v", resp)

	logs.Logger.Debugf("[spaceflight.scraper] reading response body")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[spaceflight.scraper]  failed to read response body: %v", err)
	}
	logs.Logger.Debugf("[spaceflight.scraper] got response body: %v", string(body))

	logs.Logger.Debugf("[spaceflight.scraper] unmarshalling response body")
	var response schemas.SpaceFlightResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("[spaceflight.scraper]  failed to unmarshal response body: %v", err)
	}
	logs.Logger.Debugf("[spaceflight.scraper] got unmarshalled response body: %v", response)

	s.Data = &response.Results
	logs.Logger.Infof("[spaceflight.scraper] got %d news", len(*s.Data))
	return nil
}

// Print prints the stories in order: long stories first, then short stories. In this case we do not need to sort the
// stories, we just need to print them in order.
func (s *Scraper) Print() {
	short, long := splitStoriesByLength(*s.Data)

	logs.Logger.Info("[spaceFlight.scraper] printing stories in order: long stories first, then short stories")
	logs.Logger.Infof("[spaceFlight.scraper] printing long stories \n")
	long.Print()
	logs.Logger.Infof("[spaceFlight.scraper] ---------------------------- \n")
	logs.Logger.Infof("[spaceFlight.scraper] ---------------------------- \n")

	logs.Logger.Infof("[spaceFlight.scraper] printing short stories \n")
	short.Print()
	logs.Logger.Infof("[spaceFlight.scraper] ---------------------------- \n")
	logs.Logger.Infof("[spaceFlight.scraper] ---------------------------- \n")
}
