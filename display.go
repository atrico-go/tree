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
	return displayTree(root, newRootDetails(), "", config)
}

func displayTree(node Node, childDetails nodeDetails, prefix string, config DisplayTreeConfig) []string {
	lines := make([]string, 0, 1)
	above := aboveChildren(node, config)
	newPrefix := addPrefix(prefix, childDetails.Above())
	for i := 0; i < above; i++ {
		lines = append(lines, displayTree(node.Children()[i], newChildDetails(i, node, Above), newPrefix, config)...)
	}
	lines = append(lines, fmt.Sprintf("%s%s─ %v", prefix, nodeChar(childDetails), nodeValue(node, config)))
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

func nodeValue(node Node, config DisplayTreeConfig) string {
	value := node.NodeValue()
	if value != nil {
		return fmt.Sprintf("%v", value)
	}
	children := childLocations(node, config)
	if children.Above {
		if children.Below {
			return "┤"
		} else {
			return "┘"
		}
	} else {
		if children.Below {
			return "┐"
		} else {
			return ""
		}
	}
}
func nodeChar(details nodeDetails) string {
	if details.IsRoot {
		return ""
	}
	if details.Rank == First && details.Position == Above {
		return "┌"
	}
	if details.Rank == Last && details.Position == Below {
		return "└"
	}
	return "├"
}

func addPrefix(previous string, details nodeDetails) string {
	ch := "│"
	if details.IsRoot {
		ch = ""
	} else if details.Rank == First && details.Position == Above {
		ch = " "
	} else if details.Rank == Last && details.Position == Below {
		ch = " "
	}
	return fmt.Sprintf("%s%s  ", previous, ch)
}
