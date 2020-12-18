package tree

type highlightPosition string

const (
	noHighlight           highlightPosition = "No Highlight"       // No highlight in this subtree
	iAmHighlight          highlightPosition = "I am the highlight" // This node is highlight
	childHighlightAbove   highlightPosition = "Child above"        // child above is highlight
	childHighlightBelow   highlightPosition = "Child below"        // child below is highlight
	siblingHighlightAbove highlightPosition = "Sibling above"      // sibling above is highlight
	siblingHighlightBelow highlightPosition = "Sibling below"      // sibling below is highlight
)

func (hl highlightPosition) String() string {
	return string(hl)
}

func processHighlightRoot(node Node, highlights []Node, config DisplayTreeConfig) (highlightPosition highlightPosition, highlightChildIdx int, highlightsTail []Node) {
	if len(highlights) > 0 && node.Equals(highlights[0]) {
		highlightsTail = highlights[1:]
		if len(highlightsTail) == 0 {
			highlightPosition = iAmHighlight
		} else {
			next := highlightsTail[0]
			_, highlightChildIdx = Contains(node.Children(), next)
			childrenAbove := numberOfChildren(node, config).above
			if highlightChildIdx < childrenAbove {
				highlightPosition = childHighlightAbove
			} else {
				highlightPosition = childHighlightBelow
			}
		}
	} else {
		highlightPosition = noHighlight
		highlightsTail = make([]Node, 0)
	}
	return highlightPosition, highlightChildIdx, highlightsTail
}

func processHighlightNode(node Node, parent nodeDetails, idx int, config DisplayTreeConfig) (highlightPosition highlightPosition, highlightChildIdx int, highlightsTail []Node) {
	highlightPosition, highlightChildIdx, highlightsTail = processHighlightRoot(node, parent.remainingHighlights, config)
	if highlightPosition == noHighlight {
		if parent.hasHighlightChild() {
			// Are both this child and highlight child above or below
			if (idx < parent.children.above) == parent.hasHighlightChildAbove() {
				if idx < parent.highlightChildIdx {
					highlightPosition = siblingHighlightBelow
				} else {
					highlightPosition = siblingHighlightAbove
				}
			}
		}
	}
	return highlightPosition, highlightChildIdx, highlightsTail
}
