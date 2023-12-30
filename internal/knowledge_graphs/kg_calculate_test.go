package knowledge_graphs

import (
	"fmt"
	"testing"
)

func TestKGCalculateSuccess(t *testing.T) {

	testCases := []struct{
		kg KnowledgeGraph[string, TestValue, TestOperation];
		termID string;
		expectedValue TestValue;
		initValue []struct{id string; value TestValue};
	}{
		{ // Trivial Case
			kg: fixtureSimpleKG(),
			termID: "T4",
			expectedValue: NewTestValue(4, 0.4),
			initValue: []struct{id string; value TestValue}{
				{id: "T4", value: NewTestValue(4, 0.4)},
			},
		},
		{ // One operation is calculated
			kg: fixtureSimpleKG(),
			termID: "T3",
			expectedValue: NewTestValue(9, 0.3),
			initValue: []struct{id string; value TestValue}{
				{id: "T4", value: NewTestValue(4, 0.4)},
				{id: "T5", value: NewTestValue(5, 0.4)},
			},
		},
		{ // A cyclical graph with both start and end nodes are on the cycle.
			kg: fixtureCyclicalKG(),
			termID: "T1",
			expectedValue: NewTestValue(2, 0.2),
			initValue: []struct{id string; value TestValue}{
				{id: "T2", value: NewTestValue(5, 0.4)},
			},
		},
		{ // Conflicting operations - select value with > weight (overwrite prev value)
			kg: fixtureSimpleKG(),
			termID: "T1",
			expectedValue: NewTestValue(3, 0.3),
			initValue: []struct{id string; value TestValue}{
				{id: "T2", value: NewTestValue(2, 0.4)},
				{id: "T4", value: NewTestValue(4, 0.4)},
				{id: "T5", value: NewTestValue(5, 0.4)},
			},
		},
		{ // Conflicting operations - select value with > weight (don't overwrite prev value)
			kg: fixtureSimpleKG(),
			termID: "T1",
			expectedValue: NewTestValue(2, 0.9),
			initValue: []struct{id string; value TestValue}{
				{id: "T1", value: NewTestValue(2, 0.9)},
				{id: "T2", value: NewTestValue(2, 0.4)},
				{id: "T4", value: NewTestValue(4, 0.4)},
				{id: "T5", value: NewTestValue(5, 0.4)},
			},
		},
		{ // Conflicting operations - select value with > weight (don't overwrite prev value)
			kg: fixtureSimpleKG(),
			termID: "T1",
			expectedValue: NewTestValue(2, 0.9),
			initValue: []struct{id string; value TestValue}{
				{id: "T1", value: NewTestValue(2, 0.9)},
				{id: "T2", value: NewTestValue(2, 0.4)},
				{id: "T4", value: NewTestValue(4, 0.4)},
				{id: "T5", value: NewTestValue(5, 0.4)},
			},
		},
	}

	for index, testCase := range testCases {

		for _, term := range testCase.initValue {
			testCase.kg.SetTerm(term.id, term.value)
		}
		
		value, err := testCase.kg.Calculate(testCase.termID)

		t.Logf("Test Case #%d: Kg Terms: %+v", index + 1, testCase.kg.terms)
		
		assertTrue(t, err == nil, "Calculate should not have returned an error.")
		assertTrue(
			t, value == testCase.expectedValue,
			fmt.Sprintf(
				"Test Case #%d: Expected value %+v, got %+v.", index + 1, testCase.expectedValue, value,
			),
		)
	}
}

// TODO: error paths -> no way to calculate the term do to disjoint, lack of base values, test with
// a cycle to know it halts.
func TestKGCalculateFails(t *testing.T) {

	testCases := []struct{
		kg KnowledgeGraph[string, TestValue, TestOperation];
		termID string;
		initValue []struct{id string; value TestValue};
	}{
		{ // No values in the KG
			kg: fixtureSimpleKG(),
			termID: "T1",
			initValue: []struct{id string; value TestValue}{},
		},
		{ // Disjoint KG with values on in the other side
			kg: fixtureDisjointKG(),
			termID: "T1",
			initValue: []struct{id string; value TestValue}{
				{id: "T4", value: NewTestValue(4, 0.4)},
				{id: "T5", value: NewTestValue(5, 0.4)},
			},
		},
		{ // No operations calculate the term.
			kg: fixtureSimpleKG(),
			termID: "T4",
			initValue: []struct{id string; value TestValue}{
				{id: "T2", value: NewTestValue(2, 0.4)},
				{id: "T3", value: NewTestValue(3, 0.4)},
			},
		},
		{ // Cycle with no way to calculate the term.
			kg: fixtureCyclicalKG(),
			termID: "T1",
			initValue: []struct{id string; value TestValue}{},
		},
		{ // Term ID does not exist
			kg: fixtureSimpleKG(),
			termID: "T17",
			initValue: []struct{id string; value TestValue}{
				{id: "T2", value: NewTestValue(2, 0.4)},
				{id: "T3", value: NewTestValue(3, 0.4)},
			},
		},
	}

	for index, testCase := range testCases {

		for _, term := range testCase.initValue {
			testCase.kg.SetTerm(term.id, term.value)
		}

		fmt.Printf("Text Case #%d\n", index + 1)
		value, err := testCase.kg.Calculate(testCase.termID)
		t.Logf("Test Case #%d: Kg Terms: %+v", index + 1, testCase.kg.terms)
		msg := fmt.Sprintf(
			"Calculate should have returned an error, instead found value: %+v", value,
		)
		assertError(t, err, msg)
	}
}


// Test that the Operation.Run() gets the required arguments.
type MockOperation struct{
	calcs string;
	reqs []string;
	t *testing.T;
	tests map[string]TestValue;
}

func (op MockOperation) Calculates() string {
	return op.calcs
}

func (op MockOperation) Requires() []string {
	return op.reqs
}

func (op MockOperation) Run(inputs map[string]TestValue) (TestValue, error) {

	passedInputKeys := make([]string, 0)
	for key, _ := range inputs {
		passedInputKeys = append(passedInputKeys, key)
	}
	assertSlicesAreSetwiseEqual(
		op.t, op.reqs, passedInputKeys,
		"Run should be passed only and all values for required terms.",
	)

	for key, expectedValue := range op.tests {
		assertEqualTestValues(
			op.t, inputs[key], expectedValue,
			fmt.Sprintf("Unexpected values were passed to Operation.Run() for term %v.", key),
		)
	}

	return NewTestValue(4, 0.2), nil
}

func TestOperationDotRunIsGivenCorrectArguments(t *testing.T) {

	kg := NewKnowledgeGraph[string, TestValue, MockOperation]()
	
	kg.AddTerm("T1")
	kg.AddTerm("T2")
	kg.AddTerm("T3")
	kg.AddTerm("T4")
	kg.AddTerm("T5")

	kg.AddOperation("O1", MockOperation{
		calcs: "T1",
		reqs: []string{"T2", "T4"},
		t: t,
		tests: map[string]TestValue{
			"T2": NewTestValue(2, 0.2),
			"T4": NewTestValue(4, 0.4),
		},
	})

	kg.SetTerm("T2", NewTestValue(2, 0.2))	
	kg.SetTerm("T3", NewTestValue(3, 0.3))
	kg.SetTerm("T4", NewTestValue(4, 0.4))

	value, err := kg.Calculate("T1")
	assertTrue(t, err == nil, "Test should have been able to calculate T1.")
	assertEqualTestValues(
		t, value, NewTestValue(4, 0.2),
		".Calculate returned a value that did not match what MockOperation.Run should return.",
	)
}


