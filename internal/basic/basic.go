package basic

import (
	"math/rand"
	"time"
)

// Scramble randomises a string:
func Scramble(value string) string {

	// Turn the value into a byte-array (so we can shuffle it):
	values := []byte(value)

	// Make this actually random:
	rand.Seed(time.Now().UnixNano())

	// Shuffle the order of the string:
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	return string(values)
}
