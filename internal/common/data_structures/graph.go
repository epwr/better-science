// Copywrite 2023 Eric Power - All Rights Reserved
package data_structures

import "errors"

// A DiGraph is the core datatype of this package. It is a directed graph that does not
// allow loops. Each node is represented by an identifier and can contain a data
// payload.
//
// Add methods are idempotent.
type DiGraph[ID comparable, Payload any] struct {
	nodes map[ID]Payload
	edges map[ID][]ID
}

// Create a new Graph object with the provided ID and Payload types.
func NewDiGraph[ID comparable, Payload any]() DiGraph[ID, Payload] {
	return DiGraph[ID, Payload]{
		nodes: make(map[ID]Payload),
		edges: make(map[ID][]ID),
	}
}

// Add a new node to the DiGraph.
//
// This function is idempotent. If the node already exists, then this
// will update its payload.
func (g *DiGraph[ID, Payload]) AddNode(id ID, p Payload) {
	g.nodes[id] = p
}

// Get a slice of the IDs of all nodes
func (g *DiGraph[ID, Payload]) GetNodes() []ID {

	keys := make([]ID, 0, len(g.nodes))
	for key := range g.nodes {
		keys = append(keys, key)
	}
	return keys
}

// Return the payload of the node.
func (g *DiGraph[ID, Payload]) GetNode(id ID) Payload{
	return g.nodes[id]
}

// Return the number of nodes in the Graph
func (g *DiGraph[ID, Payload]) GetNodeCount() int {
	return len(g.nodes)
}

// Create a directed edge between two nodes. Adding an edge is idempotent,
// so if the edge already exists then nothing will happen.
//
// If the edge would be a loop, or if the nodes do not exist in the graph,
// then this function returns an error.
func (g *DiGraph[ID, Payload]) AddEdge(nodeOne ID, nodeTwo ID) error {

	if g.DoesEdgeExist(nodeOne, nodeTwo) {
		return nil
	}

	if nodeOne == nodeTwo {
		return errors.New("Tried to create a self-loop")
	}

	if !g.DoesNodeExist(nodeOne) || !g.DoesNodeExist(nodeTwo) {
		return errors.New("Tried to create an edge with a node that does not exist")
	}

	g.edges[nodeOne] = append(g.edges[nodeOne], nodeTwo)

	return nil
}

// Checks if an edge exists between two nodes.
func (g *DiGraph[ID, Payload]) DoesEdgeExist(nodeOne ID, nodeTwo ID) bool {

	nodeOneNeighbours := g.edges[nodeOne]
	exists := false
	for _, neighbour := range nodeOneNeighbours {
		if neighbour == nodeTwo {
			exists = true
		}
	}
	return exists

}

// Checks if a node exists in the graph.
func (g *DiGraph[ID, Payload]) DoesNodeExist(nodeID ID) bool {

	_, exists := g.nodes[nodeID]
	return exists

}

// Returns the neighbouring nodes' IDs.
func (g *DiGraph[ID, Payload]) GetNeighbours(node ID) []ID {

	return g.edges[node]

}
