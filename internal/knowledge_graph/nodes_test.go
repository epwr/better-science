// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
)

func TestGetNodeTypeWorksForOperationNode(t *testing.T) {

    node := kg.OperationNode{
	Name: "test_operation",
    }
    
    assertEquals(t, node.GetNodeType(), "OperationNode")
}

func TestGetNodeTypeWorksForValueNode(t *testing.T) {
    
    node := kg.ValueNode{
	Name: "test_node",
	Type: kg.IntType,
    }
    
    assertEquals(t, node.GetNodeType(), "ValueNode")
}







