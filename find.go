package tree

// Find the first instance of value
// Return true if found and index of child where value was found (-1 if found at this node)
func FindValue(node Node, value interface{}) (found bool, path []Node) {
	if node.NodeValue() != nil && node.NodeValue() == value {
		return true, []Node{node}
	}
	for _,child := range node.Children() {
		result,nodes := FindValue(child, value)
		if result {
			return true, append([]Node{node},nodes...)
		}
	}
	return false,nil
}
