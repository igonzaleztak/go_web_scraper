package config

import "intelygenz/pkg/enums"

// Flags holds all the values required to execute the program
type Flags struct {
	Mode    enums.ModeType    // Mode defines whether the scraper is going to fetch data from the API or from the HTML
	Verbose enums.VerboseMode // Verbose defines de log level
	Sources []string          //list of hacker news categories that the server can scrap
}

var CmdFlags = &Flags{}
