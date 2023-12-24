package knowledge_graphs

import (
	"testing"
)

// INSTRUCTIONS:
//
// Start with creating a new knowledge graph. These knowledge graphs are intended
// to calculate values from probabilistic inputs & operations.
//
// Let's use describing a person's characteristics as an example.
//
// When you set up a knowledge graph, you provide three types:
// 1. The type for the node IDs.
// 2. The Value type which 
// 3. The type for the payload of the Operation nodes.
//
// These provide a flexible way of choosing how to index information, what information
// you want have to evaluate an operation, and how to store the value of a term.
//
// Both Operation and Value nodes have an interface they must satisfy: OperationInterface and
// ValueInterface respectively.
func TestCreateAKnowledgeGraph(t *testing.T) {

	// First, create a knowledge graph.
	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

	// Add Terms to the KG. When you first add a term it
	// is unset. 
	kg.AddTerm("age")
	kg.AddTerm("height")
	kg.AddTerm("weight")

	// Add the Operations you have. They must provide the OperationInterface interface.
	kg.AddOperation(
		"calc_weight_from_height",
		TestOperation{
			calc: "weight",
			reqs: []string{"height"},
			return_value: NewTestValue(150, 0.2),
			return_error: nil,
			return_weight: 1,
		},
	)

	kg.AddOperation(
		"calc_height_from_age",
		TestOperation{
			calc: "height",
			reqs: []string{"age"},
			return_value: NewTestValue(68, 0.2),
			return_error: nil,
			return_weight: 1,
		},
	)

	// Typically, but not always, you'll want to create a copy of the KnowledgeGraph before you start
	// initializing the Terms.
	prepared_kg := kg.Copy()

	prepared_kg.SetTerm("age", NewTestValue(25, 0.2))

	value, error := prepared_kg.Calculate("weight")
	assertTrue(
		t, error == nil,
		"Knowledge Graph should be able to calculate the weight value.",
	)
	assertTrue(
		t, value.value == 150,
		"The Knowledge Graph should have set the final term's value to the provided return_value.",
	)
}
