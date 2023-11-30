// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package graphs

import (
	"fmt"
	"testing"
)

func assertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Errorf("Assertion Failed: " + message)
	}
}

func assertError(t *testing.T, err error, message string) {
	fmt.Println("Error:", err)
	if err == nil {
		t.Errorf("Assertion Failed: " + message)
	}
}

// Assert that the graph.edge adjacency map contains exactly one reference
// from nodeOne to nodeTwo and exactly one reference from nodeTwo to nodeOne
func assertEdgeExistsExactlyOnce[ID comparable, Payload any](
	t *testing.T,
	g Graph[ID, Payload],
	nodeOne ID,
	nodeTwo ID,
) {

	nodeOneNeighbours, foundNodeOne := g.edges[nodeOne]
	nodeTwoNeighbours, foundNodeTwo := g.edges[nodeTwo]
	assertTrue(t, foundNodeOne && foundNodeTwo, "A node was not found in the graph's edges.")

	countNodeTwoReferences := 0
	for _, neighbour := range nodeOneNeighbours {
		if neighbour == nodeTwo {
			countNodeTwoReferences++
		}
	}
	countNodeOneReferences := 0
	for _, neighbour := range nodeTwoNeighbours {
		if neighbour == nodeOne {
			countNodeOneReferences++
		}
	}

	edgeExistsOnlyOnce := countNodeTwoReferences == 1 && countNodeOneReferences == 1
	assertTrue(
		t,
		edgeExistsOnlyOnce,
		"The graph's edges did not contain a reference to the other node exactly once.",
	)
}
