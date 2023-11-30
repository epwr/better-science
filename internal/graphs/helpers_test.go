// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package graphs_test

import (
	"testing"
)

func assertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Errorf("Assertion Failed: " + message)
	}
}
