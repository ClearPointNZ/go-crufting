package cruft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerCruft(t *testing.T) {

	testValue := "I aM sOme kINd of TEst vAlue"

	// Prepare a lowerCruft to test:
	lowerCruft := &LowerCruft{
		maxLength: 8,
	}

	// Ensure that lowerCruft implements the Crufter interface:
	assert.Implements(t, (*Crufter)(nil), lowerCruft)

	// Invoke the Encruft method:
	encruftedValue, err := lowerCruft.Encruft(testValue)

	// Ensure that there was no error:
	assert.NoError(t, err)

	// Ensure that we got the expected result:
	assert.Equal(t, "i am som", encruftedValue)

	// Ensure that the result was the correct length:
	assert.Len(t, encruftedValue, 8)
}
