// Copywrite 2023 Eric Power - All Rights Reserved
//
// The Nodes enforce the requirement that ValueNodes are only directly connected to OperationNodes and
// vice-versa.
package knowledge_graph

import (
    slog "github.com/epwr/better-science/internal/common/eplog"
    "errors"
)

// Node represents a node in a Graph
type Node interface {
    GetNodeType() string	
    GetName() string
    addSuccessor(Node) error
}

// OperationNode represents an operation node in the graph.
type OperationNode struct {
    Name string
    Predecessors []*ValueNode
    Successor *ValueNode
}

// GetNodeType returns the type of the node.
func (n OperationNode) GetNodeType() string {
    return "OperationNode"
}

// GetName returns the type of the node.
func (n OperationNode) GetName() string {
    return n.Name
}

// addSuccessor adds a successor ValueNode to the source node, and adds the source node as
// a predecessor to the successor node.
func (n OperationNode) addSuccessor(otherNode Node) error {

    valueAddr, ok := otherNode.(ValueNode)
    if !ok {
	msg := "Trying to add the wrong type of Node"
	slog.Error(msg, "source_node", n, "other_node", otherNode)
	return errors.New(msg)
    }
    
    n.Successor = &valueAddr

    return nil
}



// value_type is an enum of the types of value nodes.
type value_type int
const(
    IntType value_type = iota
    FloatType
    EnumType
)

// ValueNode is the description of a potential value.
type ValueNode struct {
    Name string
    Type value_type
    Predecessors []*OperationNode
    Successors []*OperationNode
}

// GetNodeType returns the type of the node.
func (n ValueNode) GetNodeType() string {
    return "ValueNode"
}

// GetName returns the type of the node.
func (n ValueNode) GetName() string {
    return n.Name
}

// addSuccessor adds a successor OperationNode to the source node, and adds the source node as
// a predecessor to the successor node.
func (n ValueNode) addSuccessor(otherNode Node) error {

    return errors.New("Not implemented")
}
