package data_structures

import (
	"testing"
	"fmt"
)

func TestAddingNodes(t *testing.T){

	addingNodeTests := []struct{nodes []string; result []string}{
		{[]string{"A1", "A2", "A3"}, []string{"A1", "A2", "A3"}},
		{[]string{"A1", "A1", "A1"}, []string{"A1"}},
		{[]string{}, []string{}},
	}
	
	for _, testCase := range addingNodeTests {
		g := NewDiGraph[string, int]()

		for _, node := range testCase.nodes {
			g.AddNode(node, 1)
		}

		inputStr := fmt.Sprintf("%v", testCase.nodes)
		expectedStr := fmt.Sprintf("%v", testCase.result)
		gotStr := fmt.Sprintf("%v", g.GetNodes())
		
		assertSlicesAreSetwiseEqual(
			t,
			g.GetNodes(),
			testCase.result,
			"The input '" + inputStr + "' should result in '" + expectedStr + "' but got '" + gotStr + "'.",
		)
	}
}

func TestAddingEdges(t *testing.T) {

	// Every test case will have the same nodes
	nodes := []string{"A", "B", "C"}

	// Provide a slice of edges (source, dest, source, dest...) and
	// then a set of tests, which test is a given given edge exists.
	testCases := []struct{
		edges []string;
		tests []struct{source string; dest string; result bool};
	}{
		{
			[]string{"A", "B"},
			[]struct{source string; dest string; result bool}{
				{"A", "B", true},
				{"B", "A", false},
			},
		},
		{
			[]string{},
			[]struct{source string; dest string; result bool}{
				{"A", "B", false},
				{"B", "A", false},
			},
		},
		{
			[]string{
				"A", "B",
				"B", "C",
			},
			[]struct{source string; dest string; result bool}{
				{"A", "B", true},
				{"B", "A", false},
				{"B", "C", true},
				{"A", "C", false},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		g := NewDiGraph[string, int]()
		for _, node := range nodes {
			g.AddNode(node, 1)
		}

		for i := 0; i < len(testCase.edges); i = i+2 {

			source := testCase.edges[i]
			dest := testCase.edges[i+1]

			g.AddEdge(source, dest)
		}

		for testIndex, test := range testCase.tests {

			result := g.DoesEdgeExist(test.source, test.dest)

			if (result != test.result) {
				
				t.Errorf(
					"Case %d, test %d: Checking edge from %s to %s returned %t. Expected %t",
					caseIndex,
					testIndex,
					test.source,
					test.dest,
					result,
					test.result,
				)
			}
		}
	}
}

func TestGetNodeCountWorksProperly(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	assertTrue(t, g.GetNodeCount() == 2, "GetNodeCount should return the number of nodes in the graph.")
}

func TestAddEdgeIsIdempotent(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")
	g.AddEdge("node1", "node2")

	assertDirectedEdgeExistsExactlyOnce(t, g, "node1", "node2")
}

func TestAddEdgeRequiresTheNodesExist(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)

	error := g.AddEdge("node1", "node2")

	assertError(t, error, "AddEdge should not add an edge between two nodes that do not exist")
}

func TestAddEdgeDoesNotAllowCreatingSelfLoops(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)

	error := g.AddEdge("node1", "node1")

	assertError(t, error, "AddEdge should not allow creating a self-loop.")
}

func TestDoesEdgeExistReturnsTrueWhenEdgeExists(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	g.AddEdge("node1", "node2")

	assertTrue(t, g.DoesEdgeExist("node1", "node2"), "DoesEdgeExist should return false when the edge exists")
}

func TestDoesEdgeExistReturnsFalseWhenEdgeDoesNotExist(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)

	assertTrue(t, !g.DoesEdgeExist("node1", "node2"), "DoesEdgeExist should return false when the edge does not exist")
}

func TestDoesNodeExistReturnsTrueWhenNodeExists(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)

	found := g.DoesNodeExist("node1")

	assertTrue(t, found, "DoesNodeExist did not find an existing node")

}

func TestDoesNodeExistReturnsFalseWhenNodeDoesNotExist(t *testing.T) {

	g := NewDiGraph[string, int]()

	found := g.DoesNodeExist("node1")

	assertTrue(t, !found, "DoesNodeExist should return fase when the node does not exist")
}

func TestGetNodeReturnsTheNodePayload(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)

	payload := g.GetNode("node1")
	assertTrue(t, payload == 1, "GetNode should return the Payload of the node")
}

func TestNodesAreNotPassedByReference(t *testing.T) {

	type NodeData struct {
		data int
	}

	g := NewDiGraph[string, NodeData]()
	node := NodeData{data: 4}
	g.AddNode("node_id", node)
	node_copy := g.GetNode("node_id")
	node_copy_2 := g.GetNode("node_id")
	
	node.data = 3
	node_copy_2.data = 2
	assertTrue(t, node_copy.data == 4, "Adding a node should create a copy of the node")
}

func TestGetNeighboursReturnsAllNeighbourNodes(t *testing.T) {

	g := NewDiGraph[string, int]()
	g.AddNode("node1", 1)
	g.AddNode("node2", 2)
	g.AddNode("node3", 3)
	g.AddNode("node4", 4)

	g.AddEdge("node1", "node2")
	g.AddEdge("node1", "node3")		
	g.AddEdge("node2", "node3")

	neighbours := g.GetNeighbours("node1")
	
	assertTrue(t, len(neighbours) == 2, "node1 should have 2 neighbours")
	if len(neighbours) == 2 {
		hasNodeTwo := neighbours[0] == "node2" || neighbours[1] == "node2"
		hasNodeThree := neighbours[0] == "node3" || neighbours[1] == "node3"
		assertTrue(t, hasNodeTwo && hasNodeThree, "node1 should have node2 and node3 as neighbours.")
	}
}


