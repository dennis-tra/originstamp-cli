package pkg

import (
	"fmt"
	"strings"
)

// EnumValue is used for command line flags that can only hold certain values
type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

// Set is used to make sure only allowed values are set
func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}
