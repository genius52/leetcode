package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node559 struct {
	Val int
	Children []*Node559
}

func maxDepth559(root *Node559) int {
	if root == nil{
		return 0
	}
	var max_depth int = 0
	for _,child := range root.Children{
		max_depth = max_int(max_depth,maxDepth559(child))
	}
	return 1 + max_depth
}