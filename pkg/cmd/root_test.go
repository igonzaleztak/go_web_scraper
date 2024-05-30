package cmd

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"intelygenz/pkg/config"
	"log"
	"os"
	"testing"
)

type CmdSuite struct {
	suite.Suite
}

func (suite *CmdSuite) SetupTest() {
	_ = os.Chdir("../../")
	_ = config.SetDefaultFlags()

	log.Println("\n-----Setup complete-----")
}

func (suite *CmdSuite) TearDownTest() {
	log.Println("\n----Teardown complete----")
}

func (suite *CmdSuite) TestExecute() {
	suite.Run("ok", func() {
		maxStories := 2
		numWords := 20
		args := []string{
			fmt.Sprintf("--max-stories=%d", maxStories),
			fmt.Sprintf("--num-words=%d", numWords),
			fmt.Sprintf("--verbose=%d", 1),
		}
		cmd := rootCmd
		cmd.SetArgs(args)

		err := cmd.Execute()
		suite.NoError(err)

		// verify flags are set correctly
		suite.Equal(maxStories, config.CmdFlags.MaxStories)
		suite.Equal(numWords, config.CmdFlags.NumWords)
	})
}

func TestCmdSuite(t *testing.T) {
	suite.Run(t, new(CmdSuite))
}
