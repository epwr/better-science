package knowledge_graphs

import (
	"errors"
	"fmt"
	"cmp"
)

type OperationInterface[ID cmp.Ordered, Value ValueInterface] interface {
	Calculates() ID // The Term that the Operation calculates
	Requires() []ID // The Terms that the Operation relies on
	Run(map[ID]Value) (Value, error) // Runs the Operation to calculate the new value
}

// An interface that defines a Value object. For this knowledge graph, it is important that a Value can
// provide a weight value. This allows the Knowledge Graph to select between Operations that provide
// differnent weights.
type ValueInterface interface {
	GetWeight() float64 // The weight of the value.
}

type KnowledgeGraph[ID cmp.Ordered, Value ValueInterface, Operation OperationInterface[ID, Value]] struct {
	terms map[ID] Value
	setTerms map[ID] bool // Tracks which terms have been set
	operations map[ID] Operation
	edges map[ID][]ID
}

// Returns a new, empty, knowledge graph.
func NewKnowledgeGraph[ID cmp.Ordered, Value ValueInterface, Operation OperationInterface[ID, Value]]() KnowledgeGraph[ID, Value, Operation] {

	return KnowledgeGraph[ID, Value, Operation] {
		terms: make(map[ID] Value),
		operations: make(map[ID] Operation),
		edges: make(map[ID][]ID),
		setTerms: make(map[ID] bool),
	}
}

// Adds an empty Value to the KnowledgeGraph. If the provided ID already exists (either as a Value or
// as an Operation), returns an error
func (kg *KnowledgeGraph[ID, Value, Operation]) AddTerm(id ID) error {

	if kg.isIDUsed(id) {
		return errors.New("Tried to add a Value with an ID that already exists in the Knowledge Graph")
	}

	kg.terms[id] = *new(Value) // Adding a Value & setting the value of a Value are distinct steps.
	return nil
}

// Sets the value of an existing Term in the KnowledgeGraph.
func (kg *KnowledgeGraph[ID, Value, Operation]) SetTerm(id ID, value Value) error {

	_, termExists := kg.terms[id]
	if !termExists {
		return errors.New(fmt.Sprintf("No term with the ID '%v'", id))
	}

	kg.terms[id] = value
	kg.setTerms[id] = true
	return nil
}

// Returns the value of a Term in the KnowledgeGraph whose Value has been set.
func (kg *KnowledgeGraph[ID, Value, Operation]) GetTerm(id ID) (Value, error) {

	_, isSet := kg.setTerms[id]
	if !isSet {
		return *new(Value), errors.New("Term is not set.")
	}

	return kg.terms[id], nil
}


// Adds an Operation to the KnowledgeGraph, and adds edges based on the Operation's Calculates & Requires
// methods.
//
// If the provided ID already exists (either as a Value or as an Operation), returns an error.
func (kg *KnowledgeGraph[ID, Value, Operation]) AddOperation(opID ID, op Operation) error {

	if kg.isIDUsed(opID) {
		return errors.New("Tried to add an Operation with an ID that already exists in the Knowledge Graph")
	}

	// Checks
	source := op.Calculates()
	targets := op.Requires()
	for _, term := range targets {
		_, exists := kg.terms[term]
		if !exists {
			termStr := fmt.Sprintf("%v", term)
			return errors.New("Operation requires a Value that does not exist: " + termStr)
		}
		if term == source {
			return errors.New("An Operation cannot calcuate a term it requires.")
		}
	}
	_, sourceExists := kg.terms[source]
	if !sourceExists {
		termStr := fmt.Sprintf("%v", source)
		return errors.New("Operation calculates a Value that does not exist: " + termStr)
	}
	
	// Add outgoing edges
	kg.edges[opID] = make([]ID, 0)
	for _, target := range targets {
		kg.edges[opID] = append(kg.edges[opID], target)
	}

	// Add incoming edge
	_, exists := kg.edges[source]
	if !exists {
		kg.edges[source] = make([]ID, 1)
	}
	kg.edges[source] = append(kg.edges[source], opID)
	
	kg.operations[opID] = op
	return nil
}

// Returns the value stored in the provided Operation
//
// Deprecated: as the KG now directly calls Calculate on an Operation, this
// method is no longer needed. 
func (kg *KnowledgeGraph[ID, Value, Operation]) GetOperation(id ID) (Operation, error) {

	operation, exists := kg.operations[id]
	if !exists {
		return operation, errors.New("Provided ID is not an operation in the KnowledgeGraph.")
	}
	return operation, nil
	
}


// Returns a deep copy of the KnowledgeGraph
func (kg *KnowledgeGraph[ID, Value, Operation]) Copy() KnowledgeGraph[ID, Value, Operation] {

	newKg := NewKnowledgeGraph[ID, Value, Operation]()

	for id, term := range kg.terms {
		newKg.terms[id] = term
	}

	for id, op := range kg.operations {
		newKg.operations[id] = op
	}
	
	for id, edge := range kg.edges {
		newKg.edges[id] = edge
	}

	for id, setFlag := range kg.setTerms {
		newKg.setTerms[id] = setFlag
	}

	return newKg
}



// Returns a slice of all the Term IDs in the Knowledge Graph
func (kg *KnowledgeGraph[ID, Value, Operation]) getTermIDs() []ID {

	ids := make([]ID, 0, len(kg.terms))
	for key := range kg.terms {
		ids = append(ids, key)
	}
	
	return ids
}

// Returns a slice of all the Operation IDs in the Knowledge Graph
func (kg *KnowledgeGraph[ID, Value, Operation]) getOperationIDs() []ID {

	ids := make([]ID, 0, len(kg.operations))
	for key := range kg.operations {
		ids = append(ids, key)
	}
	
	return ids
}

// Check if the ID represents a Value/Operation or not.
func (kg *KnowledgeGraph[ID, Value, Operation]) isIDUsed(id ID) bool {

	_, termExists := kg.terms[id]
	_, operationExists := kg.operations[id]

	return termExists || operationExists
}


