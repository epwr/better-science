// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph

import (
    slog "github.com/epwr/better-science/internal/common/eplog"
    "errors"
)

// Graph is a graph of Operation and Value Nodes.
//
// Do not instantiate directly, instead use the NewGraph constructor.
type Graph struct {
    value_nodes map[string] ValueNode
    operation_nodes map[string] OperationNode
}

// NewGraph returns a new Graph object that has instantiated maps
// to store the nodes of the graph.
func NewGraph() Graph {
    g := Graph{}
    g.operation_nodes = make(map[string] OperationNode)
    g.value_nodes = make(map[string] ValueNode)
    return g
}

// AddNode adds a Node to the Graph.
func (g Graph) AddNode(node Node) error {

    name := node.GetName()
    existingNode, _ := g.GetNode(name)
    if existingNode != nil {
	return errors.New("A node with the name '" + name + "' already exists.")
    }

    opNode, isOp := node.(OperationNode);
    if isOp {
	g.add_operation_node(name, opNode)
	return nil
    }

    valNode, isVal := node.(ValueNode);
    if isVal {
	g.add_value_node(name, valNode)
	return nil
    }

    slog.Critical("AddNode received a type that implements the Node interface that it is not familiar with.", "node", node)
    return errors.New("AddNode does not know about the " + node.GetNodeType() + " Node type")
}

// GetNode returns the node with the provided name, or nil when the node does not exist.
func (g Graph) GetNode(name string) (Node, error) {

    op_node, op_node_exists := g.operation_nodes[name]
    val_node, val_node_exists := g.value_nodes[name]

    if !(op_node_exists || val_node_exists) {
	return nil, errors.New("No node found with name '" + name + "'.")
    }
    if op_node_exists {
	return op_node, nil
    } else {
	return val_node, nil
    }
}

// CreateEdge takes two Nodes and creates a link between them.
func (g Graph) CreateEdge(sourceNode Node, destinationNode Node) error {

    err := sourceNode.addSuccessor(destinationNode)
    if err != nil {
	return err
    }

    return nil
}


// add_operation_node adds a OperationNode to the Graph, does not check if the node already exists.
func (g  Graph) add_operation_node(name string, node OperationNode) {
    g.operation_nodes[name] = node
}

// add_value_node adds a ValueNode to the Graph, or returns an error if a node already exists
// with that name.
func (g  Graph) add_value_node(name string, node ValueNode) error {

    _, exists := g.value_nodes[name]
    if exists {
	return errors.New("A node with name '" + node.GetName() + "' and type '" + node.GetNodeType() + "' already exists.")
    } else {
	g.value_nodes[name] = node
    }
    return nil
}

