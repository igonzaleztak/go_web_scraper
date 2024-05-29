package hackerNews

import (
	"github.com/stretchr/testify/suite"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"log"
	"math/rand"
	"os"
	"testing"
)

type StoriesSuite struct {
	suite.Suite
}

func (suite *StoriesSuite) SetupTest() {
	// initialize default config from root dir so the config file can be reached
	os.Chdir("../../..")
	_ = config.SetDefaultFlags()
	_ = logs.InitLogger()

	// reduce the number of news to decrease the execution time
	config.CmdFlags.MaxStories = 15

	log.Println("\n-----Setup complete-----")
}

func (suite *StoriesSuite) TearDownTest() {
	log.Println("\n----Teardown complete----")
}

func (suite *StoriesSuite) TestShortData() {
	fakeShortData := ShortData{
		{ID: rand.Int(), Score: 23},
		{ID: rand.Int(), Score: 100},
		{ID: rand.Int(), Score: 0},
	}

	suite.Run("ok. Sort", func() {
		fakeShortData.Sort()

		oldScore := fakeShortData[0].Score
		for _, d := range fakeShortData {
			suite.LessOrEqual(d.Score, oldScore)
			oldScore = d.Score
		}
	})

	suite.Run("ok. Len", func() {
		suite.NotZero(fakeShortData.Len())
	})

	suite.Run("ok. Get", func() {
		suite.NotEmpty(fakeShortData.Get())
	})

}

func (suite *StoriesSuite) TestLongData() {
	fakeLongData := LongData{
		{ID: rand.Int(), Descendants: 23},
		{ID: rand.Int(), Descendants: 100},
		{ID: rand.Int(), Descendants: 0},
	}

	suite.Run("ok. Sort", func() {
		fakeLongData.Sort()

		oldScore := fakeLongData[0].Descendants
		for _, d := range fakeLongData {
			suite.LessOrEqual(d.Descendants, oldScore)
			oldScore = d.Descendants
		}
	})

	suite.Run("ok. Len", func() {
		suite.NotZero(fakeLongData.Len())
	})

	suite.Run("ok. Get", func() {
		suite.NotEmpty(fakeLongData.Get())
	})

}

func TestStoriesSuite(t *testing.T) {
	suite.Run(t, new(StoriesSuite))
}
