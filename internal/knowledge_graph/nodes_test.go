// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
)



func assertNodeIsNotNil(t *testing.T, n kg.Node) {

    if n == nil {
	t.Errorf("The node should not be `nil`, but it is")
    }
    
}

func TestValueNodesConstructorWorks(t *testing.T) {
    
    node := kg.ValueNode{
	Name: "test_node",
	Type: kg.IntType,
    }
    
    assertNodeIsNotNil(t, node)
}

