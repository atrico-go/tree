package tree

// Basic tree node
type Node interface {
	// The value associated with this node
	NodeValue() interface{}
	// This node's children
	Children() []Node
	// Is this node the same node (or copy of)
	Equals(rhs Node) bool
}
