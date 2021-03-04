package reruns

import "fmt"

// DoSomethingWithANumber does something with a number, returns a new number:
func DoSomethingWithANumber(value int32) (int32, error) {

	// Sometimes an error will occur:
	if value < 5 {
		return value, fmt.Errorf("Value is too low")
	}

	return value * 2, nil
}
