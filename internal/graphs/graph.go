// Copywrite 2023 Eric Power - All Rights Reserved
package graphs

import "errors"

// A Graph is the core datatype of this package. It is undirected and does not
// allow loops. Each node is represented by an identifier and can contain a data
// payload.
//
// Add methods are idempotent.
type Graph[ID comparable, Payload any] struct {
	nodes map[ID]Payload
	edges map[ID][]ID
}

// Create a new Graph object with the provided ID and Payload types.
func NewGraph[ID comparable, Payload any]() Graph[ID, Payload] {
	return Graph[ID, Payload]{nodes: make(map[ID]Payload),
		edges: make(map[ID][]ID),
	}
}

// Add a new node to the Graph.
//
// This function is idempotent. If the node already exists, then this
// will update its payload.
func (g *Graph[ID, Payload]) AddNode(id ID, p Payload) {
	g.nodes[id] = p
}

// Return the number of nodes in the Graph
func (g *Graph[ID, Payload]) GetNodeCount() int {
	return len(g.nodes)
}

// Create an Edge between two nodes. Adding an Edge is idempotent,
// so if the edge already exists then nothing will happen.
//
// If the edge would be a loop, or if the nodes do not exist in the graph,
// then this function returns an error.
func (g *Graph[ID, Payload]) AddEdge(nodeOne ID, nodeTwo ID) error {

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
	g.edges[nodeTwo] = append(g.edges[nodeTwo], nodeOne)

	return nil
}

// Checks if an edge exists between two nodes.
func (g *Graph[ID, Payload]) DoesEdgeExist(nodeOne ID, nodeTwo ID) bool {

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
func (g *Graph[ID, Payload]) DoesNodeExist(nodeID ID) bool {

	_, exists := g.nodes[nodeID]
	return exists

}
