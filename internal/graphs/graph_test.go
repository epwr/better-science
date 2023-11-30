// Copywrite 2023 Eric Power - All Rights Reserved
package graphs_test

import (
	"testing"
)

type TestPayload {
    data string
}


func TestCreatingAGraph(t *testing.T) {

    graph := graphs.NewGraph[TestPayload]()
    
}
