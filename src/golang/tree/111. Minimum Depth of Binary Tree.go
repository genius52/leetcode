package tree

import "container/list"

func min_int(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return min_int(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q list.List
	q.PushBack(root)
	var depth int = 1
	for q.Len() > 0 {
		var l int = q.Len()
		for i := 0; i < l; i++ {
			var node *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}
