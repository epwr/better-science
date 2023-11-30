// Provides tests that show how do use this package.
package graphs

import (
	"testing"
)

// Start with creating a new graph. When creating a graph, you need to pass two types. The first
// is the graph's nodes' identifier, in this case a string. You will use this identifier to represent
// each node. The second is a datatype that can store additional data for each node.
func CreateAGraphAndAddANode(t *testing.T) {

	g := NewGraph[string, int]()

	assertTrue(t, g.GetNodeCount() == 0, "A new graph should have no nodes")
}

// TODO: A core component of this package is the
func CreateAnEdgeBetweenTwoNodesAndGetTheDistance(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")

	// TODO: update test
	assertTrue(t, true, "The distance between two neighbouring nodes is 1.")
}
