package treedisplay

import (
	"fmt"
	"strings"

	"github.com/atrico-go/console/box_drawing"

	"github.com/atrico-go/tree"
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
			text.WriteRune(box_drawing.MustGetBoxChar(false, true, false, true, box_drawing.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // ┌
		} else if details.parentPosition == belowParent && details.siblingLocation == lastSibling {
			text.WriteRune(box_drawing.MustGetBoxChar(true, false, false, true, box_drawing.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // └
		} else if details.isHighlighted() {
			text.WriteRune(box_drawing.MustGetBoxCharMixed(box_drawing.BoxParts{
				Up:    box_drawing.BoxSingle.HeavyIf(details.parentPosition == belowParent),
				Down:  box_drawing.BoxSingle.HeavyIf(details.parentPosition == aboveParent),
				Left:  box_drawing.BoxNone,
				Right: box_drawing.BoxHeavy})) // ├
		} else if details.isOnHighlightPath() {
			text.WriteRune(box_drawing.MustGetBoxCharMixed(box_drawing.BoxParts{
				Up:    box_drawing.BoxSingle.HeavyIf(details.parentPosition == belowParent),
				Down:  box_drawing.BoxSingle.HeavyIf(details.parentPosition == aboveParent),
				Left:  box_drawing.BoxNone,
				Right: box_drawing.BoxHeavy})) // ├
		} else {
			passThrough := (details.highlightPosition == siblingHighlightAbove && details.parentPosition == aboveParent) ||
				(details.highlightPosition == siblingHighlightBelow && details.parentPosition == belowParent)
			text.WriteRune(box_drawing.MustGetBoxCharMixed(box_drawing.BoxParts{
				Up:    box_drawing.BoxSingle.HeavyIf(passThrough),
				Down:  box_drawing.BoxSingle.HeavyIf(passThrough),
				Left:  box_drawing.BoxNone,
				Right: box_drawing.BoxSingle})) // ├
		}
	}
	// Dash
	text.WriteRune(box_drawing.MustGetBoxChar(false, false, true, true, box_drawing.BoxSingle.HeavyIf(details.isOnHighlightPath()))) // -
	// Name of node (and branch chars)
	parts := box_drawing.BoxParts{}
	parts.Up = box_drawing.ConditionalBoxType(details.children.above > 0, box_drawing.BoxSingle.HeavyIf(details.hasHighlightChildAbove()), box_drawing.BoxNone)
	parts.Down = box_drawing.ConditionalBoxType(details.children.below > 0, box_drawing.BoxSingle.HeavyIf(details.hasHighlightChildBelow()), box_drawing.BoxNone)
	parts.Left = box_drawing.BoxSingle.HeavyIf(details.isOnHighlightPath())
	value := details.node.NodeValue()
	parts.Right = box_drawing.ConditionalBoxType(value != nil || len(details.node.Children()) == 0, box_drawing.BoxSingle.HeavyIf(details.isHighlighted()), box_drawing.BoxNone)
	text.WriteRune(box_drawing.MustGetBoxCharMixed(parts))
	if value != nil {
		text.WriteRune(box_drawing.MustGetBoxChar(false, false, false, false, box_drawing.BoxNone))
		text.WriteString(fmt.Sprintf("%v", value))
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
				prefix.WriteRune(box_drawing.MustGetBoxChar(true, true, false, false, box_drawing.BoxSingle.HeavyIf(highlight))) // |
			} else {
				prefix.WriteRune(box_drawing.MustGetBoxChar(false, false, false, false, box_drawing.BoxNone)) // space
			}
		}
		prefix.WriteRune(box_drawing.MustGetBoxChar(false, false, false, false, box_drawing.BoxNone)) // space
	}
	return prefix.String()
}
