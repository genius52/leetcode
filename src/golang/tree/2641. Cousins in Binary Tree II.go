package tree

import "container/list"

func replaceValueInTree(root *TreeNode) *TreeNode {
	var q list.List
	q.PushBack(root)
	var depth_sum map[int]int = make(map[int]int)
	var depth int = 0
	for q.Len() > 0 {
		var cur_len int = q.Len()
		var sum int = 0
		for i := 0; i < cur_len; i++ {
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			sum += cur.Val
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		depth_sum[depth] = sum
		depth++
	}
	q.PushBack(root)
	depth = 1
	root.Val = 0
	for q.Len() > 0 {
		var cur_len int = q.Len()
		for i := 0; i < cur_len; i++ {
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			var pre_left_val int = 0
			if cur.Left != nil {
				pre_left_val = cur.Left.Val
				cur.Left.Val = depth_sum[depth] - cur.Left.Val
				if cur.Right != nil {
					cur.Left.Val -= cur.Right.Val
				}
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Val = depth_sum[depth] - cur.Right.Val
				if cur.Left != nil {
					cur.Right.Val -= pre_left_val
				}
				q.PushBack(cur.Right)
			}
		}
		depth++
	}
	return root
}
