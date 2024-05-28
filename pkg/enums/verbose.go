package enums

import "fmt"

// This enums uses the *FlagSet interface defined in the cobra pflag package:
// https://github.com/spf13/pflag

// By using this interface, we can ensure that input arguments defined by users are correct.

type VerboseMode int

const (
	VerboseModeLog VerboseMode = iota
	VerboseModeInfo
)

// String returns the enum value in the form of a string
func (e *VerboseMode) String() string {
	switch *e {
	case VerboseModeLog:
		return "log"
	case VerboseModeInfo:
		return "info"
	default:
		return ""
	}
}

// Int returns the value in the form of an int
func (e *VerboseMode) Int() int {
	return int(*e)
}

// Set sets the value of the flag
func (e *VerboseMode) Set(v string) error {
	switch v {
	case "0":
		*e = VerboseModeLog
		return nil
	case "1":
		*e = VerboseModeInfo
		return nil
	default:
		return fmt.Errorf("invalid verbose mode: %s", v)
	}
}

// Type is only used in help text
func (e *VerboseMode) Type() string {
	return "verbose"
}
