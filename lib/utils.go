package lib

import (
	"fmt"
)

// PrintComparisons formats one two lines the expected and actual values.
// Meant to visually test the results of functions
func PrintComparisons(expected string, actual string) {
	fmt.Println("Expected: " + expected + "\nActual:   " + actual)
}
