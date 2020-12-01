package tree

// Basic tree node
type Node interface {
	// The value associated with this node
	NodeValue() interface{}
	// This node's children
	Children() []Node
}

// Binary tree node
type BinaryNode interface {
	// The value associated with this node
	NodeValue() interface{}
	// This node's children
	Left() BinaryNode
	Right() BinaryNode
}
