// Copywrite 2023 Eric Power - All Rights Reserved
//
// This package provides an implementation of a knowledge graph that is design to determine how to
// calculate a desired value.
//
// Conceptually these KnowledgeGraphs are a directed, bipartiate graph where the nodes are split into
// Terms and Operations. Terms only connect to Operations and vice-versa. Terms represent a node that
// can hold a Value, and an Operation represents a way to calculate a Value for a Term.
//
// WARNING:
// There is one important noteworthy odd behaviour in this package. Currently, if two operations
// produce different values with the same weight, via GetWeight(), then the first operation to be
// run will set the value for the term. If one Operation can be run in an earlier pass than the
// other, the one that can be run first will be used. If they are run in the same pass, then the
// Operations will be run based on their orderding, which is deterministic and based on their IDs.
package knowledge_graphs
