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

	assertTrue(t, g.GetNodeCount() == 1, "AddNode should add exactly one node to Graph.")
}

func TestAddAnEdgeBetweenTwoNodes(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")

	assertEdgeExistsExactlyOnce(t, g, "node1", "node2")
}

func TestGetNodeCountWorksProperly(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	assertTrue(t, g.GetNodeCount() == 2, "GetNodeCount should return the number of nodes in the graph.")
}

func TestAddEdgeIsIdempotent(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")
	g.AddEdge("node1", "node2")

	assertEdgeExistsExactlyOnce(t, g, "node1", "node2")
}

func TestAddEdgeRequiresTheNodesExist(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)

	error := g.AddEdge("node1", "node2")

	assertError(t, error, "AddEdge should not add an edge between two nodes that do not exist")
}

func TestAddEdgeDoesNotAllowCreatingSelfLoops(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)

	error := g.AddEdge("node1", "node1")

	assertError(t, error, "AddEdge should not allow creating a self-loop.")
}

func TestDoesEdgeExistReturnsTrueWhenEdgeExists(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")

	assertTrue(t, g.DoesEdgeExist("node1", "node2"), "DoesEdgeExist should return false when the edge exists")
}

func TestDoesEdgeExistReturnsFalseWhenEdgeDoesNotExist(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	assertTrue(t, !g.DoesEdgeExist("node1", "node2"), "DoesEdgeExist should return false when the edge does not exist")
}

func TestDoesNodeExistReturnsTrueWhenNodeExists(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)

	found := g.DoesNodeExist("node1")

	assertTrue(t, found, "DoesNodeExist did not find an existing node")

}

func TestDoesNodeExistReturnsFalseWhenNodeDoesNotExist(t *testing.T) {

	g := NewGraph[string, int]()

	found := g.DoesNodeExist("node1")

	assertTrue(t, !found, "DoesNodeExist should return fase when the node does not exist")
}
