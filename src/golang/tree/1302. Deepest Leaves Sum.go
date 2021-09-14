package tree

import "container/list"

func deepestLeavesSum(root *TreeNode) int {
	var q list.List
	q.PushBack(root)
	var res int = 0
	for q.Len() > 0{
		var l int = q.Len()
		res = 0
		for i := 0;i < l;i++{
			cur := q.Front().Value.(*TreeNode)
			res += cur.Val
			q.Remove(q.Front())
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
		}
	}
	return res
}