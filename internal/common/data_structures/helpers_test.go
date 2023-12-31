// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package data_structures

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
				message + "[For element '" + elemStr + "', found " +  countOneStr +
                " and " + countTwoStr + " elements in each slice respectively.]",
			)
		}
	}
}


// Assert that the graph.edge adjacency map contains exactly one reference
// from nodeOne to nodeTwo and exactly one reference from nodeTwo to nodeOne
func assertDirectedEdgeExistsExactlyOnce[ID comparable, Payload any](
	t *testing.T,
	g DiGraph[ID, Payload],
	nodeOne ID,
	nodeTwo ID,
) {

	t.Helper()
	
	nodeOneNeighbours, foundOutgoingEdges := g.edges[nodeOne]
	assertTrue(t, foundOutgoingEdges, "A node was not found in the graph's edges.")
	
	_, foundNodeOne := g.nodes[nodeOne]
	_, foundNodeTwo := g.nodes[nodeTwo]
	assertTrue(t, foundNodeOne && foundNodeTwo, "Both nodes should exist in the graph's nodes.")

	countNodeTwoReferences := 0
	for _, neighbour := range nodeOneNeighbours {
		if neighbour == nodeTwo {
			countNodeTwoReferences++
		}
	}

	nodeOneStr := fmt.Sprintf("%v", nodeOne)
	assertTrue(
		t,
		countNodeTwoReferences == 1,
		nodeOneStr + "'s edges did not contain a reference to the other node exactly once.",
	)
}

// Initialize a DiGraph
func initGraph[ID comparable, Payload any](nodes []ID, edges []ID) DiGraph[ID, Payload] {

	g := NewDiGraph[ID, Payload]()
	for index := 0; index < len(nodes); index++ {
		g.AddNode(nodes[index], *new(Payload))
	}

	for index := 0; index < len(edges); index = index + 2 {
		g.AddEdge(edges[index], edges[index + 1])
	}

	return g
}


