package unit_tests

import (
	"github.com/segmentio/ksuid"

	"github.com/atrico-go/tree"
)

// -------------------------------------------------------------------------------------------------
// Tree
// -------------------------------------------------------------------------------------------------

type testTreeNode struct {
	uid      ksuid.KSUID
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
	return testTreeNode{ksuid.New(), value, children}
}

func createEmptyNode(children ...tree.Node) tree.Node {
	return testTreeNode{ksuid.New(), nil, children}
}

func (n testTreeNode) Equals(rhs tree.Node) bool {
	switch nt := rhs.(type) {
	case testTreeNode:
		return n.uid == nt.uid
	}
	return false
}

// -------------------------------------------------------------------------------------------------
// Binary tree
// -------------------------------------------------------------------------------------------------
type testBinaryTreeNode struct {
	uid      ksuid.KSUID
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
func (n testBinaryTreeNode) Equals(rhs tree.BinaryNode) bool {
	switch nt := rhs.(type) {
	case testBinaryTreeNode:
		return n.uid == nt.uid
	}
	return false
}
func createBinaryNode(value string, left tree.BinaryNode, right tree.BinaryNode) tree.BinaryNode {
	return testBinaryTreeNode{ksuid.New(),value, left, right}
}
func createBinaryEmptyNode(left tree.BinaryNode, right tree.BinaryNode) tree.BinaryNode {
	return testBinaryTreeNode{ksuid.New(),nil, left, right}
}
func createBinaryLeafNode(value string) tree.BinaryNode {
	return createBinaryNode(value, nil, nil)
}
