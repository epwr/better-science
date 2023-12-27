package knowledge_graphs

import (
	"fmt"
	"reflect"
	"testing"
)


func TestAddingTermsToKnowledgeGraph(t *testing.T){

	addTermsTests := []struct{terms []string; result []string; expectError bool}{
		{[]string{"T1", "T2", "T3"}, []string{"T1", "T2", "T3"}, false},
		{[]string{"T1", "T1", "T1"}, []string{"T1"}, true},
		{[]string{}, []string{}, false},
	}
	
	for _, testCase := range addTermsTests {
		kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

		for index, term := range testCase.terms {
			error := kg.AddTerm(term)
			if error != nil && !testCase.expectError {
				t.Errorf(
					"Unexpected error while adding term '%s' (index %d) in testCase: %+v",
					term, index, testCase,
				)
			}
		}

		termIDs := kg.getTermIDs()
		
		inputStr := fmt.Sprintf("%v", testCase.terms)
		expectedStr := fmt.Sprintf("%v", testCase.result)
		gotStr := fmt.Sprintf("%v", termIDs)
		
		assertSlicesAreSetwiseEqual(
			t,
			termIDs,
			testCase.result,
			"The input '" + inputStr + "' should result in '" + expectedStr + "' but got '" + gotStr + "'.",
		)
	}
}


func TestAddingOperationsToKnowledgeGraph(t *testing.T){

	addOperationsTests := []struct{operations []string; result []string; expectError bool}{
		// Operations are added appropriately & with no errors.
		{[]string{"O1", "O2", "O3"}, []string{"O1", "O2", "O3"}, false},
		// Duplicates are not added.
		{[]string{"O1", "O1", "O1"}, []string{"O1"}, true},
		// GetOperations works when there is no operations.
		{[]string{}, []string{}, false},
	}
	
	for _, testCase := range addOperationsTests {
		kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
		kg.AddTerm("T1")
		kg.AddTerm("T2")

		for index, operationID := range testCase.operations {
			error := kg.AddOperation(operationID, MakeExampleTestOperation())
			if error != nil && !testCase.expectError {
				t.Errorf(
					"Unexpected error while adding operation '%s' (index %d) in testCase: %+v",
					operationID, index, testCase,
				)
			}
		}

		opIDs := kg.getOperationIDs()
		
		inputStr := fmt.Sprintf("%v", testCase.operations)
		expectedStr := fmt.Sprintf("%v", testCase.result)
		gotStr := fmt.Sprintf("%v", opIDs)
		
		assertSlicesAreSetwiseEqual(
			t,
			opIDs,
			testCase.result,
			"The input '" + inputStr + "' should result in '" + expectedStr + "' but got '" + gotStr + "'.",
		)
	}
}

func TestGetOperationPayload(t *testing.T) {

	type Op struct {
		id string
		payload TestOperation
	}
	
	getOpsTests := []struct{operations []Op; getID string; expectOp TestOperation; expectError bool}{
		{
			operations: []Op{{id: "A", payload: MakeExampleTestOperation()}},
			getID: "A", expectOp: MakeExampleTestOperation(), expectError: false,
		},
		{
			operations: []Op{
				{id: "A", payload: NewTestOperation("T1", "T2")},
				{id: "A", payload: NewTestOperation("T1", "T3")},
				{id: "A", payload: NewTestOperation("T1", "T4")},
			},
			getID: "A", expectOp: NewTestOperation("T1", "T2"), expectError: false,
		},
		{
			operations: []Op{
				{id: "A", payload: NewTestOperation("T1", "T2")},
				{id: "B", payload: NewTestOperation("T1", "T3")},
			},
			getID: "C", expectOp: TestOperation{}, expectError: true,
		},
		{
			operations: []Op{
				{id: "A", payload: NewTestOperation("T1", "T2")},
				{id: "B", payload: NewTestOperation("T1", "T3", "T2")},
			},
			getID: "C", expectOp: TestOperation{}, expectError: true,
		},
	}

	for _, testCase := range getOpsTests {

		t.Logf("For testCase: %+v", testCase)
		
		kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
		for _, term := range []string{"T1", "T2", "T3", "T4"} {
			kg.AddTerm(term)
		}

		for _, operation := range testCase.operations {
			error := kg.AddOperation(operation.id, operation.payload)
			t.Logf(
				"Adding ID '%s' with payload '%+v'. Result: %v",
				operation.id, operation.payload, error,
			)
		}

		result, err := kg.GetOperation(testCase.getID)
		
		assertTrue(
			t, !((err == nil) && testCase.expectError),
			"Expected an error, but no error was returned.",
		)
		assertTrue(
			t, !((err != nil) && !testCase.expectError),
			"Expected no error, but an error was returned.",
		)
		assertSlicesAreSetwiseEqual(
			t, result.reqs, testCase.expectOp.reqs,
			"GetOperation did not return the expected value.",
		)
	}
}

