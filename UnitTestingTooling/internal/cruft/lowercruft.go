package cruft

import (
	"fmt"
	"strings"
)

// LowerCruft implements the Crufter interface:
type LowerCruft struct {
	maxLength int
}

// Encruft switches the given value to lower case:
func (lc *LowerCruft) Encruft(value string) (string, error) {
	if len(value) == 0 {
		return value, fmt.Errorf("You gave me an empty string!")
	}

	return strings.ToLower(value)[:lc.maxLength], nil
}
