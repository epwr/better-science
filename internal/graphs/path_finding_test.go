package graphs

import "testing"

func TestGetPathBetweeTwoNodes(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)
	g.AddNode("node3", 3)

	g.AddEdge("node1", "node2")
	g.AddEdge("node2", "node3")

	path := g.GetPathBetween("node1", "node2")

	// TODO:
	// - Do I want the path to start with the first node?
	// - What are "intermediary nodes" called?
	assertTrue(t, path[0] == "node1", "The path should start with the initial node.")
	assertTrue(t, path[1] == "node2", "The path should include all intermidiary nodes.")
	assertTrue(t, path[2] == "node3", "The path should end with the final node node.")
}

// TODO: write a test that checks g.GetPathBetween("node1", "node1") when there is a cycle.
// TODO: write a test that checks g.GetPathBetween("node1", "node1") when the node is isolated.
// TODO: write a test that checks g.GetPathBetween("node1", "node1") when the node is connected, but not
// part of a cycle.
