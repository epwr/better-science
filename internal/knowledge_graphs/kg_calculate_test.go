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


// TODO: test that the Operation.Run() gets the required arguments.
func TestOperationDotRunIsGivenCorrectArguments(t *testing.T) {

	type MockOperation struct{
		value TestValue;
		weight float64;
		calcs string;
		reqs string;
		tester *testing.T;
	}

	func (op MockOperation) 
	
	

	kg := NewKnowledgeGraph[string, TestValue, MockOperation]{}
	
	kg.AddTerm("T1")
	kg.AddTerm("T2")
	kg.AddTerm("T3")

	kg

}


