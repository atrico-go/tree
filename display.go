package tree

import (
	"fmt"
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
	newPrefix := addPrefix(prefix, childDetails.Above(), config)
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(node.Children()[i], newChildDetails(i, node, Above), newPrefix, config)...)
	}
	lines = append(lines, fmt.Sprintf("%s%s%s %v", prefix, nodeChar(childDetails, config), config.getChar(BoxNone, BoxNone, BoxSingle, BoxSingle), nodeValue(node, config)))
	newPrefix = addPrefix(prefix, childDetails.Below(), config)
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

func nodeValue(node Node, config DisplayTreeConfig) string {
	value := node.NodeValue()
	if value != nil {
		return fmt.Sprintf("%v", value)
	}
	children := childLocations(node, config)
	if children.Above {
		if children.Below {
			return config.getChar(BoxSingle, BoxSingle, BoxSingle, BoxNone)
		} else {
			return config.getChar(BoxSingle, BoxNone, BoxSingle, BoxNone)
		}
	} else {
		if children.Below {
			return config.getChar(BoxNone, BoxSingle, BoxSingle, BoxNone)
		} else {
			return ""
		}
	}
}

func nodeChar(details nodeDetails, config DisplayTreeConfig) string {
	if details.IsRoot {
		return ""
	}
	if details.Rank == First && details.Position == Above {
		return config.getChar(BoxNone, BoxSingle, BoxNone, BoxSingle)
	}
	if details.Rank == Last && details.Position == Below {
		return config.getChar(BoxSingle, BoxNone, BoxNone, BoxSingle)
	}
	return config.getChar(BoxSingle, BoxSingle, BoxNone, BoxSingle)
}

func addPrefix(previous string, details nodeDetails, config DisplayTreeConfig) string {
	ch := config.getChar(BoxSingle, BoxSingle, BoxNone, BoxNone)
	if details.IsRoot {
		ch = ""
	} else if details.Rank == First && details.Position == Above {
		ch = config.getChar(BoxNone, BoxNone, BoxNone, BoxNone)
	} else if details.Rank == Last && details.Position == Below {
		ch = config.getChar(BoxNone, BoxNone, BoxNone, BoxNone)
	}
	return fmt.Sprintf("%s%s%s%s", previous, ch, config.getChar(BoxNone, BoxNone, BoxNone, BoxNone),config.getChar(BoxNone, BoxNone, BoxNone, BoxNone))
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