func TestAddingOperationsCreatesAppropriateEdges(t *testing.T) {

	testSteps := []struct{
		opID string
		operation TestOperation
		addOperationFails bool
		checkEdges []struct{source string; targets []string}
	}{
		{ // Base case
			opID: "O1",
			operation: NewTestOperation("T1", "T2", "T3"),
			addOperationFails: false,
			checkEdges: []struct{source string; targets []string}{
				{source: "T1", targets: []string{"O1"}},
				{source: "O1", targets: []string{"T2", "T3"}},
			},
		},
		{ // An op cannot be added twice
			opID: "O1",
			operation: NewTestOperation("T1", "T2", "T3"),
			addOperationFails: true,
		},
		{ // A term can be calculated multiple ops
			opID: "O2",
			operation: NewTestOperation("T1", "T3"),
			addOperationFails: false,
			checkEdges: []struct{source string; targets []string}{
				{source: "T1", targets: []string{"O1", "O2"}},
				{source: "O2", targets: []string{"T3"}},
			},
		},
		{ // An op can require multiple terms
			opID: "O3",
			operation: NewTestOperation("T1", "T2", "T3", "T4"),
			addOperationFails: false,
			checkEdges: []struct{source string; targets []string}{
				{source: "T1", targets: []string{"O3"}},
				{source: "O3", targets: []string{"T2", "T3", "T4"}},
			},
		},
		{ // An op cannot require another op
			opID: "O4",
			operation: NewTestOperation("T1", "O2"),
			addOperationFails: true,
		},
		{ // An op cannot calculate & require the same term.
			opID: "O5",
			operation: NewTestOperation("T2", "T2"),
			addOperationFails: true,
		},
	}

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

    kg.AddTerm("T1")
    kg.AddTerm("T2")
    kg.AddTerm("T3")
	kg.AddTerm("T4")

    for index, step := range testSteps {

		t.Logf(
			"Test #%d: Adding Operation '%s' -> operation: %+v, knowledge graph: %+v",
			index + 1, step.opID, step.operation, kg,
		)
		err := kg.AddOperation(step.opID, step.operation)
		if step.addOperationFails {
			assertError(t, err, "Adding Operation should have failed.")
		} else {
			for _, edgeCheck := range step.checkEdges {
				for _, target := range edgeCheck.targets {
					t.Logf("Checking Edge: %s -> %s", edgeCheck.source, target)
					targets, _ := kg.edges[edgeCheck.source]
					assertElementIsInSlice(
						t, target, targets,
						"The Target vertex was not found in the source's outgoing vertices.",
					)
				}
			}
		}
	}
}

func TestAddOperationIsAtomic(t *testing.T) {

	// No edges/ids should be added:
	// - If id exists
	// - if caculates/requires terms do not exists
	// - if op requires another op
	// - calculates a term that does not exist

	testCases := []struct{id string; operation TestOperation}{
		{ // Id already exists
			id: "O0",
			operation: NewTestOperation("T1", "T2"),
		},
		{ // Required Terms do not exist
			id: "O1",
			operation: NewTestOperation("T1", "TX"),
		},
		{ // Calculated Term does not exist
			id: "O2",
			operation: NewTestOperation("TX", "T2"),
		},
		{ // Op requires another Op
			id: "O3",
			operation: NewTestOperation("T1", "O0"),
		},
	}
	
	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

    kg.AddTerm("T1")
    kg.AddTerm("T2")
    kg.AddOperation("O0", NewTestOperation("T1", "T2"))

	for testNumber, testCase := range testCases {
		t.Logf(
			"Test Case #%d: Adding Operation %+v to Knowledge Graph %+v",
			testNumber, testCase.operation, kg,
		)

		beforeTerms := kg.terms
		beforeOperations := kg.operations
		beforeEdges := kg.edges

		error := kg.AddOperation(testCase.id, testCase.operation)
		assertError(t, error, "Adding operation should have failed.")

		afterTerms := kg.terms
		afterOperations := kg.operations
		afterEdges := kg.edges

		assertTrue(t, reflect.DeepEqual(beforeTerms, afterTerms), "The Terms should not have changed")
		assertTrue(t, reflect.DeepEqual(beforeEdges, afterEdges), "The Edges should not have changed")
		assertTrue(t, reflect.DeepEqual(beforeOperations, afterOperations), "The Operations should not have changed")
	}
}

func TestSetTerm(t *testing.T) {

	testCases := []struct{id string; value TestValue; expectError bool}{
		{ // Base Case (should pass)
			id: "T1",
			value: NewTestValue(5, 1.0),
			expectError: false,
		},
		{ // Term is already set.
			id: "T1",
			value: NewTestValue(5, 1.0), 
			expectError: false,
		},
		{ // Term ID does not exist
			id: "T0",
			value: NewTestValue(5, 1.0),
			expectError: true,
		},
		{ // empty Term value should work as the KG does not know when/if that's inappropriate.
	      // NOTE: this test requires that the value is set to something else before setting it to this value.
			id: "T1",
			value: *new(TestValue),
			expectError: false,
		},
	}
	
	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
    kg.AddTerm("T1")

	for index, testCase := range testCases {

		err := kg.SetTerm(testCase.id, testCase.value)
		if testCase.expectError {
			assertError(
				t, err,
				fmt.Sprintf(
					"Test Case #%d: Expected error, got %v", index + 1, err,
				),
			)
		} else {
			assertTrue(
				t, err == nil,
				fmt.Sprintf(
					"Test Case #%d: Unexpected error: %v", index + 1, err,
				),
			)
			assertTrue(
				t, kg.terms[testCase.id] == testCase.value,
				fmt.Sprintf(
					"Test Case #%d: Expected value: %v, found: %v",
					index + 1, testCase.value, kg.terms[testCase.id],
				),
			)
		}
		
	}
}

