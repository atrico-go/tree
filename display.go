package tree

import (
	"fmt"
)

type DisplayTreeConfig struct {
	Type DisplayType
}

type DisplayType int

const (
	TopDown           DisplayType = iota
	Balanced          DisplayType = iota
	BalancedFavourTop DisplayType = iota
	BottomUp          DisplayType = iota
)

func DisplayTree(root Node, config DisplayTreeConfig) []string {
	return displayTree(root, childDetails{Root, Above}, "", config)
}

func displayTree(node Node, childDetails childDetails, prefix string, config DisplayTreeConfig) []string {
	lines := make([]string, 0, 1)
	above := aboveChildren(node, config)
	newPrefix := addPrefix(prefix, childDetails.Above())
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(node.Children()[i], getChildDetails(i, node, Above), newPrefix, config)...)
	}
	lines = append(lines, fmt.Sprintf("%s%s─ %v", prefix, nodeChar(childDetails), node.NodeValue()))
	newPrefix = addPrefix(prefix, childDetails.Below())
	for i := above; i < len(node.Children()); i++ {
		lines = append(lines, displayTree(node.Children()[i], getChildDetails(i, node, Below), newPrefix, config)...)
	}
	return lines
}

type childRank int
type childPosition int

const (
	Root  childRank     = iota
	First childRank     = iota
	Mid   childRank     = iota
	Last  childRank     = iota
	Above childPosition = iota
	Below childPosition = iota
)

type childDetails struct {
	Rank     childRank
	Position childPosition
}

func (d childDetails) IsRoot() bool {
	return d.Rank == Root
}

func (d childDetails) Above() childDetails {
	return childDetails{d.Rank, Above}
}

func (d childDetails) Below() childDetails {
	return childDetails{d.Rank, Below}
}

func getChildDetails(idx int, parent Node, pos childPosition) childDetails {
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
	return childDetails{rank, pos}
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

func nodeChar(childDetails childDetails) string {
	if childDetails.IsRoot() {
		return ""
	}
	if childDetails.Rank == First && childDetails.Position == Above {
		return "┌"
	}
	if childDetails.Rank == Last && childDetails.Position == Below {
		return "└"
	}
	return "├"
}

func addPrefix(previous string, childDetails childDetails) string {
	ch := "│"
	if childDetails.IsRoot() {
		ch = ""
	} else if childDetails.Rank == First && childDetails.Position == Above {
		ch = " "
	} else if childDetails.Rank == Last && childDetails.Position == Below {
		ch = " "
	}
	return fmt.Sprintf("%s%s  ", previous, ch)
}
