// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph


type Graph struct {
	nodes []Node
}

// AddNode returns a graph with the node added, but does not
// alter the original graph.
func (g Graph) AddNode(Node) Graph {

	return g

}

// ContainsValueNode returns whether or not a ValueNode with the provided
// name exists in the graph.
func (g Graph) ContainsValueNode(name string) bool {

	// Iterate through the list and check if any element has the target name
	found := false
	for _, node := range g.nodes {
		if node.GetName() == name {
			found = true
			break // Found the target name, no need to continue searching
		}
	}
	return found
}

