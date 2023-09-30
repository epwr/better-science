// Copywrite 2023 Eric Power - All Rights Reserved
package knowledge_graph

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
	GetNodeType() string	
	GetName() string
}

// OperationNode represents an operation node in the graph.
type OperationNode struct {
	Name string
	Inputs []*ValueNode
	Result *ValueNode
}

// GetNodeType returns the type of the node.
func (n OperationNode) GetNodeType() string {
	return "OperationNode"
}

// GetName returns the type of the node.
func (n OperationNode) GetName() string {
	return n.Name
}

// ValueNode is the description of a potential value.
type value_type int
const(
	IntType value_type = iota
	FloatType
	EnumType
)

type ValueNode struct {
	Name string
	Type value_type
	CalculatedBy []*OperationNode
	UsedBy []*OperationNode
}

// GetNodeType returns the type of the node.
func (n ValueNode) GetNodeType() string {
	return "ValueNode"
}

// GetName returns the type of the node.
func (n ValueNode) GetName() string {
	return n.Name
}

