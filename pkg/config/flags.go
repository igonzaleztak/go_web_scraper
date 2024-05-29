package config

import "intelygenz/pkg/enums"

// Flags holds all the values required to execute the program
type Flags struct {
	Mode       enums.ModeType    // Mode defines whether the scraper is going to fetch data from the API or from the HTML
	Verbose    enums.VerboseMode // Verbose defines de log level
	MaxStories int               // MaxStories defines the maximum number of stories that the server can fetch
	NumWords   int               // NumWords defines the number of words that a title must have to be considered long
	Section    enums.SectionType // Sources categories defined by the hacker news API
}

var CmdFlags = &Flags{}
