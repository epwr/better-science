// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
)


func assertEquals(t *testing.T, object_1 any, object_2 any) {

    if object_1 != object_2 {
	t.Errorf("Expected two objects with the same value, got different values")
    }
} 


func assertNodeIsOperationNode(t *testing.T, n kg.Node) {

    if n.GetNodeType() != "OperationNode" {
	t.Errorf("Expected a 'ValueNode', got a '" + n.GetNodeType() + "'.")
    }
    
}

func assertNodeIsValueNode(t *testing.T, n kg.Node) {

    if n.GetNodeType() != "ValueNode" {
	t.Errorf("Expected a 'ValueNode', got a '" + n.GetNodeType() + "'.")
    }
    
}









