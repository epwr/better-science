package graphs

import (
	"errors"
	
	"github.com/epwr/better-science/internal/common/data_structures"
)

func (g *Graph[ID, Payload]) GetShortestPathBetween(startNode ID, endNode ID) ([]ID, error) {

	parentNodes := make(map[ID]ID)
	visitedNodes := make(map[ID]bool)
	toVisit := data_structures.NewQueue[ID]()
	
	toVisit.Push(startNode)
	parentNodes[startNode] = startNode 

	for !toVisit.IsEmpty() {
		
		node, _ := toVisit.Pop()

		if node == endNode {
			return buildPath[ID](startNode, node, parentNodes), nil
		}
		
		if !visitedNodes[node] {
			neighbours := g.GetNeighbours(node)
			visitedNodes[node] = true

			for i := 0; i < len(neighbours); i++ {
				_, exists := parentNodes[neighbours[i]]
				if !exists {
					parentNodes[neighbours[i]] = node
				}
				toVisit.Push(neighbours[i])
			}
		}
		
	}
	
	return []ID{}, errors.New("No path could be found")
}

// prepend an element onto a slice. 
func prepend[ID comparable](element ID, slice []ID) []ID {
	newSlice := make([]ID, len(slice) + 1)
	newSlice[0] = element
	copy(newSlice[1:], slice)
	return newSlice
}

// build a path from the start node to the
// end node.
func buildPath[ID comparable](startNode ID, endNode ID, parentNodes map[ID]ID) []ID {
	
	path := []ID{endNode}
	node := endNode
	for node != startNode {
		path = prepend(parentNodes[node], path)
		node = parentNodes[node]
	}

	return path
}


