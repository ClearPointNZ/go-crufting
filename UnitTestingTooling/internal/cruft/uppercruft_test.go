package cruft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperCruft(t *testing.T) {

	testValue := "I aM sOme kINd of TEst vAlue"

	// Prepare a upperCruft to test:
	upperCruft := &UpperCruft{
		prefix: "cruft",
	}

	// Ensure that upperCruft implements the Crufter interface:
	assert.Implements(t, (*Crufter)(nil), upperCruft)

	// Invoke the Encruft method:
	encruftedValue, err := upperCruft.Encruft(testValue)

	// Ensure that there was no error:
	assert.NoError(t, err)

	// Ensure that we got the expected result:
	assert.Equal(t, "cruftI AM SOME KIND OF TEST VALUE", encruftedValue)

	// Ensure that the result contained the configured prefix:
	assert.Contains(t, encruftedValue, "cruft")
}
