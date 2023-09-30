// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
)



func assertNodeIsValueNode(t *testing.T, n kg.Node) {

    if n.GetNodeType() != "ValueNode" {
	t.Errorf("Expected a 'ValueNode', got a '" + n.GetNodeType() + "'.")
    }
    
}

func TestValueNodeConstructorReturnsAValueNode(t *testing.T) {
    
    node := kg.ValueNode{
	Name: "test_node",
	Type: kg.IntType,
    }
    
    assertNodeIsValueNode(t, node)
}



