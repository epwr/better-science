package graphs

import (
	"fmt"
	"testing"
)


func TestGetShortestPathWhereTheStartAndEndNodesAreTheSame(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)

	path, _ := g.GetShortestPathBetween("node1", "node1")
	assertTrue(t, len(path) == 1, "The path should only contain one node.")
	
}


func TestGetPathBetweenTwoNodes(t *testing.T) {

	g := NewGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)
	g.AddNode("node3", 3)

	g.AddEdge("node1", "node2")
	g.AddEdge("node2", "node3")

	path, _ := g.GetShortestPathBetween("node1", "node3")

	fmt.Println(path)

	assertSliceIsEqual(
		t,
		path,
		[]string{"node1", "node2", "node3"},
		"The path should contain a path from start to end nodes.",
	)
}


func TestBuildPathProvidesCorrectPath(t *testing.T) {

	parentNodes := map[string]string{
		"end": "mid3",
		"mid3": "mid2",
		"mid2": "mid1",
		"mid1": "start",
	}

	path := buildPath("start", "end", parentNodes)

	assertSliceIsEqual(
		t,
		path,
		[]string{"start", "mid1", "mid2", "mid3", "end"},
		"The path should list the nodes in reverse order from what is in the parentNodes map.",
	)
}

func TestNoPathCanBeFoundBetweenTwoNodesInDisjointGraphs(t *testing.T) {

	g := initGraph[string, int](
		[]string{"A1", "A2", "A3", "B1", "B2"},
		[]string{
			"A1", "A2",
			"A2", "A3",
			"B1", "B2",
		},
	)

	_, error := g.GetShortestPathBetween("A1", "B2")

	assertError(t, error, "No path should be found between two nodes in disjoint graphs.")
}


func TestAPathCanBeFoundEvenWhenACycleExists(t *testing.T) {

	g := initGraph[string, int](
		[]string{"A1", "A2", "A3", "A4", "A5"},
		[]string{
			"A1", "A2",
			"A2", "A3",
			"A3", "A4",
			"A4", "A2",
			"A4", "A5",
			
		},
	)

	path, _ := g.GetShortestPathBetween("A1", "A5")

	assertSliceIsEqual(
		t,
		path,
		[]string{"A1", "A2", "A4", "A5"},
		"GetShortestPathBetween should find the shortest path, even when there is a cycle in the graph.",
	)
}

// TODO: write a test that checks a path when there is a cycle.
// TODO: write a test that checks g.GetPathBetween("node1", "node1") when the node is isolated.
// TODO: write a test that checks g.GetPathBetween("node1", "node1") when the node is connected, but not
// part of a cycle.
