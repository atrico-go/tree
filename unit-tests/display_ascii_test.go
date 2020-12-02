package unit_tests

import (
	"testing"

	"github.com/atrico-go/tree"
)

func TestDisplayAscii(t *testing.T) {
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
		"- \\",
		"  +- Child1",
		"  |  +- Child11",
		"  |  |  \\- Child111",
		"  |  \\- \\",
		"  |     \\- Child121",
		"  \\- Child2",
		"     +- Child21",
		"     |  \\- Child211",
		"     \\- Child22",
		"        \\- Child221",
	}
	expectedBalanced := []string{
		"     /- Child11",
		"     |  \\- Child111",
		"  /- Child1",
		"  |  \\- \\",
		"  |     \\- Child121",
		"- +",
		"  |  /- Child21",
		"  |  |  \\- Child211",
		"  \\- Child2",
		"     \\- Child22",
		"        \\- Child221",
	}
	expectedBalancedFavourTop := []string{
		"        /- Child111",
		"     /- Child11",
		"  /- Child1",
		"  |  |  /- Child121",
		"  |  \\- /",
		"- +",
		"  |     /- Child211",
		"  |  /- Child21",
		"  \\- Child2",
		"     |  /- Child221",
		"     \\- Child22",
	}
	expectedBottomUp := []string{
		"        /- Child111",
		"     /- Child11",
		"     |  /- Child121",
		"     +- /",
		"  /- Child1",
		"  |     /- Child211",
		"  |  /- Child21",
		"  |  |  /- Child221",
		"  |  +- Child22",
		"  +- Child2",
		"- /",
	}
	testConfig := testConfig{}.ForStandardTree(root).WithDisplayCharacterType(tree.ASCII)
	testDisplayTypesImpl(t, testConfig, expectedTopDown, expectedBalanced, expectedBalancedFavourTop, expectedBottomUp)
}

func TestDisplayBinaryAscii(t *testing.T) {
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
		"- Root",
		"  +- ChildL",
		"  |  +- ChildLL",
		"  |  |  +- ChildLLL",
		"  |  |  \\- ",
		"  |  \\- ChildLR",
		"  |     +- ",
		"  |     \\- ChildLRR",
		"  \\- ChildR",
		"     +- ChildRL",
		"     |  +- ",
		"     |  \\- ChildRLR",
		"     \\- ChildRR",
		"        +- ChildRRL",
		"        \\- ",
	}
	expectedBalanced := []string{
		"        /- ChildLLL",
		"     /- ChildLL",
		"     |  \\- ",
		"  /- ChildL",
		"  |  |  /- ",
		"  |  \\- ChildLR",
		"  |     \\- ChildLRR",
		"- Root",
		"  |     /- ",
		"  |  /- ChildRL",
		"  |  |  \\- ChildRLR",
		"  \\- ChildR",
		"     |  /- ChildRRL",
		"     \\- ChildRR",
		"        \\- ",
	}
	expectedBottomUp := []string{
		"        /- ChildLLL",
		"        +- ",
		"     /- ChildLL",
		"     |  /- ",
		"     |  +- ChildLRR",
		"     +- ChildLR",
		"  /- ChildL",
		"  |     /- ",
		"  |     +- ChildRLR",
		"  |  /- ChildRL",
		"  |  |  /- ChildRRL",
		"  |  |  +- ",
		"  |  +- ChildRR",
		"  +- ChildR",
		"- Root",
	}
	testConfig := testConfig{}.ForBinaryTree(root).WithDisplayCharacterType(tree.ASCII)
	testDisplayTypesImpl(t, testConfig, expectedTopDown, expectedBalanced, expectedBalanced, expectedBottomUp)
}

