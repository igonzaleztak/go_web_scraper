package spaceFlight

import (
	"github.com/stretchr/testify/suite"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"log"
	"os"
	"testing"
)

type ScraperSuite struct {
	suite.Suite
}

func (suite *ScraperSuite) SetupTest() {
	// initialize default config from root dir so the config file can be reached
	os.Chdir("../../..")
	_ = config.SetDefaultFlags()
	_ = logs.InitLogger()

	// reduce the number of news to decrease the execution time
	config.CmdFlags.MaxStories = 3

	log.Println("\n-----Setup complete-----")
}

func (suite *ScraperSuite) TearDownTest() {
	log.Println("\n----Teardown complete----")
}

func (suite *ScraperSuite) TestNewScraper() {
	apiURL := config.AppConfig.HackersNewsAPI
	suite.Run("ok", func() {
		s := NewScraper(apiURL)
		suite.NotNil(s)
		suite.Equal(s.API, apiURL)
	})
}

func (suite *ScraperSuite) TestScrap() {
	apiURL := config.AppConfig.SpaceFlightNewsAPI
	s := NewScraper(apiURL)

	suite.Run("ok", func() {
		err := s.Scrap()

		suite.NoError(err)
		suite.NotEmpty(s.Data)
	})
}

func TestScraperSuite(t *testing.T) {
	suite.Run(t, new(ScraperSuite))
}
