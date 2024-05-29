package config

import "intelygenz/pkg/enums"

// Flags holds all the values required to execute the program
type Flags struct {
	Verbose    enums.VerboseMode // Verbose defines de log level
	MaxStories int               // MaxStories defines the maximum number of stories that the server can fetch
	NumWords   int               // NumWords defines the number of words that a title must have to be considered long
}

var CmdFlags = &Flags{}
