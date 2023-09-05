// Copywrite 2023 Eric Power - All Rights Reserved
package kg_engine

/* NOTE
 * I'm not dealing with tracking the precision of values, data lineage, or units. I would like
 * to do this. I suspect once I get a basic KG Engine working, I will need to set up a
 * data_representation layer with a Value that tracks it's precision, units, and lineage. 
 * 
 * This data_representation layer should also provide all the basic Operations (add,
 * subtract, multiply, divide, exponents, equality, ordering, min, max, mean, etc) for Numeric and
 * Enumeration values.
 *
 * Data lineage probably isn't needed early on. Nor are units. Maybe track sigfigs only?
 */

// Node represents a node in a Graph
type Node interface {
	GetType() string
	GetName() string
}

// OperationNode represents an operation node in the graph.
type OperationNode struct {
	Name string
	Confidence float64
	Inputs []*ValueNode
	Result ValueNode
}

// GetType returns the type of the node.
func (n OperationNode) GetType() string {
	return "OperationNode"
}

// GetName returns the type of the node.
func (n OperationNode) GetName() string {
	return n.Name
}

// ValueNode represents a value node in the graph that stores an int
type ValueNode struct {
	Name string
	Value int
	Confidence float64
	CalculatedBy []*OperationNode
	UsedBy []*OperationNode
}

// GetType returns the type of the node.
func (n ValueNode) GetType() string {
	return "ValueNode"
}

// GetName returns the type of the node.
func (n ValueNode) GetName() string {
	return n.Name
}




