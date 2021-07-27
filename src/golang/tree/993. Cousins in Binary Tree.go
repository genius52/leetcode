package tree

import "container/list"

func isCousins(root *TreeNode, x int, y int) bool {
	var q list.List
	q.PushBack(root)
	for q.Len() > 0{
		var l int = q.Len()
		var exist bool = false
		for i := 0;i < l;i++{
			var cur *TreeNode = q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if cur.Val == x || cur.Val == y{
				if exist{
					return true
				}else{
					exist = true
				}
			}
			if cur.Left != nil && cur.Right != nil{
				if (cur.Left.Val == x || cur.Left.Val == y) && (cur.Right.Val == x || cur.Right.Val == y) {
					return false
				}
			}
			if cur.Left != nil{
				q.PushBack(cur.Left)
			}
			if cur.Right != nil{
				q.PushBack(cur.Right)
			}
		}
		if exist{
			return false
		}
	}
	return false
}