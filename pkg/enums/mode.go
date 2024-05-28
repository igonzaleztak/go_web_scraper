package enums

import "fmt"

// This enums uses the *FlagSet interface defined in the cobra pflag package:
// https://github.com/spf13/pflag

// By using this interface, we can ensure that input arguments defined by users are correct.

type ModeType string

const (
	ModeTypeAPI ModeType = "api"
	ModeTypeWeb ModeType = "web"
)

// String returns the enum value in the form of a string
func (e *ModeType) String() string {
	return string(*e)
}

// Set sets the value of the flag
func (e *ModeType) Set(v string) error {
	switch v {
	case "api", "web":
		*e = ModeTypeAPI
		return nil
	default:
		return fmt.Errorf("invalid mode type: %s. Must be %s or %s", v, ModeTypeAPI, ModeTypeWeb)
	}
}

// Type is only used in help text
func (e *ModeType) Type() string {
	return "mode"
}
