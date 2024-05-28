package utils

import (
	"encoding/json"
	"fmt"
)

// PrettyPrintStruct prints prettily a struct
func PrettyPrintStruct(d any) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Print(string(s))
}
