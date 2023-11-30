// Copywrite 2023 Eric Power - All Rights Reserved
package graphs

// A Graph is the core datatype of this package. It stores generic nodes
// (defined as an ID and a Payload) and undirected edges between the nodes.
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

// Add a new node to the Graph. This function is idempotent. If the
// node already exists, then update its payload.
func (g *Graph[ID, Payload]) AddNode(id ID, p Payload) {
	g.nodes[id] = p
}

// Return the number of nodes in the Graph
func (g *Graph[ID, Payload]) GetNodeCount() int {
	return len(g.nodes)
}

// Create an Edge between two nodes
func (g *Graph[ID, Payload]) AddEdge(nodeOne ID, nodeTwo ID) {
	g.edges[nodeOne] = append(g.edges[nodeOne], nodeTwo)
	g.edges[nodeTwo] = append(g.edges[nodeTwo], nodeOne)
}
