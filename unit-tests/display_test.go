package unit_tests

import (
	"testing"

	"github.com/atrico-go/tree"
)

func TestDisplaySingleNode(t *testing.T) {
	root := createNode("Root")
	expected := []string{
		"─ Root",
	}
	testTypes(t, root, expected, expected, expected, expected)
}

func TestDisplaySingleChild(t *testing.T) {
	child1 := createNode("Child1")
	root := createNode("Root", child1)
	expectedTopDown := []string{
		"─ Root",
		"  └─ Child1",
	}
	expectedBalanced := expectedTopDown
	expectedBalancedFavourTop := []string{
		"  ┌─ Child1",
		"─ Root",
	}
	expectedBottomUp := expectedBalancedFavourTop
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplaySingleLevelChildrenEven(t *testing.T) {
	child1 := createNode("Child1")
	child2 := createNode("Child2")
	root := createNode("Root", child1, child2)
	expectedTopDown := []string{
		"─ Root",
		"  ├─ Child1",
		"  └─ Child2",
	}
	expectedBalanced := []string{
		"  ┌─ Child1",
		"─ Root",
		"  └─ Child2",
	}
	expectedBalancedFavourTop := expectedBalanced
	expectedBottomUp := []string{
		"  ┌─ Child1",
		"  ├─ Child2",
		"─ Root",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplaySingleLevelChildrenOdd(t *testing.T) {
	child1 := createNode("Child1")
	child2 := createNode("Child2")
	child3 := createNode("Child3")
	root := createNode("Root", child1, child2, child3)
	expectedTopDown := []string{
		"─ Root",
		"  ├─ Child1",
		"  ├─ Child2",
		"  └─ Child3",
	}
	expectedBalanced := []string{
		"  ┌─ Child1",
		"─ Root",
		"  ├─ Child2",
		"  └─ Child3",
	}
	expectedBalancedFavourTop := []string{
		"  ┌─ Child1",
		"  ├─ Child2",
		"─ Root",
		"  └─ Child3",
	}
	expectedBottomUp := []string{
		"  ┌─ Child1",
		"  ├─ Child2",
		"  ├─ Child3",
		"─ Root",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplayMultipleLevelChildren(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createNode("Child1", child11, child12)
	child21 := createNode("Child21")
	child22 := createNode("Child22")
	child2 := createNode("Child2", child21, child22)
	root := createNode("Root", child1, child2)
	expectedTopDown := []string{
		"─ Root",
		"  ├─ Child1",
		"  │  ├─ Child11",
		"  │  └─ Child12",
		"  └─ Child2",
		"     ├─ Child21",
		"     └─ Child22",
	}
	expectedBalanced := []string{
		"     ┌─ Child11",
		"  ┌─ Child1",
		"  │  └─ Child12",
		"─ Root",
		"  │  ┌─ Child21",
		"  └─ Child2",
		"     └─ Child22",
	}
	expectedBalancedFavourTop := expectedBalanced
	expectedBottomUp := []string{
		"     ┌─ Child11",
		"     ├─ Child12",
		"  ┌─ Child1",
		"  │  ┌─ Child21",
		"  │  ├─ Child22",
		"  ├─ Child2",
		"─ Root",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplayManyLevelChildren(t *testing.T) {
	// Arrange
	child111 := createNode("Child111")
	child11 := createNode("Child11", child111)
	child121 := createNode("Child121")
	child12 := createNode("Child12", child121)
	child1 := createNode("Child1", child11, child12)
	child211 := createNode("Child211")
	child21 := createNode("Child21", child211)
	child221 := createNode("Child221")
	child22 := createNode("Child22", child221)
	child2 := createNode("Child2", child21, child22)
	root := createNode("Root", child1, child2)
	expectedTopDown := []string{
		"─ Root",
		"  ├─ Child1",
		"  │  ├─ Child11",
		"  │  │  └─ Child111",
		"  │  └─ Child12",
		"  │     └─ Child121",
		"  └─ Child2",
		"     ├─ Child21",
		"     │  └─ Child211",
		"     └─ Child22",
		"        └─ Child221",
	}
	expectedBalanced := []string{
		"     ┌─ Child11",
		"     │  └─ Child111",
		"  ┌─ Child1",
		"  │  └─ Child12",
		"  │     └─ Child121",
		"─ Root",
		"  │  ┌─ Child21",
		"  │  │  └─ Child211",
		"  └─ Child2",
		"     └─ Child22",
		"        └─ Child221",
	}
	expectedBalancedFavourTop := []string{
		"        ┌─ Child111",
		"     ┌─ Child11",
		"  ┌─ Child1",
		"  │  │  ┌─ Child121",
		"  │  └─ Child12",
		"─ Root",
		"  │     ┌─ Child211",
		"  │  ┌─ Child21",
		"  └─ Child2",
		"     │  ┌─ Child221",
		"     └─ Child22",
	}
	expectedBottomUp := []string{
		"        ┌─ Child111",
		"     ┌─ Child11",
		"     │  ┌─ Child121",
		"     ├─ Child12",
		"  ┌─ Child1",
		"  │     ┌─ Child211",
		"  │  ┌─ Child21",
		"  │  │  ┌─ Child221",
		"  │  ├─ Child22",
		"  ├─ Child2",
		"─ Root",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplayEmptyNodesSingleNode(t *testing.T) {
	root := createEmptyNode()
	expected := []string{
		"─ ",
	}
	testTypes(t, root, expected, expected, expected, expected)
}

func TestDisplayEmptyNodesManyLevelChildren(t *testing.T) {
	// Arrange
	child111 := createNode("Child111")
	child11 := createNode("Child11", child111)
	child121 := createNode("Child121")
	child12 := createEmptyNode(child121)
	child1 := createNode("Child1", child11, child12)
	child211 := createNode("Child211")
	child21 := createNode("Child21", child211)
	child221 := createNode("Child221")
	child22 := createNode("Child22", child221)
	child2 := createNode("Child2", child21, child22)
	root := createEmptyNode(child1, child2)
	expectedTopDown := []string{
		"─ ┐",
		"  ├─ Child1",
		"  │  ├─ Child11",
		"  │  │  └─ Child111",
		"  │  └─ ┐",
		"  │     └─ Child121",
		"  └─ Child2",
		"     ├─ Child21",
		"     │  └─ Child211",
		"     └─ Child22",
		"        └─ Child221",
	}
	expectedBalanced := []string{
		"     ┌─ Child11",
		"     │  └─ Child111",
		"  ┌─ Child1",
		"  │  └─ ┐",
		"  │     └─ Child121",
		"─ ┤",
		"  │  ┌─ Child21",
		"  │  │  └─ Child211",
		"  └─ Child2",
		"     └─ Child22",
		"        └─ Child221",
	}
	expectedBalancedFavourTop := []string{
		"        ┌─ Child111",
		"     ┌─ Child11",
		"  ┌─ Child1",
		"  │  │  ┌─ Child121",
		"  │  └─ ┘",
		"─ ┤",
		"  │     ┌─ Child211",
		"  │  ┌─ Child21",
		"  └─ Child2",
		"     │  ┌─ Child221",
		"     └─ Child22",
	}
	expectedBottomUp := []string{
		"        ┌─ Child111",
		"     ┌─ Child11",
		"     │  ┌─ Child121",
		"     ├─ ┘",
		"  ┌─ Child1",
		"  │     ┌─ Child211",
		"  │  ┌─ Child21",
		"  │  │  ┌─ Child221",
		"  │  ├─ Child22",
		"  ├─ Child2",
		"─ ┘",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}
func TestDisplayEmptyNodesAllEmpty(t *testing.T) {
	// Arrange
	child11 := createNode("Child11")
	child12 := createNode("Child12")
	child1 := createEmptyNode(child11, child12)
	child21 := createNode("Child21")
	child22 := createNode("Child22")
	child2 := createEmptyNode(child21, child22)
	root := createEmptyNode(child1, child2)
	expectedTopDown := []string{
		"─ ┐",
		"  ├─ ┐",
		"  │  ├─ Child11",
		"  │  └─ Child12",
		"  └─ ┐",
		"     ├─ Child21",
		"     └─ Child22",
	}
	expectedBalanced := []string{
		"     ┌─ Child11",
		"  ┌─ ┤",
		"  │  └─ Child12",
		"─ ┤",
		"  │  ┌─ Child21",
		"  └─ ┤",
		"     └─ Child22",
	}
	expectedBalancedFavourTop := expectedBalanced
	expectedBottomUp := []string{
		"     ┌─ Child11",
		"     ├─ Child12",
		"  ┌─ ┘",
		"  │  ┌─ Child21",
		"  │  ├─ Child22",
		"  ├─ ┘",
		"─ ┘",
	}
	testTypes(t, root, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func testTypes(t *testing.T, root tree.Node, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	displayFunc := func(config tree.DisplayTreeConfig) []string {
		return tree.DisplayTree(root, config)
	}
	testTypesImpl(t, displayFunc, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

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
