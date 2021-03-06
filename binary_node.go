package tree

// Binary tree node
type BinaryNode interface {
	// The value associated with this node
	NodeValue() interface{}
	// This node's children
	Left() BinaryNode
	Right() BinaryNode
	// Is this node the same node (or copy of)
	Equals(rhs BinaryNode) bool
}

// ---------------------------------------------------------------------
// Binary tree wrapper
// ---------------------------------------------------------------------
type BinaryNodeTreeWrapper struct {
	BinaryNode
}

func (n BinaryNodeTreeWrapper) NodeValue() interface{} {
	if n.BinaryNode != nil {
		return n.BinaryNode.NodeValue()
	}
	return nil
}

func (n BinaryNodeTreeWrapper) Children() []Node {
	if n.BinaryNode != nil && (n.BinaryNode.Left() != nil || n.BinaryNode.Right() != nil) {
		return []Node{BinaryNodeTreeWrapper{n.BinaryNode.Left()}, BinaryNodeTreeWrapper{n.BinaryNode.Right()}}
	}
	return make([]Node, 0)
}

func (n BinaryNodeTreeWrapper) Equals(rhs Node) bool {
	switch nt := rhs.(type) {
	case BinaryNodeTreeWrapper:
		if n.BinaryNode != nil {
			return n.BinaryNode.Equals(nt.BinaryNode)
		}
	}
	return false
}
