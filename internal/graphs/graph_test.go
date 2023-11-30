package graphs

import (
	"testing"
)

// Start with creating a new graph. When creating a graph, you need to pass two types. The first
// is the graph's nodes' identifier, in this case a string. You will use this identifier to represent
// each node. The second is a datatype that can store additional data for each node.
func TestAddingANode(t *testing.T) {

	g := NewGraph[string, int]()

	g.AddNode("node1", 1)

	assertTrue(t, g.GetNodeCount() == 1, "AddNode does not add exactly one node to Graph.")
}

func TestAddAnEdgeBetweenTwoNodes(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")

	// note: this fails maybe because the g.edges doesn't create slices? So maybe the append's in AddEdge
	// return a nil or something? Idk
	assertEdgeExistsProperly(t, g, "node1", "node2")

}

// test: add an edge is idempotent, add an edge requires that the node exists, no self loops?
