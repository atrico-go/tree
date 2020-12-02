package unit_tests

import "github.com/atrico-go/tree"

// -------------------------------------------------------------------------------------------------
// Tree
// -------------------------------------------------------------------------------------------------

type testTreeNode struct {
	value    interface{}
	children []tree.Node
}

func (n testTreeNode) NodeValue() interface{} {
	return n.value
}

func (n testTreeNode) Children() []tree.Node {
	return n.children
}

func createNode(value string, children ...tree.Node) tree.Node {
	return testTreeNode{value, children}
}

func createEmptyNode(children ...tree.Node) tree.Node {
	return testTreeNode{nil, children}
}

// -------------------------------------------------------------------------------------------------
// Binary tree
// -------------------------------------------------------------------------------------------------
type testBinaryTreeNode struct {
	value interface{}
	left  tree.BinaryNode
	right tree.BinaryNode
}

func (n testBinaryTreeNode) NodeValue() interface{} {
	return n.value
}
func (n testBinaryTreeNode) Left() tree.BinaryNode {
	return n.left
}
func (n testBinaryTreeNode) Right() tree.BinaryNode {
	return n.right
}
func createBinaryNode(value string, left tree.BinaryNode, right tree.BinaryNode) tree.BinaryNode {
	return testBinaryTreeNode{value, left, right}
}
func createBinaryEmptyNode(left tree.BinaryNode, right tree.BinaryNode) tree.BinaryNode {
	return testBinaryTreeNode{nil, left, right}
}
func createBinaryLeafNode(value string) tree.BinaryNode {
	return createBinaryNode(value, nil, nil)
}

