// Copywrite 2023 Eric Power - All Rights Reserved
package kg_engine_test

import (
    "testing"
)

// TestAddNode tests that adding a node to a graph creates a new graph with the node.
func TestAddNode(t *testing.T) {

	graph := Graph{}
	node_name = "example_node"
	node := ValueNode{
		Name: node_name,
		Value: 1.2,
		Confidence: 0.3,
	}

	new_graph = graph.AddNode(node)

	if !new_graph.ContainsValueNode(node_name) {
		t.Errorf("Expected AddNode to create a new graph which contains the node.")
	}
	if graph.ContainsValueNode(node_name) {
		t.Errorf("Expected AddNode to not change the original graph.")
	}
}

