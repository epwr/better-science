// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package knowledge_graphs

import (
	"fmt"
	"testing"
)


// Assert an expression is true.
func assertTrue(t *testing.T, value bool, message string) {

	t.Helper()
	if !value {
		t.Errorf("Assertion Failed: " + message)
	}
}

// Assert the err value is not nil.
func assertError(t *testing.T, err error, message string) {

	t.Helper()
	if err == nil {
		t.Errorf("Assertion Failed: " + message)
	}
}

// Assert two slices have the same length and the same elements in each position.
func assertSlicesAreIdentical[T comparable](t *testing.T, sliceOne []T, sliceTwo []T, message string) {

	t.Helper()
	assertTrue(t, len(sliceOne) == len(sliceTwo), message + " [The two slices should have the same length]")
	if len(sliceOne) == len(sliceTwo) {
		for i := 0; i < len(sliceOne); i++ {
			indexStr := fmt.Sprintf("%v", i)
			assertTrue(
				t,
				sliceOne[i] == sliceTwo[i],
				message + " [The two slices should have the same element at index " + indexStr + "]",
			)
		}
	}
}

// Assert two slices have the same elements, and the same number of each element
func assertSlicesAreSetwiseEqual[T comparable](t *testing.T, sliceOne []T, sliceTwo []T, message string) {

	t.Helper()
	assertTrue(t, len(sliceOne) == len(sliceTwo), message + " [The two slices should have the same length]")
	if len(sliceOne) == len(sliceTwo) {
		for _, element := range sliceOne {
			
			countOne := 0
			countTwo := 0
			for i := 0; i < len(sliceOne); i++ {
				if sliceOne[i] == element {
					countOne++
				}
				if sliceTwo[i] == element {
					countTwo++
				}
			}

			elemStr := fmt.Sprintf("%v", element)
			countOneStr := fmt.Sprintf("%v", countOne)
			countTwoStr := fmt.Sprintf("%v", countTwo)
			assertTrue(
				t,
				countOne == countTwo,
				message + " [For element '" + elemStr + "', found " +  countOneStr +
                " and " + countTwoStr + " elements in each slice respectively.]",
			)
		}
	}
}

// Assert an element is in a Slice
func assertElementIsInSlice[T comparable](t *testing.T, element T, slice []T, message string) {

	t.Helper()
	for _, value := range slice {
		if element == value {
			return
		} 
	}
	t.Error(message)
}


func assertEqualTestValues(t *testing.T, v1 TestValue, v2 TestValue, msg string) {

	t.Helper()

	assertTrue(t, v1.value == v2.value, msg + " [The .value fields were different.]")
	assertTrue(t, v1.weight == v2.weight, msg + " [The .weight fields were different.]")
}
