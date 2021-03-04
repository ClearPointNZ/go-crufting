package cruft

import (
	"fmt"
	"strings"
)

// UpperCruft implements the Crufter interface:
type UpperCruft struct {
	prefix string
}

// Encruft switches the given value to upper case:
func (uc *UpperCruft) Encruft(value string) (string, error) {
	if len(value) == 0 {
		return value, fmt.Errorf("You gave me an empty string!")
	}

	// Return the upper-case value with our configured prefix:
	return fmt.Sprintf("%s%s", uc.prefix, strings.ToUpper(value)), nil
}
