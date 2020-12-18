package unit_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

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

func testTypes(t *testing.T, root tree.Node, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	t.Run("Top down", func(t *testing.T) { testImpl(t, root, tree.TopDown, expectedTopDown) })
	t.Run("Balanced", func(t *testing.T) { testImpl(t, root, tree.Balanced, expectedBalanced) })
	t.Run("Balanced favour top", func(t *testing.T) { testImpl(t, root, tree.BalancedFavourTop, expectedBalancedFavourTop) })
	t.Run("Bottom up", func(t *testing.T) { testImpl(t, root, tree.BottomUp, expectedBottomUp) })
}

func testImpl(t *testing.T, root tree.Node, displayType tree.DisplayType, expected []string) {
	// Arrange
	config := tree.DisplayTreeConfig{Type: displayType}

	// Act
	result := tree.DisplayTree(root, config)
	display(result)

	// Assert
	assertDisplay(t, result, expected)
}

func display(lines []string) {
	for _, ln := range lines {
		fmt.Println(ln)
	}
}

func assertDisplay(t *testing.T, actual []string, expected []string) {
	Assert(t).That(len(actual), is.EqualTo(len(expected)), "Correct number of lines")
	for i := range actual {
		Assert(t).That(actual[i], is.EqualTo(expected[i]), fmt.Sprintf("Line %d", i))
	}
}

type testTreeNode struct {
	value    string
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
