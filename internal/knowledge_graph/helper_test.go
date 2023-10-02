// Copywrite 2023 Eric Power - All Rights Reserved
//
// Provides a set of helper functions for testing this package
package knowledge_graph_test

import (
    "testing"
    kg "github.com/epwr/better-science/internal/knowledge_graph"
)


func assertEquals(t *testing.T, object_1 any, object_2 any, message string) {
    assertTrue(
	t,
	object_1 == object_2,
	message,
    )
} 

func assertNotEquals(t *testing.T, object_1 any, object_2 any, message string) {
    assertTrue(
	t,
	object_1 != object_2,
	message,
    )
}

func assertNil(t *testing.T, object any, message string) {
    assertTrue(
	t,
	object == nil,
	message,
    )
}

func assertNodeIsOperationNode(t *testing.T, n kg.Node, message string) {
    assertTrue(
	t,
	n.GetNodeType() == "OperationNode",
	message,
    ) 
}

func assertNodeIsValueNode(t *testing.T, n kg.Node, message string) {
    assertTrue(
	t,
	n.GetNodeType() == "ValueNode",
	message,
    )
}

func assertTrue(t *testing.T, value bool, message string) {
    if !value {
	t.Errorf("Assertion Failed: " + message)
    }
}

func assertFalse(t *testing.T, value bool, message string) {
    assertTrue(
	t,
	!value,
	message,
    )
}









