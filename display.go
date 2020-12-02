package tree

import (
	"fmt"
	"strings"
)

func DisplayTree(root Node, config DisplayTreeConfig) []string {
	return displayTree(root, newRootDetails(), "", config)
}

func DisplayBinaryTree(root BinaryNode, config DisplayTreeConfig) []string {
	return DisplayTree(binaryTreeNodeWrapper{root}, config)
}

func displayTree(node Node, childDetails nodeDetails, prefix string, config DisplayTreeConfig) []string {
	lines := make([]string, 0, 1)
	above := aboveChildren(node, config)
	newPrefix := addPrefix(prefix, childDetails.Above())
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(node.Children()[i], newChildDetails(i, node, Above), newPrefix, config)...)
	}
	lines = append(lines, fmt.Sprintf("%s%s%s%v", prefix, nodeChar(childDetails), getBoxChar(BoxNone, BoxNone, BoxSingle, BoxSingle), nodeValue(node, config)))
	newPrefix = addPrefix(prefix, childDetails.Below())
	for i := above; i < len(node.Children()); i++ {
		lines = append(lines, displayTree(node.Children()[i], newChildDetails(i, node, Below), newPrefix, config)...)
	}
	return lines
}

type childRank int
type childPosition int

const (
	First childRank     = iota
	Mid   childRank     = iota
	Last  childRank     = iota
	Above childPosition = iota
	Below childPosition = iota
)

type nodeDetails struct {
	IsRoot   bool
	Rank     childRank
	Position childPosition
}

func (d nodeDetails) Above() nodeDetails {
	return nodeDetails{d.IsRoot, d.Rank, Above}
}

func (d nodeDetails) Below() nodeDetails {
	return nodeDetails{d.IsRoot, d.Rank, Below}
}

func newChildDetails(idx int, parent Node, pos childPosition) nodeDetails {
	rank := Mid
	if pos == Above {
		if idx == 0 {
			rank = First
		} else if idx == len(parent.Children())-1 {
			rank = Last
		}
	}
	if pos == Below {
		if idx == len(parent.Children())-1 {
			rank = Last
		} else if idx == 0 {
			rank = First
		}
	}
	return nodeDetails{false, rank, pos}
}

func newRootDetails() nodeDetails {
	return nodeDetails{true, 0, 0}
}

func aboveChildren(node Node, config DisplayTreeConfig) int {
	switch config.Type {
	case TopDown:
		return 0
	case Balanced:
		return len(node.Children()) / 2
	case BalancedFavourTop:
		return (len(node.Children()) + 1) / 2
	case BottomUp:
		return len(node.Children())
	default:
		panic("Unrecognised display type")
	}
}

type childrenLocations struct {
	Above bool
	Below bool
}

func childLocations(node Node, config DisplayTreeConfig) childrenLocations {
	above := aboveChildren(node, config)
	return childrenLocations{Above: above > 0, Below: above < len(node.Children())}
}
func ifThenElse(c bool, t BoxType, f BoxType) BoxType {
	if c {
		return t
	} else {
		return f
	}
}

func nodeValue(node Node, config DisplayTreeConfig) string {
	text := strings.Builder{}
	children := childLocations(node, config)
	up := ifThenElse(children.Above, BoxSingle, BoxNone)
	down := ifThenElse(children.Below, BoxSingle, BoxNone)
	left := BoxSingle
	value := node.NodeValue()
	right := ifThenElse(value != nil || !(children.Above || children.Below), BoxSingle, BoxNone)
	text.WriteString(getBoxChar(up, down, left, right))
	if value != nil {
		text.WriteString(fmt.Sprintf("%s%v", getBoxChar(BoxNone,BoxNone,BoxNone,BoxNone),value))
	}
	return text.String()
}

func nodeChar(details nodeDetails) string {
	if details.IsRoot {
		return ""
	}
	if details.Rank == First && details.Position == Above {
		return getBoxChar(BoxNone, BoxSingle, BoxNone, BoxSingle)
	}
	if details.Rank == Last && details.Position == Below {
		return getBoxChar(BoxSingle, BoxNone, BoxNone, BoxSingle)
	}
	return getBoxChar(BoxSingle, BoxSingle, BoxNone, BoxSingle)
}

func addPrefix(previous string, details nodeDetails) string {
	ch := getBoxChar(BoxSingle, BoxSingle, BoxNone, BoxNone)
	if details.IsRoot {
		ch = ""
	} else if details.Rank == First && details.Position == Above {
		ch = getBoxChar(BoxNone, BoxNone, BoxNone, BoxNone)
	} else if details.Rank == Last && details.Position == Below {
		ch = getBoxChar(BoxNone, BoxNone, BoxNone, BoxNone)
	}
	return fmt.Sprintf("%s%s%s", previous, ch,  getBoxChar(BoxNone, BoxNone, BoxNone, BoxNone))
}

// ---------------------------------------------------------------------
// Binary tree wrapper
// ---------------------------------------------------------------------
type binaryTreeNodeWrapper struct {
	BinaryNode
}

func (n binaryTreeNodeWrapper) NodeValue() interface{} {
	if n.BinaryNode != nil {
		return n.BinaryNode.NodeValue()
	}
	return nil
}

func (n binaryTreeNodeWrapper) Children() []Node {
	if n.BinaryNode != nil && (n.BinaryNode.Left() != nil || n.BinaryNode.Right() != nil) {
		return []Node{binaryTreeNodeWrapper{n.BinaryNode.Left()}, binaryTreeNodeWrapper{n.BinaryNode.Right()}}
	}
	return make([]Node, 0)
}
