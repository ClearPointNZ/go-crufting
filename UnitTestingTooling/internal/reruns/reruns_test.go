package reruns

import (
	"fmt"
	"testing"
)

type testCase struct {
	value         int32
	errorExpected bool
}

func TestReRuns(t *testing.T) {

	// Set up some test data we'd like to run:
	var testCases = []testCase{
		{
			value:         1,
			errorExpected: true,
		},
		{
			value:         3,
			errorExpected: true,
		},
		{
			value:         5,
			errorExpected: false,
		},
		{
			value:         7,
			errorExpected: false,
		},
		{
			value:         9,
			errorExpected: false,
		},
	}

	for testCaseNumber, testCase := range testCases {

		runName := fmt.Sprintf("Run %d", testCaseNumber)

		t.Run(runName, func(t *testing.T) {
			_, err := DoSomethingWithANumber(testCase.value)

			if err == nil && testCase.errorExpected {
				t.Fail()
			}

			if err != nil && !testCase.errorExpected {
				t.Fail()
			}

		})
	}

}
