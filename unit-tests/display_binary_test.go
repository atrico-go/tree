package unit_tests

import (
	"testing"
)

func TestDisplayBinarySingleNode(t *testing.T) {
	root := createBinaryLeafNode("Root")
	expected := []string{
		"── Root",
	}
	testDisplayTypesBinary(t, root, expected, expected, expected)
}

func TestDisplayBinarySingleChildLeft(t *testing.T) {
	childL := createBinaryLeafNode("ChildL")
	root := createBinaryNode("Root", childL, nil)
	expectedTopDown := []string{
		"─┬ Root",
		" ├── ChildL",
		" └──",
	}
	expectedBalanced := []string{
		" ┌── ChildL",
		"─┼ Root",
		" └──",
	}
	expectedBottomUp := []string{
		" ┌── ChildL",
		" ├──",
		"─┴ Root",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}
func TestDisplayBinarySingleChildRight(t *testing.T) {
	childR := createBinaryLeafNode("ChildR")
	root := createBinaryNode("Root", nil, childR)
	expectedTopDown := []string{
		"─┬ Root",
		" ├──",
		" └── ChildR",
	}
	expectedBalanced := []string{
		" ┌──",
		"─┼ Root",
		" └── ChildR",
	}
	expectedBottomUp := []string{
		" ┌──",
		" ├── ChildR",
		"─┴ Root",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}

func TestDisplayBinarySingleLevelChildrenEven(t *testing.T) {
	childL := createBinaryLeafNode("ChildL")
	childR := createBinaryLeafNode("ChildR")
	root := createBinaryNode("Root", childL, childR)
	expectedTopDown := []string{
		"─┬ Root",
		" ├── ChildL",
		" └── ChildR",
	}
	expectedBalanced := []string{
		" ┌── ChildL",
		"─┼ Root",
		" └── ChildR",
	}
	expectedBottomUp := []string{
		" ┌── ChildL",
		" ├── ChildR",
		"─┴ Root",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}

func TestDisplayBinaryMultipleLevelChildren(t *testing.T) {
	// Arrange
	childLL := createBinaryLeafNode("ChildLL")
	childLR := createBinaryLeafNode("ChildLR")
	childL := createBinaryNode("ChildL", childLL, childLR)
	childRL := createBinaryLeafNode("ChildRL")
	childRR := createBinaryLeafNode("ChildRR")
	childR := createBinaryNode("ChildR", childRL, childRR)
	root := createBinaryNode("Root", childL, childR)
	expectedTopDown := []string{
		"─┬ Root",
		" ├─┬ ChildL",
		" │ ├── ChildLL",
		" │ └── ChildLR",
		" └─┬ ChildR",
		"   ├── ChildRL",
		"   └── ChildRR",
	}
	expectedBalanced := []string{
		"   ┌── ChildLL",
		" ┌─┼ ChildL",
		" │ └── ChildLR",
		"─┼ Root",
		" │ ┌── ChildRL",
		" └─┼ ChildR",
		"   └── ChildRR",
	}
	expectedBottomUp := []string{
		"   ┌── ChildLL",
		"   ├── ChildLR",
		" ┌─┴ ChildL",
		" │ ┌── ChildRL",
		" │ ├── ChildRR",
		" ├─┴ ChildR",
		"─┴ Root",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}

func TestDisplayBinaryManyLevelChildren(t *testing.T) {
	// Arrange
	childLLL := createBinaryLeafNode("ChildLLL")
	childLL := createBinaryNode("ChildLL", childLLL, nil)
	childLRR := createBinaryLeafNode("ChildLRR")
	childLR := createBinaryNode("ChildLR", nil, childLRR)
	childL := createBinaryNode("ChildL", childLL, childLR)
	childRLR := createBinaryLeafNode("ChildRLR")
	childRL := createBinaryNode("ChildRL", nil, childRLR)
	childRRL := createBinaryLeafNode("ChildRRL")
	childRR := createBinaryNode("ChildRR", childRRL, nil)
	childR := createBinaryNode("ChildR", childRL, childRR)
	root := createBinaryNode("Root", childL, childR)
	expectedTopDown := []string{
		"─┬ Root",
		" ├─┬ ChildL",
		" │ ├─┬ ChildLL",
		" │ │ ├── ChildLLL",
		" │ │ └──",
		" │ └─┬ ChildLR",
		" │   ├──",
		" │   └── ChildLRR",
		" └─┬ ChildR",
		"   ├─┬ ChildRL",
		"   │ ├──",
		"   │ └── ChildRLR",
		"   └─┬ ChildRR",
		"     ├── ChildRRL",
		"     └──",
	}
	expectedBalanced := []string{
		"     ┌── ChildLLL",
		"   ┌─┼ ChildLL",
		"   │ └──",
		" ┌─┼ ChildL",
		" │ │ ┌──",
		" │ └─┼ ChildLR",
		" │   └── ChildLRR",
		"─┼ Root",
		" │   ┌──",
		" │ ┌─┼ ChildRL",
		" │ │ └── ChildRLR",
		" └─┼ ChildR",
		"   │ ┌── ChildRRL",
		"   └─┼ ChildRR",
		"     └──",
	}
	expectedBottomUp := []string{
		"     ┌── ChildLLL",
		"     ├──",
		"   ┌─┴ ChildLL",
		"   │ ┌──",
		"   │ ├── ChildLRR",
		"   ├─┴ ChildLR",
		" ┌─┴ ChildL",
		" │   ┌──",
		" │   ├── ChildRLR",
		" │ ┌─┴ ChildRL",
		" │ │ ┌── ChildRRL",
		" │ │ ├──",
		" │ ├─┴ ChildRR",
		" ├─┴ ChildR",
		"─┴ Root",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}

func TestDisplayBinaryEmptyNodesSingleNode(t *testing.T) {
	root := createBinaryEmptyNode(nil, nil)
	expected := []string{
		"──",
	}
	testDisplayTypesBinary(t, root, expected, expected, expected)
}

func TestDisplayBinaryEmptyNodesManyLevelChildren(t *testing.T) {
	// Arrange
	childLLL := createBinaryLeafNode("ChildLLL")
	childLL := createBinaryNode("ChildLL", childLLL, nil)
	childLRR := createBinaryLeafNode("ChildLRR")
	childLR := createBinaryEmptyNode(nil, childLRR)
	childL := createBinaryNode("ChildL", childLL, childLR)
	childRLR := createBinaryLeafNode("ChildRLR")
	childRL := createBinaryNode("ChildRL", nil, childRLR)
	childRRL := createBinaryLeafNode("ChildRRL")
	childRR := createBinaryNode("ChildRR", childRRL, nil)
	childR := createBinaryNode("ChildR", childRL, childRR)
	root := createBinaryEmptyNode(childL, childR)
	expectedTopDown := []string{
		"─┐",
		" ├─┬ ChildL",
		" │ ├─┬ ChildLL",
		" │ │ ├── ChildLLL",
		" │ │ └──",
		" │ └─┐",
		" │   ├──",
		" │   └── ChildLRR",
		" └─┬ ChildR",
		"   ├─┬ ChildRL",
		"   │ ├──",
		"   │ └── ChildRLR",
		"   └─┬ ChildRR",
		"     ├── ChildRRL",
		"     └──",
	}
	expectedBalanced := []string{
		"     ┌── ChildLLL",
		"   ┌─┼ ChildLL",
		"   │ └──",
		" ┌─┼ ChildL",
		" │ │ ┌──",
		" │ └─┤",
		" │   └── ChildLRR",
		"─┤",
		" │   ┌──",
		" │ ┌─┼ ChildRL",
		" │ │ └── ChildRLR",
		" └─┼ ChildR",
		"   │ ┌── ChildRRL",
		"   └─┼ ChildRR",
		"     └──",
	}
	expectedBottomUp := []string{
		"     ┌── ChildLLL",
		"     ├──",
		"   ┌─┴ ChildLL",
		"   │ ┌──",
		"   │ ├── ChildLRR",
		"   ├─┘",
		" ┌─┴ ChildL",
		" │   ┌──",
		" │   ├── ChildRLR",
		" │ ┌─┴ ChildRL",
		" │ │ ┌── ChildRRL",
		" │ │ ├──",
		" │ ├─┴ ChildRR",
		" ├─┴ ChildR",
		"─┘",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}

func TestDisplayBinaryEmptyNodesAllEmpty(t *testing.T) {
	// Arrange
	childLL := createBinaryLeafNode("ChildLL")
	childLR := createBinaryLeafNode("ChildLR")
	childL := createBinaryEmptyNode(childLL, childLR)
	childRL := createBinaryLeafNode("ChildRL")
	childRR := createBinaryLeafNode("ChildRR")
	childR := createBinaryEmptyNode(childRL, childRR)
	root := createBinaryEmptyNode(childL, childR)
	expectedTopDown := []string{
		"─┐",
		" ├─┐",
		" │ ├── ChildLL",
		" │ └── ChildLR",
		" └─┐",
		"   ├── ChildRL",
		"   └── ChildRR",
	}
	expectedBalanced := []string{
		"   ┌── ChildLL",
		" ┌─┤",
		" │ └── ChildLR",
		"─┤",
		" │ ┌── ChildRL",
		" └─┤",
		"   └── ChildRR",
	}
	expectedBottomUp := []string{
		"   ┌── ChildLL",
		"   ├── ChildLR",
		" ┌─┘",
		" │ ┌── ChildRL",
		" │ ├── ChildRR",
		" ├─┘",
		"─┘",
	}
	testDisplayTypesBinary(t, root, expectedTopDown, expectedBalanced, expectedBottomUp)
}