func TestSetTermInternallyMarksTermsAsSet(t *testing.T) {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
	ids := []string{"T1", "T2", "T3", "T4", "T5"}

	for _, id := range ids {
		kg.AddTerm(id)
	}
	
	for index, id := range ids {
		
		// Ensure all later ids are not set.
		for i := index; i < len(ids); i++ {
			assertTrue(
				t, !kg.setTerms[ids[i]],
				fmt.Sprintf(
					"Term '%s' was incorrectly marked set. Index: %d, i: %d", id, index, i,
				),
			)
		}
		// Ensure all prev ids are set
		for i := index - 1; i > 0; i-- {
			assertTrue(
				t, kg.setTerms[ids[i]],
				fmt.Sprintf(
					"Term '%s' was incorrectly not marked set. Index: %d, i: %d", id, index, i,
				),
			)
		}

		kg.SetTerm(id, NewTestValue(1, 2))
	}
}

func TestCopyCreatesIdenticalKnowledgeGraph(t *testing.T) {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
	
	kg.AddTerm("T1")
	kg.AddTerm("T2")	
	kg.AddTerm("T3")

	kg.AddOperation("O1", NewTestOperation("T1", "T2", "T3"))
	kg.AddOperation("O2", NewTestOperation("T2", "T1", "T3"))

	kg.SetTerm("T2", NewTestValue(1, 3)) 

	newKg := kg.Copy()

	assertTrue(
		t, reflect.DeepEqual(kg, newKg),
		fmt.Sprintf(
			"The copy should be deeply equal to the original. \n    Orig: %+v\n    Copy: %+v",
			kg, newKg,
		),	
	)
}


func TestCopyCreatesDistinctKnowledgeGraphs(t *testing.T) {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()
	
	kg.AddTerm("T1")
	kg.AddTerm("T2")	
	kg.AddTerm("T3")

	kg.AddOperation("O1", NewTestOperation("T1", "T2", "T3"))
	kg.AddOperation("O2", NewTestOperation("T2", "T1", "T3"))

	kg.SetTerm("T2", NewTestValue(1,5)) 

	t.Logf("The original knowledge graph is: %+v", kg)
	
	newKg := kg.Copy()
	
	newKg.AddTerm("T4")
	_, exists := kg.terms["T4"]
	assertTrue(
		t, !exists,
		fmt.Sprintf(
			"The original KG should not be modified by the copy's AddTerm(). \n    Orig: %+v\n    Copy: %+v",
			kg, newKg,
		),
	)

	newKg.SetTerm("T3", NewTestValue(1, 4))
	oldValue := kg.terms["T3"]
	assertEqualTestValues(
		t, oldValue, *new(TestValue),
		fmt.Sprintf(
			"The original KG should not be modified by the copy's SetTerm(). \n    Orig: %+v\n    Copy: %+v",
			kg, newKg,
		),
	)

	newKg.AddOperation("O3", MakeExampleTestOperation())
	_, exists = kg.operations["O3"]
	assertTrue(
		t, !exists,
		fmt.Sprintf(
			"The original KG should not be modified by the copy's AddOperation(). \n    Orig: %+v\n    Copy: %+v",
			kg, newKg,
		),
	)
}


func TestGetTermValue(t *testing.T) {

	testCases := []struct{id string; value TestValue; expectError bool}{
		{
			id: "T1", value: NewTestValue(5, 1.0),
		},
		{
			id: "T2", expectError: true,
		},
		{
			id: "T2", expectError: true,
		},
	}


	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

	kg.AddTerm("T1")
	kg.AddTerm("T2")

	kg.SetTerm("T1", NewTestValue(5, 1.0))
	
	for index, test := range testCases {

		value, err := kg.GetTerm(test.id)

		if test.expectError {
			assertError(
				t, err,
				fmt.Sprintf(
					"Test Case #%d: Expect error, got value %v", index + 1, value,
				),
			)
		} else {
			assertTrue(
				t, err == nil,
				fmt.Sprintf(
					"Test Case #%d: Expected no error, got error %v", index + 1, err,
				),
			)
			assertEqualTestValues(
				t, value, test.value,
				fmt.Sprintf(
					"Test Case #%d: Expected value %v, got %v", index + 1, test.value, value,
				),
			)
		}
	}
}

