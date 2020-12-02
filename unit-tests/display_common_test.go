package unit_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/tree"
)

type displayFuncDef func(config tree.DisplayTreeConfig) []string
type testConfig struct {
	displayFunc   displayFuncDef
	displayConfig tree.DisplayTreeConfig
}
func NewTestConfig() testConfig {
	return testConfig{nil, tree.NewDisplayConfig()}
}
func (orig testConfig) ForStandardTree(root tree.Node) testConfig {
	return testConfig{displayFunc: func(config tree.DisplayTreeConfig) []string {
		return tree.DisplayTree(root, config)
	}, displayConfig: orig.displayConfig}
}
func (orig testConfig) ForBinaryTree(root tree.BinaryNode) testConfig {
	return testConfig{displayFunc: func(config tree.DisplayTreeConfig) []string {
		return tree.DisplayBinaryTree(root, config)
	}, displayConfig: orig.displayConfig}
}
func (orig testConfig) WithDisplayType(value tree.DisplayType) testConfig {
	return testConfig{displayFunc: orig.displayFunc, displayConfig: orig.displayConfig.WithDisplayType(value)}
}

func (orig testConfig) WithDisplayCharacterType(value tree.CharacterType) testConfig {
	return testConfig{displayFunc: orig.displayFunc, displayConfig: orig.displayConfig.WithCharacterType(value)}
}

func testDisplayTypes(t *testing.T, root tree.Node, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	testConfig := NewTestConfig().ForStandardTree(root)
	testDisplayTypesImpl(t, testConfig,  expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func testDisplayTypesBinary(t *testing.T, root tree.BinaryNode, expectedTopDown []string, expectedBalanced []string, expectedBottomUp []string) {
	testConfig := NewTestConfig().ForBinaryTree(root)
	// Favour up ignored, treat as balanced
	testDisplayTypesImpl(t, testConfig,  expectedTopDown, expectedBalanced, expectedBalanced, expectedBottomUp)
}

func testDisplayTypesImpl(t *testing.T, testConfig testConfig, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	t.Run("Top down", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(tree.TopDown), expectedTopDown) })
	t.Run("Balanced", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(tree.Balanced), expectedBalanced) })
	t.Run("Balanced favour top", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(tree.BalancedFavourTop), expectedBalancedFavourTop) })
	t.Run("Bottom up", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(tree.BottomUp), expectedBottomUp) })
}

func testImpl(t *testing.T, testConfig testConfig, expected []string) {
	// Act
	result := testConfig.displayFunc(testConfig.displayConfig)
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
func createEmptyBinaryNode(left tree.BinaryNode, right tree.BinaryNode) tree.BinaryNode {
	return testBinaryTreeNode{nil, left, right}
}
func createBinaryLeafNode(value string) tree.BinaryNode {
	return createBinaryNode(value, nil, nil)
}
