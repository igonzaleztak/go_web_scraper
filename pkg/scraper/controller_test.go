package scraper

import (
	"github.com/stretchr/testify/suite"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/scraper/hackerNews"
	"log"
	"os"
	"testing"
)

type ControllerSuite struct {
	suite.Suite
}

func (suite *ControllerSuite) SetupTest() {
	// initialize default config from root dir so the config file can be reached
	os.Chdir("../..")
	_ = config.SetDefaultFlags()
	_ = logs.InitLogger()

	// reduce the number of news to decrease the execution time
	config.CmdFlags.MaxStories = 15

	log.Println("\n-----Setup complete-----")
}

func (suite *ControllerSuite) TearDownTest() {
	log.Println("\n----Teardown complete----")
}

func (suite *ControllerSuite) TestStartScraperProcess() {

	suite.Run("ok default flags", func() {
		// initialize new hacker news scraper
		s := hackerNews.NewScraper(config.AppConfig.HackersNewsAPI)
		err := StartScraperProcess(s)
		suite.NoError(err)

		// check data is not zero
		scrapedData := s.GetData()
		suite.NotZero(scrapedData)
	})

	suite.Run("not ok. Invalid API URL", func() {
		newURI := "this is not an API URL"
		s := hackerNews.NewScraper(newURI)
		err := StartScraperProcess(s)
		suite.Error(err)
	})
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}
