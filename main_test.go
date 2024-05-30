package main_test

import (
	"github.com/stretchr/testify/suite"
	"os/exec"
	"testing"
)

type IntegrationSuite struct {
	suite.Suite
}

func (suite *IntegrationSuite) TestCommand() {

	suite.Run("ok", func() {
		cmd := exec.Command("go", "run", "main.go", "--max-stories=2", "--num-words=20", "--verbose=1")
		err := cmd.Run()

		suite.NoError(err)
	})

	suite.Run("fail. Invalid args", func() {
		cmd := exec.Command("go", "run", "main.go", "--max-a=2")
		err := cmd.Run()

		suite.Error(err)
	})

	suite.Run("fail. Invalid flag", func() {
		cmd := exec.Command("go", "run", "main.go", "--max-stories=2", "--num-words=20", "--verbose=2")
		err := cmd.Run()

		suite.Error(err)
	})

}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}
