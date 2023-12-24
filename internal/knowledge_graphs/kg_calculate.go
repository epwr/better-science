package knowledge_graphs

import (
	"errors"
	"fmt"
	"slices"
)

// Tries to calculate the provided value.
func (kg *KnowledgeGraph[ID, Value, Operation]) Calculate(goalID ID) (Value, error) {

	// Get determenistic ordering of operations
	orderedOpIDs := []ID{}
	for key, _ := range kg.operations {
		orderedOpIDs = append(orderedOpIDs, key)
	}
	slices.Sort(orderedOpIDs)
	
	updateMadeInPass := true
	for updateMadeInPass {
		updateMadeInPass = false

		for _, opID := range orderedOpIDs {
			
			operation := kg.operations[opID]
			opCanBeRun := true
			requiredTermIDs := operation.Requires()
			requiredTerms := make(map[ID]Value)

			// Check if op can be run & build map of required Terms
			for i := 0; i < len(requiredTermIDs) && opCanBeRun; i++ {
				if !kg.setTerms[requiredTermIDs[i]] {
					opCanBeRun = false
				}
				requiredTerms[requiredTermIDs[i]] = kg.terms[requiredTermIDs[i]]
			}

			if opCanBeRun {

				fmt.Printf("Running Operation %v\n", opID)
				
				newValue, err := operation.Run(requiredTerms)
				if err == nil {
					termID := operation.Calculates()
					oldValue, exists := kg.terms[termID]
					if exists {
						if newValue.GetWeight() > oldValue.GetWeight() {
							kg.SetTerm(termID, newValue)
							updateMadeInPass = true
						}
					} else {
						kg.SetTerm(termID, newValue)
						updateMadeInPass = true
					}	
				}
			}
		}
	}

	if kg.setTerms[goalID] {
		return kg.terms[goalID], nil
	}
	return *new(Value), errors.New("Could not calculate the term")
}
