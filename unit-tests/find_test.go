package unit_tests

import (
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/tree"
)

func TestFindNotFoundSingleNode(t *testing.T) {
	// Arrange
	root := createNode("root")

	// Act
	here, _ := tree.FindValue(root, "not_there")

	// Assert
	Assert(t).That(here, is.False, "Not found")
}

func TestFindNotFoundManyNodes(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createNode("Child1", child11, child12)
	child21 := createNode("Child21")
	child2 := createNode("Child2", child21)
	root := createNode("root", child1, child2)

	// Act
	here, _ := tree.FindValue(root, "not_there")

	// Assert
	Assert(t).That(here, is.False, "Not found")
}

func TestFindFoundAtRoot(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createNode("Child1", child11, child12)
	child21 := createNode("Child21")
	child2 := createNode("Child2", child21)
	root := createNode("root", child1, child2)

	// Act
	here, nodes := tree.FindValue(root, "root")

	// Assert
	Assert(t).That(here, is.True, "Found")
	Assert(t).That(nodes, is.DeepEqualTo([]tree.Node{root}), "Root")
}

func TestFindFoundAtChild(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createNode("Child1", child11, child12)
	child21 := createNode("Child21")
	child2 := createNode("Child2", child21)
	root := createNode("root", child1, child2)

	// Act
	here, nodes := tree.FindValue(root, "Child1")

	// Assert
	Assert(t).That(here, is.True, "Found")
	Assert(t).That(nodes, is.DeepEqualTo([]tree.Node{root,child1}), "Root-Child1")
}

func TestFindFoundAtSubChild(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createNode("Child1", child11, child12)
	child21 := createNode("Child21")
	child2 := createNode("Child2", child21)
	root := createNode("root", child1, child2)

	// Act
	here, nodes := tree.FindValue(root, "Child12")

	// Assert
	Assert(t).That(here, is.True, "Found")
	Assert(t).That(nodes, is.DeepEqualTo([]tree.Node{root,child1,child12}), "Root-Child1-Child12")
}
