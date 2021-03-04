package basic

import "testing"

func TestScrable(t *testing.T) {
	testValue := "something"
	scrambledValue := Scramble(testValue)

	// Fail the test if our value wasn't scrambled:
	if scrambledValue == testValue {
		t.Errorf("Our test value (%s) wasn't scrambled: %s", testValue, scrambledValue)
		t.Fail()
	}
}
