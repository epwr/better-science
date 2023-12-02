// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package data_structures

import (
	"testing"
)

func assertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Errorf("Assertion Failed: " + message)
	}
}

func assertError(t *testing.T, err error, message string) {
	if err == nil {
		t.Errorf("Assertion Failed: " + message)
	}
}
