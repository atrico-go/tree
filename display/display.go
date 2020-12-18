package treedisplay

import (
	"fmt"
	"github.com/atrico-go/console"
	"github.com/atrico-go/tree"
	"strings"
)

func DisplayTree(root tree.Node, config DisplayTreeConfig) []string {
	highlights := make([]tree.Node, 0)
	if config.Highlight != nil {
		if found, foundHl := tree.FindValue(root, config.Highlight); found {
			highlights = foundHl
		}
	}
	return displayTree(newRootDetails(root, highlights, config), config)
}

func DisplayBinaryTree(root tree.BinaryNode, config DisplayTreeConfig) []string {
	return DisplayTree(tree.BinaryNodeTreeWrapper{BinaryNode: root}, config)
}

func displayTree(details nodeDetails, config DisplayTreeConfig) []string {
	lines := make([]string, 0, 1)
	for i := 0; i < details.children.above; i++ {
		lines = append(lines, displayTree(newNodeDetails(&details, i, config), config)...)
	}
	lines = append(lines, formatNode(details))
	for i := details.children.above; i < details.children.total; i++ {
		lines = append(lines, displayTree(newNodeDetails(&details, i, config), config)...)
	}
	return lines
}

func formatNode(details nodeDetails) string {
	text := strings.Builder{}
	// Prefix (previous branches)
	text.WriteString(formatPrefix(details.parent, details.parentPosition))
	// Attachment to tree structure
	if !details.isRoot() {
		if details.parentPosition == aboveParent && details.siblingLocation == firstSibling {
			text.WriteString(console.MustGetBoxChar(false, true, false, true, console.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // ┌
		} else if details.parentPosition == belowParent && details.siblingLocation == lastSibling {
			text.WriteString(console.MustGetBoxChar(true, false, false, true, console.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // └
		} else if details.isHighlighted() {
			text.WriteString(console.MustGetBoxCharMixed(
				console.BoxSingle.HeavyIf(details.parentPosition == belowParent),
				console.BoxSingle.HeavyIf(details.parentPosition == aboveParent),
				console.BoxNone,
				console.BoxHeavy)) // ├
		} else if details.isOnHighlightPath() {
			text.WriteString(console.MustGetBoxCharMixed(
				console.BoxSingle.HeavyIf(details.parentPosition == belowParent),
				console.BoxSingle.HeavyIf(details.parentPosition == aboveParent),
				console.BoxNone,
				console.BoxHeavy)) // ├
		} else {
			passThrough := (details.highlightPosition == siblingHighlightAbove && details.parentPosition == aboveParent) ||
				(details.highlightPosition == siblingHighlightBelow && details.parentPosition == belowParent)
			text.WriteString(console.MustGetBoxCharMixed(
				console.BoxSingle.HeavyIf(passThrough),
				console.BoxSingle.HeavyIf(passThrough),
				console.BoxNone,
				console.BoxSingle)) // ├
		}
	}
	// Dash
	text.WriteString(console.MustGetBoxChar(false, false, true, true, console.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // -
	// Name of node (and branch chars)
	up := console.ConditionalBoxType(details.children.above > 0, console.BoxSingle.HeavyIf(details.hasHighlightChildAbove()), console.BoxNone)
	down := console.ConditionalBoxType(details.children.below > 0, console.BoxSingle.HeavyIf(details.hasHighlightChildBelow()), console.BoxNone)
	left := console.BoxSingle.HeavyIf(details.isOnHighlightPath())
	value := details.node.NodeValue()
	right := console.ConditionalBoxType(value != nil || len(details.node.Children()) == 0, console.BoxSingle.HeavyIf(details.isHighlighted()), console.BoxNone)
	text.WriteString(console.MustGetBoxCharMixed(up, down, left, right))
	if value != nil {
		text.WriteString(fmt.Sprintf("%s%v", console.MustGetBoxChar(false,false,false,false, console.BoxNone), value))
	}
	// TODO - Start
	// text.WriteString(fmt.Sprintf(":par=%v,sib=%v,hl=%v,hlIx=%v", details.parentPosition, details.siblingLocation, details.highlightPosition, details.highlightChildIdx))
	// TODO - End

	return text.String()
}

func formatPrefix(details *nodeDetails, context parentPosition) string {
	prefix := strings.Builder{}
	if details != nil {
		// Root level has no prefix
		if !details.isRoot() {
			prefix.WriteString(formatPrefix(details.parent, details.parentPosition))
			bypass := false
			switch details.siblingLocation {
			// First has bypass if calling child is below
			case firstSibling:
				bypass = context == belowParent
			// Mid sibling must have bypass
			case midSibling:
				bypass = true
			// Last has bypass if calling child is above
			case lastSibling:
				bypass = context == aboveParent
			}
			if bypass {
				highlight := (context == belowParent && details.isOnHighlightPath() && details.parentPosition == aboveParent) ||
					(context == aboveParent && details.isOnHighlightPath() && details.parentPosition == belowParent) ||
					(context == belowParent && details.hasHighlightSiblingBelow()) ||
					(context == aboveParent && details.hasHighlightSiblingAbove())
				prefix.WriteString(console.MustGetBoxChar(true, true, false, false, console.BoxSingle.HeavyIf(highlight))) // |
			} else {
				prefix.WriteString(console.MustGetBoxChar(false,false,false,false, console.BoxNone)) // space
			}
		}
		prefix.WriteString(console.MustGetBoxChar(false,false,false,false, console.BoxNone)) // space
	}
	return prefix.String()
}
