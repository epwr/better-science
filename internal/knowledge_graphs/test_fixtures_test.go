package knowledge_graphs

/*
 *     Mock Operation
 */

// TestValue
type TestValue struct {
	value int
	weight float64
}

func (val TestValue) GetWeight() float64 {
	return val.weight
}

func NewTestValue(value int, weight float64) TestValue {
	return TestValue{value: value, weight: weight}
}



// Operations must provide Calculates and Requires methods. This is an example type that
// is used throughout testing
type TestOperation struct {
	calc string
	reqs []string
	return_value TestValue
	return_error error
	return_weight int
}

func (op TestOperation) Calculates() string {
	return op.calc
}

func (op TestOperation) Requires() []string {
	return op.reqs
}

func (op TestOperation) Run(terms map[string]TestValue) (TestValue, error) {
	return op.return_value, op.return_error
}

func MakeExampleTestOperation() TestOperation {
	return TestOperation{calc: "T1", reqs: []string{"T2"}}
}

func NewTestOperation(calc string, reqs ...string) TestOperation {
	return TestOperation{calc: calc, reqs: reqs}
}


/*
 *     Knowledge Graph Fixtures
*/

func fixtureSimpleKG() KnowledgeGraph[string, TestValue, TestOperation] {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

	kg.AddTerm("T1")
	kg.AddTerm("T2")
	kg.AddTerm("T3")
	kg.AddTerm("T4")
	kg.AddTerm("T5")

	kg.AddOperation("O1", TestOperation{
		calc: "T1",
		reqs: []string{"T2"},
		return_value: NewTestValue(2, 0.2),
	})

	kg.AddOperation("O2", TestOperation{
		calc: "T1",
		reqs: []string{"T3"},
		return_value: NewTestValue(3, 0.3),
	})

	kg.AddOperation("O3", TestOperation{
		calc: "T3",
		reqs: []string{"T4", "T5"},
		return_value: NewTestValue(9, 0.3),
	})

	return kg
}


func fixtureDisjointKG() KnowledgeGraph[string, TestValue, TestOperation] {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

	kg.AddTerm("T1")
	kg.AddTerm("T2")
	kg.AddTerm("T3")
	kg.AddTerm("T4")
	kg.AddTerm("T5")

	kg.AddOperation("O1", TestOperation{
		calc: "T1",
		reqs: []string{"T2"},
		return_value: NewTestValue(2, 0.2),
	})

	kg.AddOperation("O2", TestOperation{
		calc: "T2",
		reqs: []string{"T1"},
		return_value: NewTestValue(1, 0.2),
	})

	kg.AddOperation("O3", TestOperation{
		calc: "T3",
		reqs: []string{"T4", "T5"},
		return_value: NewTestValue(9, 0.2),
	})

	return kg
	
}


func fixtureCyclicalKG() KnowledgeGraph[string, TestValue, TestOperation] {

	kg := NewKnowledgeGraph[string, TestValue, TestOperation]()

	kg.AddTerm("T1")
	kg.AddTerm("T2")
	kg.AddTerm("T3")
	kg.AddTerm("T4")
	kg.AddTerm("T5")

	kg.AddOperation("O1", TestOperation{
		calc: "T1",
		reqs: []string{"T2"},
		return_value: NewTestValue(2, 0.2),
	})

	kg.AddOperation("O2", TestOperation{
		calc: "T2",
		reqs: []string{"T3"},
		return_value: NewTestValue(3, 0.2),
	})

	kg.AddOperation("O3", TestOperation{
		calc: "T3",
		reqs: []string{"T4"},
		return_value: NewTestValue(4, 0.2),
	})

	kg.AddOperation("O4", TestOperation{
		calc: "T4",
		reqs: []string{"T5"},
		return_value: NewTestValue(5, 0.2),
	})
	
	kg.AddOperation("O5", TestOperation{
		calc: "T5",
		reqs: []string{"T1"},
		return_value: NewTestValue(1, 0.2),
	})

	return kg
	
}

