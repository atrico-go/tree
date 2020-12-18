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

func Contains(list []Node, node Node) (found bool,idx int) {
	for i,n := range list {
		if n.Equals(node) {
			return true,i
		}
	}
	return false,-1
}