package tree

import (
	"fmt"
	"strings"
)

func DisplayTree(root Node, config DisplayTreeConfig) []string {
	highlights := make([]Node, 0)
	if config.Highlight != nil {
		if found, foundHl := FindValue(root, config.Highlight); found {
			highlights = foundHl
		}
	}
	return displayTree(newRootDetails(root, highlights, config), config)
}

func DisplayBinaryTree(root BinaryNode, config DisplayTreeConfig) []string {
	return DisplayTree(BinaryNodeTreeWrapper{root}, config)
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
			text.WriteString(GetBoxCharHeavy(BoxNone, BoxSingle, BoxNone, BoxSingle, details.isOnHighlightPath())) // ┌
		} else if details.parentPosition == belowParent && details.siblingLocation == lastSibling {
			text.WriteString(GetBoxCharHeavy(BoxSingle, BoxNone, BoxNone, BoxSingle, details.isOnHighlightPath())) // └
		} else if details.isHighlighted() {
			text.WriteString(GetBoxChar(
				BoxSingle.MakeHeavy(details.parentPosition == belowParent),
				BoxSingle.MakeHeavy(details.parentPosition == aboveParent),
				BoxNone,
				BoxHeavy)) // ├
		} else if details.isOnHighlightPath() {
			text.WriteString(GetBoxChar(
				BoxSingle.MakeHeavy(details.parentPosition == belowParent),
				BoxSingle.MakeHeavy(details.parentPosition == aboveParent),
				BoxNone,
				BoxHeavy)) // ├
		} else {
			passThrough := (details.highlightPosition == siblingHighlightAbove && details.parentPosition == aboveParent) ||
				(details.highlightPosition == siblingHighlightBelow && details.parentPosition == belowParent)
			text.WriteString(GetBoxChar(
				BoxSingle.MakeHeavy(passThrough),
				BoxSingle.MakeHeavy(passThrough),
				BoxNone,
				BoxSingle)) // ├
		}
	}
	// Dash
	text.WriteString(GetBoxCharHeavy(BoxNone, BoxNone, BoxSingle, BoxSingle, details.isOnHighlightPath())) // -
	// Name of node (and branch chars)
	up := IfThenElseBoxType(details.children.above > 0, BoxSingle.MakeHeavy(details.hasHighlightChildAbove()), BoxNone)
	down := IfThenElseBoxType(details.children.below > 0, BoxSingle.MakeHeavy(details.hasHighlightChildBelow()), BoxNone)
	left := IfThenElseBoxType(details.isOnHighlightPath(), BoxHeavy, BoxSingle)
	value := details.node.NodeValue()
	right := IfThenElseBoxType(value != nil || len(details.node.Children()) == 0, BoxSingle.MakeHeavy(details.isHighlighted()), BoxNone)
	text.WriteString(GetBoxChar(up, down, left, right))
	if value != nil {
		text.WriteString(fmt.Sprintf("%s%v", GetBoxChar(BoxNone, BoxNone, BoxNone, BoxNone), value))
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
				highlight := (details.isOnHighlightPath() && (
					(context == belowParent && details.parentPosition == aboveParent) ||
					(context == aboveParent && details.parentPosition == belowParent))) ||
					(context == belowParent && details.hasHighlightSiblingBelow()) ||
				(context == aboveParent && details.hasHighlightSiblingAbove())
				prefix.WriteString(GetBoxCharHeavy(BoxSingle, BoxSingle, BoxNone, BoxNone, highlight)) // |
			} else {
				prefix.WriteString(GetBoxChar(BoxNone, BoxNone, BoxNone, BoxNone)) // space
			}
		}
		prefix.WriteString(GetBoxChar(BoxNone, BoxNone, BoxNone, BoxNone)) // space
	}
	return prefix.String()
}
