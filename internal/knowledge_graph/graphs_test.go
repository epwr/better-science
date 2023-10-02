// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
    "fmt"
)


func TestAddingAOperationNodeToAGraph(t *testing.T) {

    graph := kg.NewGraph()
    node := kg.OperationNode{
	Name: "node name",
    }

    err := graph.AddNode(node)
    assertTrue(t, err == nil, "The node should be added to the graph.")
}

func TestErrorWhenAddingTwoOperationNodesWithTheSameName(t *testing.T) {

    graph := kg.NewGraph()
    name := "generic node name"
    first_node := kg.OperationNode{
	Name: name,
    }
    second_node := kg.OperationNode{
	Name: name,
    }

    first_add_status := graph.AddNode(first_node)
    second_add_status := graph.AddNode(second_node)

    assertTrue(t, first_add_status == nil, "The first node should be added to the graph.")
    assertFalse(t, second_add_status == nil, "The second node should be added to the graph.")
}

func TestAddingAValueNodeToAGraph(t *testing.T) {

    graph := kg.NewGraph()
    node := kg.ValueNode{
	Name: "node name",
    }

    err := graph.AddNode(node)
    assertTrue(t, err == nil, "The node should be added to the graph.")
}

func TestErrorWhenAddingTwoValueNodesWithTheSameName(t *testing.T) {

    graph := kg.NewGraph()
    name := "generic node name"
    first_node := kg.ValueNode{
	Name: name,
    }
    second_node := kg.ValueNode{
	Name: name,
    }

    first_add_status := graph.AddNode(first_node)
    second_add_status := graph.AddNode(second_node)

    assertTrue(t, first_add_status == nil, "The first node should be added to the graph.")
    assertFalse(t, second_add_status == nil, "The second node should be added to the graph.")
}

func TestErrorWhenAddingTwoNodesWithTheSameName(t *testing.T) {

    graph := kg.NewGraph()
    name := "generic node name"
    first_node := kg.ValueNode{
	Name: name,
    }
    second_node := kg.OperationNode{
	Name: name,
    }

    first_add_status := graph.AddNode(first_node)
    second_add_status := graph.AddNode(second_node)

    assertTrue(t, first_add_status == nil, "The first node should be added to the graph.")
    assertFalse(t, second_add_status == nil, "The second node should be added to the graph.")    
}


func TestCreatingAnEdge(t *testing.T) {

    graph := kg.NewGraph()
    valNode := kg.ValueNode{
	Name: "value node",
    }
    opNode := kg.OperationNode{
	Name: "operation node",
    }
    graph.AddNode(valNode)
    graph.AddNode(opNode)

    fmt.Printf("Pre Change:\n")
    fmt.Printf("Opr Node -> Address: %p \n", &opNode)
    fmt.Printf("Opr Node -> Successor: %p \n", opNode.Successor)
    fmt.Printf("Val Node -> Address: %p \n", &valNode)
    
    err := graph.CreateEdge(opNode, valNode)

    fmt.Printf("Post Change:\n")
    fmt.Printf("Opr Node -> Address: %p \n", &opNode)
    fmt.Printf("Opr Node -> Successor: %p \n", opNode.Successor)
    fmt.Printf("Val Node -> Address: %p \n", &valNode)
    
    assertTrue(t, err == nil, "Creating the edge should have been successful.")
    assertTrue(t, opNode.Successor == &valNode, "The ValueNode's address should be set as the OperationNode's successor." )

    hasCorrectPredecessor := len(valNode.Predecessors) == 1 && valNode.Predecessors[0] == &opNode 
    assertTrue(t, hasCorrectPredecessor, "The OperationNode's address should be set as the ValueNode's predecessor.")
  
}

// func TestErrorWhenCreatingAnEdgeBetweenTwoNodesOfTheSameType(t *testing.T) {

//     graph := kg.NewGraph()
//     nodeOne := kg.ValueNode{
// 	Name: "value node 1",
//     }
//     nodeTwo := kg.ValueNode{
// 	Name: "value node 2",
//     }
    
//     graph.AddNode(nodeOne)
//     graph.AddNode(nodeTwo)

//     err := graph.CreateEdge(nodeOne, nodeTwo)

//     assertNil(t, err, "The edge should not have been created, as it was between two value nodes.")
// }

