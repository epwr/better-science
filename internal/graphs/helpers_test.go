// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package graphs

import (
	"strconv"
	"testing"
)

// Assert an expression is true.
func assertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Errorf("Assertion Failed: " + message)
	}
}

// Assert the err value is not nil.
func assertError(t *testing.T, err error, message string) {
	if err == nil {
		t.Errorf("Assertion Failed: " + message)
	}
}

// Assert two slices have the same length and the same elements in each position.
func assertSliceIsEqual[T comparable](t *testing.T, sliceOne []T, sliceTwo []T, message string) {

	assertTrue(t, len(sliceOne) == len(sliceTwo), message + " [The two slices should have the same length]")
	if len(sliceOne) == len(sliceTwo) {
		for i := 0; i < len(sliceOne); i++ {
			indexStr := strconv.Itoa(i)
			assertTrue(
				t,
				sliceOne[i] == sliceTwo[i],
				message + " [The two slices should have the same element at index " + indexStr + "]",
			)
		}
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




// Initialize a Graph
func initGraph[ID comparable, Payload any](nodes []ID, edges []ID) Graph[ID, Payload] {

	g := NewGraph[ID, Payload]()
	for index := 0; index < len(nodes); index++ {
		g.AddNode(nodes[index], *new(Payload))
	}

	for index := 0; index < len(edges); index = index + 2 {
		g.AddEdge(edges[index], edges[index + 1])
	}

	return g
}
