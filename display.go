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
	lines := make([]string, 0, 1)
	above := aboveChildren(root, config)
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(root.Children()[i], getChildRank(i, root, Above), Above, "  ", config)...)
	}
	lines = append(lines, fmt.Sprintf("─ %v", root.NodeValue()))
	for i := above; i < len(root.Children()); i++ {
		lines = append(lines, displayTree(root.Children()[i], getChildRank(i, root, Below), Below, "  ", config)...)
	}
	return lines
}

func displayTree(node Node, childRank childRank, childPosition childPosition, prefix string, config DisplayTreeConfig) []string {
	lines := make([]string, 0, 1)
	above := aboveChildren(node, config)
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(node.Children()[i], getChildRank(i, node, Above), Above, fmt.Sprintf("%s%s  ", prefix, prefixChar(childRank, Above)), config)...)
	}
	lines = append(lines, fmt.Sprintf("%s%s─ %v", prefix, nodeChar(childRank, childPosition), node.NodeValue()))
	for i := above; i < len(node.Children()); i++ {
		lines = append(lines, displayTree(node.Children()[i], getChildRank(i, node, Below), Below, fmt.Sprintf("%s%s  ", prefix, prefixChar(childRank, Below)), config)...)
	}
	return lines
}

type childRank int

const (
	First childRank = iota
	Mid   childRank = iota
	Last  childRank = iota
)

type childPosition int

const (
	Above childPosition = iota
	Below childPosition = iota
)

func getChildRank(idx int, parent Node, pos childPosition) childRank {
	if pos == Above {
		if idx == 0 {
			return First
		} else if idx == len(parent.Children())-1 {
			return Last
		}
	}
	if pos == Below {
		if idx == len(parent.Children())-1 {
			return Last
		} else if idx == 0 {
			return First
		}
	}
	return Mid
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

func nodeChar(childRank childRank, childPos childPosition) string {
	if childRank == First && childPos == Above {
		return "┌"
	}
	if childRank == Last && childPos == Below {
		return "└"
	}
	return "├"
}

func prefixChar(childRank childRank, childPos childPosition) string {
	if childRank == First && childPos == Above {
		return " "
	}
	if childRank == Last && childPos == Below {
		return " "
	}
	return "│"
}
