package treedisplay

import "github.com/atrico-go/tree"

type parentPosition string
type siblingLocation string

const (
	aboveParent  parentPosition  = "Above parent" // Node is above the parent
	belowParent  parentPosition  = "Below parent" // Node is below the parent
	noParent     parentPosition  = "No Parent"
	firstSibling siblingLocation = "First sibling"
	midSibling   siblingLocation = "Mid sibling"
	lastSibling  siblingLocation = "Last sibling"
)

func (pp parentPosition) String() string {
	return string(pp)
}
func (sl siblingLocation) String() string {
	return string(sl)
}

type numberSplit struct {
	above int
	below int
	total int
}

type nodeDetails struct {
	parent              *nodeDetails
	node                tree.Node
	parentPosition      parentPosition
	siblingLocation     siblingLocation
	children            numberSplit
	highlightPosition   highlightPosition
	highlightChildIdx   int
	remainingHighlights []tree.Node
}

func newRootDetails(node tree.Node, highlights []tree.Node, config DisplayTreeConfig) nodeDetails {
	highlightPosition, highlightChildIdx, highlightsTail := processHighlightRoot(node, highlights, config)
	return nodeDetails{nil, node, noParent, midSibling, numberOfChildren(node, config), highlightPosition, highlightChildIdx, highlightsTail}
}

func newNodeDetails(parent *nodeDetails, idx int, config DisplayTreeConfig) nodeDetails {
	node := parent.node.Children()[idx]
	highlightPosition, highlightChildIdx, highlightsTail := processHighlightNode(node, *parent, idx, config)

	siblings := numberOfChildren(parent.node, config)
	var parentPosition parentPosition
	siblingLocation := midSibling
	if idx < siblings.above {
		parentPosition = aboveParent
		if idx == 0 {
			siblingLocation = firstSibling
		}
	} else {
		parentPosition = belowParent
		if idx == len(parent.node.Children())-1 {
			siblingLocation = lastSibling
		}
	}
	return nodeDetails{parent, node, parentPosition, siblingLocation, numberOfChildren(node, config), highlightPosition, highlightChildIdx, highlightsTail}
}

func (d nodeDetails) isRoot() bool {
	return d.parentPosition == noParent
}

func (d nodeDetails) isHighlighted() bool {
	return d.highlightPosition == iAmHighlight
}

func (d nodeDetails) hasHighlightChildAbove() bool {
	return d.highlightPosition == childHighlightAbove
}
func (d nodeDetails) hasHighlightChildBelow() bool {
	return d.highlightPosition == childHighlightBelow
}
func (d nodeDetails) hasHighlightChild() bool {
	return d.hasHighlightChildAbove() || d.hasHighlightChildBelow()
}
func (d nodeDetails) isOnHighlightPath() bool {
	return d.isHighlighted() || d.hasHighlightChild()
}
func (d nodeDetails) hasHighlightSiblingAbove() bool {
	return d.highlightPosition == siblingHighlightAbove
}
func (d nodeDetails) hasHighlightSiblingBelow() bool {
	return d.highlightPosition == siblingHighlightBelow
}

// ---------------------------------------------------------------
// Internal functions
// ---------------------------------------------------------------
func numberOfChildren(node tree.Node, config DisplayTreeConfig) numberSplit {
	total := len(node.Children())
	var above int
	switch config.Type {
	case TopDown:
		above = 0
	case Balanced:
		above = total / 2
	case BalancedFavourTop:
		above = (total + 1) / 2
	case BottomUp:
		above = total
	default:
		panic("Unrecognised display type")
	}
	return numberSplit{above, total - above, total}
}
