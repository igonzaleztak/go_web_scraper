package hackerNews

import (
	"github.com/stretchr/testify/suite"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"log"
	"math/rand"
	"os"
	"testing"
)

type OperationSuite struct {
	suite.Suite
}

func (suite *OperationSuite) SetupTest() {
	// initialize default config from root dir so the config file can be reached
	os.Chdir("../../..")
	_ = config.SetDefaultFlags()
	_ = logs.InitLogger()

	// reduce the number of news to decrease the execution time
	config.CmdFlags.MaxStories = 15

	log.Println("\n-----Setup complete-----")
}

func (suite *OperationSuite) TearDownTest() {
	log.Println("\n----Teardown complete----")
}

func (suite *OperationSuite) TestOperation() {
	suite.Run("ok. More long stories than short stores", func() {
		shortFakeStory := schemas.StoryHackerNews{ID: rand.Int(), Title: "short story"}
		fakeStories := []schemas.StoryHackerNews{
			{
				ID:    rand.Int(),
				Title: "This is a very very long story",
			},
			{
				ID:    rand.Int(),
				Title: "This is another very very long story",
			},
			shortFakeStory,
		}
		short, long := splitStoriesByLength(fakeStories)

		suite.NotEmpty(short)
		suite.NotEmpty(long)

		suite.Greater(long.Len(), short.Len())
		suite.Equal(2, long.Len())
		suite.Equal(1, short.Len())

		suite.Equal(shortFakeStory.Title, short.Get()[0].Title)
	})

	suite.Run("ok. Same number of long stories and short stories", func() {
		fakeStories := []schemas.StoryHackerNews{
			{
				ID:    rand.Int(),
				Title: "This is another very very long story",
			},
			{
				ID:    rand.Int(),
				Title: "short story",
			},
		}
		short, long := splitStoriesByLength(fakeStories)

		suite.NotEmpty(short)
		suite.NotEmpty(long)

		suite.Equal(long.Len(), short.Len())
		suite.Equal(1, long.Len())
		suite.Equal(1, short.Len())
	})

}

func TestOperationSuite(t *testing.T) {
	suite.Run(t, new(OperationSuite))
}
