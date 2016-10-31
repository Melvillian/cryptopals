package lib

import (
	"fmt"
)

// PrintComparisons formats on two lines the expected and actual values.
// Meant to visually test the results of functions
func PrintComparisons(expected string, actual string) {
	fmt.Println("Expected: " + expected + "\nActual:   " + actual)
}

// CheckAbort checks if an error exists, and if so ends the process
func CheckAbort(err error) {
	if err != nil {
		panic(err)
	}
}