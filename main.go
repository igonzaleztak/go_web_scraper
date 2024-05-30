package main

import (
	"fmt"
	"intelygenz/pkg/cmd"
	"intelygenz/pkg/config"
	"os"

	_ "go.uber.org/automaxprocs"
)

func main() {
	// init app configuration and set default flags values
	if err := config.SetDefaultFlags(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// execute the root command
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
