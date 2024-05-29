package spaceFlight

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

func (suite *StoriesSuite) TestDataSpaceFlight() {
	fakeShortData := DataSpaceFlight{
		{ID: rand.Int()},
		{ID: rand.Int()},
		{ID: rand.Int()},
	}

	suite.Run("ok. Len", func() {
		suite.NotZero(fakeShortData.Len())
	})

	suite.Run("ok. Get", func() {
		suite.NotEmpty(fakeShortData.Get())
	})

}

func TestStoriesSuite(t *testing.T) {
	suite.Run(t, new(StoriesSuite))
}
